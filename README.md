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
