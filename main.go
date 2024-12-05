package main

import (
	"cloc-desktop/backend/cloc"
	"cloc-desktop/backend/sys"
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// TODO: 系统托盘
// TODO: CLOC结果页Table展示

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	appCloc := cloc.NewAppCloc()
	appSys := sys.NewSystemTp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "cloc-desktop",
		Width:     1024,
		Height:    768,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			appCloc.Startup(ctx)
			appSys.Startup(ctx)
		},
		Bind: []interface{}{
			appCloc,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			BackdropType:                      windows.Acrylic,
			DisableFramelessWindowDecorations: false,
		},
		Mac: &mac.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
