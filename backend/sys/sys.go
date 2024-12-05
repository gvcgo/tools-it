package sys

import (
	"context"
	_ "embed"

	"github.com/energye/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed favicon.ico
var icon []byte

type SystemTp struct {
	ctx context.Context
}

func NewSystemTp() *SystemTp {
	return &SystemTp{}
}
func (a *SystemTp) systemTray() {
	systray.SetIcon(icon)

	show := systray.AddMenuItem("main", "Shows main window.")
	show.Click(func() { runtime.WindowShow(a.ctx) })
	show.SetIcon(icon)
	systray.AddSeparator()
	exit := systray.AddMenuItem("exit", "Exits it-tools.")
	exit.Click(func() { runtime.Quit(a.ctx) })

	systray.SetOnClick(func(menu systray.IMenu) { runtime.WindowShow(a.ctx) })
	systray.SetOnRClick(func(menu systray.IMenu) { menu.ShowMenu() })
	systray.SetTooltip("IT-Tools is a great application.")
}
func (a *SystemTp) Startup(ctx context.Context) {
	a.ctx = ctx
	systray.Run(a.systemTray, func() {})
}
