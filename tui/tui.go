package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

func updateSelectedTheme(list *tview.List, currentTheme string) {
	for i := 0; i < list.GetItemCount(); i++ {
		theme, _ := list.GetItemText(i)
		theme = strings.TrimPrefix(theme, "[*] ")

		if theme == currentTheme {
			theme = "[*] " + theme
		}

		list.SetItemText(i, theme, "")
	}
}

var bgColor = tcell.NewHexColor(0x06070d)

func Start() error {
	app := tview.NewApplication()
	pages := tview.NewPages()
	tview.Styles.PrimitiveBackgroundColor = bgColor
	pages.Box.SetBackgroundColor(bgColor)

	mainPage, err := CreateMainpage(app, pages)
	if err != nil {
		return err
	}
	themesPage, err := CreateThemesPage(app, pages)
	if err != nil {
		return err
	}
	presetsPage, err := CreatePresetsPage()
	if err != nil {
		return err
	}

	pages.AddPage("main", mainPage, true, true)
	pages.AddPage("themes", themesPage, true, false)
	pages.AddPage("presets", presetsPage, true, false)

	return app.SetRoot(pages, true).SetFocus(pages).Run()
}
