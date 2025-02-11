package tui

import (
	"fmt"
	"github.com/rivo/tview"
	"log"
	"presetsManager/themes"
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

func StartThemes() error {
	app := tview.NewApplication()
	list := tview.NewList()

	themeList, err := themes.GetThemes()
	if err != nil {
		log.Println("Ошибка получения списка тем:", err)
		return err
	}

	currentTheme, err := themes.GetCurrentTheme()
	if err != nil {
		currentTheme = "unknown"
	}

	for _, theme := range themeList {
		t := theme
		displayName := t
		if t == currentTheme {
			displayName = fmt.Sprintf("[*] %s", t)
		}
		list.AddItem(displayName, "", 0, func() {
			err := themes.SetTheme(t)
			if err == nil {
				updateSelectedTheme(list, t)
			}
		})
	}

	list.SetBorder(true).SetTitle("Выбор темы").SetTitleAlign(tview.AlignLeft)
	list.SetDoneFunc(func() {
		app.Stop()
	})

	return app.SetRoot(list, true).Run()
}
