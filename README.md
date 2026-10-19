# CSV 讀取器

使用 Wails、Go 與 Svelte 製作的 **CSV 卡片瀏覽小工具**。將 CSV 檔案拖入視窗，即可把每列資料轉成卡片，方便瀏覽素材網站、指令筆記與程式碼範例。

https://github.com/user-attachments/assets/9ccf5aaf-bc21-43e2-96cd-d1942e40eaaf

## 功能

- 拖放 CSV 檔案載入資料。
- 顯示目前檔案的完整路徑與資料筆數。
- 以卡片呈現標題、說明、網址、範例與分類。
- 搜尋標題、說明、網址、範例或分類。
- 點擊卡片標題，使用系統預設瀏覽器開啟網址。
- 支援多個分類，並產生去重、排序後的分類清單。
- 支援 UTF-8 BOM 標題列。

## CSV 欄位

第一列為欄位名稱，請使用以下拼字與大小寫，例如 `Name`，而不是 `NAME`。

| 欄位 | 必填 | 用途 |
| --- | --- | --- |
| `Name` | 是 | 卡片標題 |
| `Notes` | 否 | 說明文字 |
| `URL` | 否 | 點擊標題時開啟的網址 |
| `Type` | 否 | 分類；可使用逗號分隔多個分類 |
| `Example` | 否 | 指令或程式碼範例 |

欄位順序不限，其他欄位（例如 `Level`）目前不會顯示。

## 使用方式

1. 準備包含 `Name` 標題欄位的 CSV 檔案。
2. 將檔案拖入應用程式的卡片區域。
3. 瀏覽卡片，或使用底部搜尋列篩選資料。
4. 按下 `Escape` 清除搜尋內容。

有 `Example` 內容的卡片會顯示範例框；沒有範例的卡片則不會顯示。

## CSV 範例

### 素材網站

```csv
Name,Notes,URL,Type,Level
icon-icons,圖示下載,[https://icon-icons.com/,Icon,1](https://icon-icons.com/,Icon,1)
Free3D,3D 模型下載,[https://free3d.com/,3D,3](https://free3d.com/,3D,3)
素材工具,多種類型素材,[https://example.com/](https://example.com/),"Sound,Image,Font",2
```

### 指令筆記

```csv
Name,Notes,URL,Example,Level
git clone <url>,複製遠端儲存庫,[https://git-scm.com/docs/git-clone,git](https://git-scm.com/docs/git-clone,git) clone [https://github.com/William-Weng/Demo.git,1](https://github.com/William-Weng/Demo.git,1)
git remote,顯示遠端儲存庫清單,[https://git-scm.com/docs/git-remote,git](https://git-scm.com/docs/git-remote,git) remote,1
```

## 格式注意事項

- 建議使用 UTF-8 編碼。
- `URL` 請填入完整的 `http://` 或 `https://` 網址，不要使用 Markdown 連結格式。
- 欄位內容包含逗號或換行時，請使用雙引號包住。
- 被雙引號包住的內容若包含雙引號，請以 `""` 表示。

例如，多個分類需放在同一個欄位：

```csv
Name,Notes,Type
素材網站,提供音效與圖片,"Sound,Image"
```

## 技術

- Wails：桌面應用程式
- Go：CSV 解析與檔案讀取
- Svelte + TypeScript：介面與搜尋
- LESS：樣式管理

## [建置指令整理](https://v3.wails.io/zh-tw/guides/build/building/)

| 目的 | 指令 |
| --- | --- |
| 安裝 Wails 3 | `go install github.com/wailsapp/wails/v3/cmd/wails3@latest` |
| 建立新專案 | `wails3 init -n <專案名稱> -t <前端框架>` |
| 安裝 LESS | `cd frontend; npm install -D less` |
| 產生 bindings | `wails3 generate bindings` |
| 更新 Windows 與 macOS 專用圖示檔案 | `wails3 generate icons -input build/appicon.png -windowsfilename build/windows/icon.ico -macfilename build/darwin/icons.icns` |
| 更新建置資源 | `wails3 update build-assets -config build/config.yml -dir build` |
| 拉取（下載）用於跨平台交叉編譯的 Docker 映像檔 | `wails3 task setup:docker` |
| 打包 macOS arm64 | `wails3 package GOOS=darwin GOARCH=arm64` |
| 建置 Windows x64 | `wails3 build GOOS=windows GOARCH=amd64` |
| 打包 Linux arm64 | `wails3 build GOOS=linux GOARCH=arm64` |
