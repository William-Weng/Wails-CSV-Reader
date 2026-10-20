<script lang="ts">
    import { Browser, Events } from "@wailsio/runtime";
    import { onMount } from "svelte";
    import { dialog } from "./lib/dialog";
    import { typeColors } from "./lib/constant";
    import { ReadResources } from "../bindings/wails-csv-reader/backend/fileservice";

    import SearchBar from "./component/SearchBar.svelte";
    import Toolbar from "./component/Toolbar.svelte";
    import CardTitle from "./component/CardTitle.svelte";
    import CardDescription from "./component/CardDescription.svelte";
    import CardExample from "./component/CardExample.svelte";
    import CardTags from "./component/CardTags.svelte";

    let types = $state<string[]>([]);
    let currentFilePath = $state("");
    let searchText = $state("");
    let items = $state<CardItem[]>([]);

    let filteredItems = $derived.by(() => filterItems(items, searchText));

    type ReadResourcesResult = Awaited<ReturnType<typeof ReadResources>>;
    type ResourceListData = NonNullable<ReadResourcesResult>;

    onMount(() => {
        const unsubscribeDrop = Events.On("csv-file-dropped", (event) => {
            void readCSV(event.data as DroppedData);
        });

        return () => {
            unsubscribeDrop();
        };
    });

    /**
     * 讀取指定路徑的 CSV，並更新分類、卡片與目前檔案路徑
     *
     * 先完成讀取與資料轉換，再更新畫面狀態；若讀取或轉換失敗，保留原本資料並顯示錯誤提示
     *
     * @param data 拖放檔案資料，使用其中的 path 作為讀取路徑
     */
    async function readCSV(data: DroppedData): Promise<void> {
        try {
            const result = await ReadResources(data.path);

            if (result == null) {
                throw new Error("沒有取得 CSV 解析結果");
            }

            // 先完成解析，再更新畫面狀態
            const nextTypes = parseTypes(result);
            const nextItems = parseItems(result);

            types = nextTypes;
            items = nextItems;

            // 資料成功載入後，才更新目前顯示的檔案路徑
            currentFilePath = data.path;
        } catch (err) {
            const message = err instanceof Error ? err.message : String(err);

            dialog("warning", "錯誤", `讀取 CSV 失敗：${message}`);
        }
    }

    /**
     * 使用系統預設瀏覽器開啟外部網址
     *
     * 僅允許 HTTP 與 HTTPS 協定；網址解析或開啟失敗時，顯示錯誤提示
     *
     * @param value 要開啟的完整網址
     */
    async function openExternalURL(value: string) {
        try {
            const url = new URL(value);

            if (url.protocol !== "https:" && url.protocol !== "http:") {
                dialog("warning", "錯誤", `不支援的網址協定：${url.protocol}`);
                return;
            }

            await Browser.OpenURL(url.href);
        } catch (error) {
            dialog("warning", "錯誤", `無法開啟外部網址：${value} - ${error}`);
        }
    }

    /**
     * 依搜尋文字篩選卡片資料
     *
     * 搜尋範圍包含標題、說明、網址、範例與分類，比對時忽略大小寫，以及搜尋文字前後的空白
     *
     * @param source 要篩選的卡片資料
     * @param search 搜尋文字
     * @returns 符合條件的卡片資料；無關鍵字時回傳原陣列
     */
    function filterItems(source: typeof items, search: string): typeof items {
        const keyword = search.trim().toLowerCase();

        if (!keyword) {
            return source;
        }

        return source.filter((item) => {
            const searchableText = [
                item.title,
                item.description,
                item.url ?? "",
                item.example ?? "",
                ...item.types,
            ]
                .join("\n")
                .toLowerCase();

            return searchableText.includes(keyword);
        });
    }

    /**
     * 取得 CSV 解析結果中的分類清單
     *
     * @param result 已確認非空的 CSV 解析結果
     * @returns 分類清單；未提供分類時回傳空陣列
     */
    function parseTypes(result: ResourceListData): string[] {
        return result.types ?? [];
    }

    /**
     * 將 CSV 資源清單轉換為卡片顯示資料
     *
     * 僅回傳轉換結果，不直接修改畫面狀態；缺少的文字欄位以空字串代替，分類以空陣列代替
     *
     * @param result 已確認非空的 CSV 解析結果
     * @returns 供卡片列表顯示的資料陣列
     */
    function parseItems(result: ResourceListData) {
        return (result.resources ?? []).map((resource, index) => ({
            id: `${index}-${resource.url ?? ""}`,
            title: resource.title ?? "",
            description: resource.description ?? "",
            url: resource.url ?? "",
            example: resource.example ?? "",
            types: resource.types ?? [],
        }));
    }

    /**
     * 依分類在清單中的位置，取得對應的標籤顏色
     *
     * 分類數量超過色票數量時，會循環使用色票；找不到分類時，回傳預設的灰色
     *
     * @param type 分類名稱
     * @returns 分類標籤使用的色碼
     */
    function getTypeColor(type: string): string {
        const index = types.indexOf(type);

        if (index < 0) {
            return "#51515C";
        }

        return typeColors[index % typeColors.length]!;
    }
</script>

<div class="app">
    <Toolbar totalCount={items.length} {currentFilePath} />
    <main class="scroll-area" data-file-drop-target>
        <div class="card-grid">
            {#each filteredItems as item (item.id)}
                <article class="card">
                    <CardTitle title={item.title} url={item.url} {openExternalURL} />
                    <CardDescription description={item.description} />
                    <CardExample example={item.example} />
                    <CardTags types={item.types} {getTypeColor} />
                </article>
            {:else}
                <p class="empty-state">找不到符合「{searchText}」的資料</p>
            {/each}
        </div>
    </main>
    <SearchBar bind:searchText totalCount={items.length} filteredCount={filteredItems.length} />
</div>

<style lang="less">
    @import "./less/variables.less";
    @import "./less/mixins.less";
    @import "./less/base.less";
    @import "./less/cards.less";
</style>
