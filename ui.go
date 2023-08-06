// Demo code for the TextArea primitive.
package main

import (
	"tui/trans"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	textArea := tview.NewTextArea().
		SetPlaceholder("Enter text here...")
	textArea.SetTitle("From").SetBackgroundColor(tcell.ColorDefault).SetBorder(true)
	textView := tview.NewTextView()
	textView.SetTitle("To").SetBackgroundColor(tcell.ColorDefault).SetBorder(true)

	mainView := tview.NewFlex().
		AddItem(textArea, 0, 1, true).
		AddItem(textView, 0, 1, false)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter || event.Key() == tcell.KeyCtrlBackslash {
			text := readInputString(*textArea)
			target := trans.TranslateText(text)
			textView.Clear()
			textView.Write([]byte(target))
		} else if event.Key() == tcell.KeyCtrlP {
			text, _ := clipboard.ReadAll()
			target := trans.TranslateText(text)
			textView.Clear()
			textArea.SetText(text, false)
			textView.Write([]byte(target))
			// textView.SetText(target, false)
		} else if event.Key() == tcell.KeyCtrlY {
			clipboard.WriteAll(textView.GetText(true))
		}

		return event
	})

	if err := app.SetRoot(mainView,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func readInputString(tArea tview.TextArea) string {
	return tArea.GetText()
}
