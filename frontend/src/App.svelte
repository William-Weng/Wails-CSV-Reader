<script lang="ts">
    import { Browser, Events } from "@wailsio/runtime";
    import { onMount } from "svelte";
    import { dialog } from "./lib/dialog";
    import { typeColors } from "./lib/constant";
    import { ReadResources } from "../bindings/wails-csv-reader/backend/fileservice";

    let types = $state<string[]>([]);
    let currentFilePath = $state("");
    let searchText = $state("");

    let items = $state<CardItem[]>([]);

    let filteredItems = $derived.by(() => {
        const keyword = searchText.trim().toLowerCase();

        if (!keyword) {
            return items;
        }

        return items.filter((item) => {
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
    });

    async function readCSV(data: DroppedData) {
        try {
            const result = await ReadResources(data.path);

            types = result.types ?? [];

            items = (result.resources ?? []).map((resource, index) => ({
                id: `${index}-${resource.url}`,
                title: resource.title,
                description: resource.description,
                url: resource.url,
                example: resource.example ?? "",
                types: resource.types ?? [],
            }));

            // 資料成功載入後，才更新目前顯示的檔案路徑。
            currentFilePath = data.path;
        } catch (err) {
            dialog("warning", "錯誤", `讀取 CSV 失敗：${err}`);
        }
    }

    onMount(() => {
        const unsubscribeDrop = Events.On("csv-file-dropped", (event) => {
            void readCSV(event.data as DroppedData);
        });

        return () => {
            unsubscribeDrop();
        };
    });

    function getTypeColor(type: string): string {
        const index = types.indexOf(type);

        if (index < 0) {
            return "#51515C";
        }

        return typeColors[index % typeColors.length]!;
    }

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
</script>

<div class="app">
    <header class="toolbar">
        <div class="toolbar-info">
            <h1>CSV 讀取器</h1>
            <span>共 {items.length} 筆</span>
        </div>

        <div class="toolbar-file">
            {#if currentFilePath}
                <span class="file-label">目前檔案</span>
                <span class="file-path" title={currentFilePath}>
                    {currentFilePath}
                </span>
            {:else}
                <span class="file-placeholder">
                    請拖入 CSV 檔案
                </span>
            {/if}
        </div>
    </header>

    <main class="scroll-area" data-file-drop-target>
        <div class="card-grid">
            {#each filteredItems as item (item.id)}
                <article class="card">
                    {#if item.url}
                        <button
                            class="card-title clickable"
                            type="button"
                            title={`${item.title}`}
                            onclick={() => {
                                if (item.url) {
                                    void openExternalURL(item.url);
                                }
                            }}
                        >
                            <span>{item.title}</span>
                        </button>
                    {:else}
                        <h2 class="card-title" title={item.title}>
                            <span>{item.title}</span>
                        </h2>
                    {/if}

                    <div class="card-description">
                        <p>{item.description}</p>
                    </div>

                    {#if item.example}
                        <div class="card-example">
                            <pre><code>{item.example}</code></pre>
                        </div>
                    {/if}

                    {#if item.types.length > 0}
                        <div class="card-tags" aria-label="分類">
                            {#each item.types as type (type)}
                                <span
                                    class="type-tag"
                                    style:--type-color={getTypeColor(type)}
                                >
                                    {type}
                                </span>
                            {/each}
                        </div>
                    {/if}
                </article>
            {:else}
                <p class="empty-state">
                    找不到符合「{searchText}」的資料
                </p>
            {/each}
        </div>
    </main>

    <footer class="search-bar">
        <div class="search-content">
            <label for="card-search">Search</label>

            <div class="search-input-wrap">
                <input
                    id="card-search"
                    type="search"
                    placeholder="搜尋標題、說明、網址、範例或分類…"
                    bind:value={searchText}
                    onkeydown={(event) => {
                        if (event.key === "Escape") {
                            searchText = "";
                        }
                    }}
                />
            </div>

            <span class="search-count" role="status">
                {filteredItems.length} / {items.length} 筆
            </span>
        </div>
    </footer>
</div>

<style lang="less">
  @import "./less/main.less";
</style>
