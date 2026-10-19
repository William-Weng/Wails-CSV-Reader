package main

import (
	"embed"
	"wails-csv-reader/backend"

	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
}

func main() {

	fileService := &backend.FileService{}

	app := application.New(application.Options{
		Name:        "CSV Reader",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(fileService),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:          "CSV 讀取器",
		Width:          1000,
		Height:         618,
		EnableFileDrop: true,
		URL:            "/",
	}).OnWindowEvent(events.Common.WindowFilesDropped, func(event *application.WindowEvent) {

		ctx := event.Context()
		files := ctx.DroppedFiles()
		target := ctx.DropTargetDetails()

		if len(files) > 0 {

			emitKey := "csv-file-dropped"
			emitData := map[string]any{
				"elementID": target.ElementID,
				"path":      files[0],
			}

			application.Get().Event.Emit(emitKey, emitData)
		}
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
