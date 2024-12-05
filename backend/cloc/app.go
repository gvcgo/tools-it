package cloc

import (
	"cloc-desktop/backend/conf"
	"context"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// AppCloc struct
type AppCloc struct {
	ctx context.Context
}

// NewAppCloc creates a new App application struct
func NewAppCloc() *AppCloc {
	return &AppCloc{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *AppCloc) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *AppCloc) GetAppName() string {
	return conf.AppName
}

func (a *AppCloc) Cloc(
	byFile bool, skipDuplicated bool,
	sortTag string, excludeExt string,
	includeLang string,
	match string, notMatch string,
	matchDir string, notMatchDir string,
	paths string,
) (r string) {
	pathList := strings.Split(paths, ",")

	r = Cloc(byFile, skipDuplicated,
		sortTag, excludeExt,
		includeLang,
		match, notMatch,
		matchDir, notMatchDir,
		pathList...)
	return
}

func (a *AppCloc) SelectDirectory() string {
	r, _ := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Directory",
	})
	return r
}
