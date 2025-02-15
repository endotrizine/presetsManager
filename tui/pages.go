package tui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"presetsManager/cfg"
	"presetsManager/themes"
	"strings"
)

func CreateMainpage(app *tview.Application, pages *tview.Pages) (tview.Primitive, error) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("This is main page. Choose where to go next.")).
		AddButtons([]string{"Themes", "Presets", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonIndex == 0 {
				pages.SwitchToPage("themes")
			} else if buttonIndex == 1 {
				pages.SwitchToPage("presets")
			} else {
				app.Stop()
			}
		}).
		SetBackgroundColor(bgColor)
	modal.Box.SetBackgroundColor(bgColor)
	return modal, nil
}

func CreateThemesPage(app *tview.Application, pages *tview.Pages) (tview.Primitive, error) {
	list := tview.NewList().
		SetSelectedBackgroundColor(bgColor).
		SetSelectedTextColor(tcell.ColorBlue)

	themeList, err := themes.GetThemes()
	if err != nil {
		return nil, err
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
			if err := themes.SetTheme(t); err == nil {
				updateSelectedTheme(list, t)
			}
		})
	}

	list.
		SetBorder(true).
		SetTitle("Выбор темы").
		SetTitleAlign(tview.AlignLeft)

	list.SetDoneFunc(func() {
		pages.SwitchToPage("main")
	})

	return list, nil
}

func CreatePresetsPage() (tview.Primitive, error) {
	table := tview.NewTable().
		SetBorders(false)

	presets, err := cfg.LoadPresets()
	if err != nil {
		return nil, fmt.Errorf("ошибка загрузки пресетов: %v", err)
	}

	row := 0
	for _, p := range presets {
		preset := p
		var presetName string
		var selectable bool

		if preset.Selected {
			presetName = fmt.Sprintf("🔹 %s | [lime]Текущий[-]", preset.Name)
			selectable = false
		} else {
			presetName = fmt.Sprintf("🔹 %s", preset.Name)
			selectable = true
		}

		titleCell := tview.NewTableCell(presetName).
			SetTextColor(tcell.ColorWhite).
			SetSelectable(selectable).
			SetAlign(tview.AlignLeft)

		table.SetCell(row, 0, titleCell)
		row++

		for key, value := range preset.Params {
			paramCell := tview.NewTableCell(fmt.Sprintf("  %s: %s", key, value)).
				SetTextColor(tcell.ColorLightGray).
				SetSelectable(false).
				SetAlign(tview.AlignLeft)

			table.SetCell(row, 0, paramCell)
			row++
		}
		row++ // разделитель
	}

	table.Select(0, 0).SetFixed(1, 1).SetSelectable(true, false)

	table.SetSelectedFunc(func(row, column int) {
		cell := table.GetCell(row, column)
		presetName := cell.Text
		rows := table.GetRowCount()
		var cells []*tview.TableCell
		for i := 0; i < rows; i++ {
			cell := table.GetCell(i, column)
			cells = append(cells, cell)
		}
		for _, c := range cells {
			if c.Text != cell.Text && strings.HasPrefix(c.Text, "🔹") {
				newName := strings.ReplaceAll(c.Text, " | [lime]Текущий[-]", "")
				c.SetText(newName).SetSelectable(true)
			}
		}

		cell.SetText(fmt.Sprintf("%s | [lime]Текущий[-]", presetName)).SetSelectable(false)
		table.Select(0, 0)
		// TODO: config.ApplyPreset(presetName)
	})

	table.SetBorder(true).SetTitle(" Список пресетов ")
	return table, nil
}
