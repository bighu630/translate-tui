// Demo code for the TextArea primitive.
package main

import (
	"os/exec"
	"tui/trans"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const imgFile = "/tmp/trans.png"

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
	app.SetMouseCapture(func(event *tcell.EventMouse, m tview.MouseAction) (*tcell.EventMouse, tview.MouseAction) {
		if event.Buttons() == tcell.ButtonSecondary {
			v := app.GetFocus()
			if ta, ok := v.(*tview.TextArea); ok {
				clipboard.WriteAll(ta.GetText())
			} else if tv, ok := v.(*tview.TextView); ok {
				clipboard.WriteAll(tv.GetText(true))
			} else {
				textView.Clear()
				textView.Write([]byte("get focus not in right way"))
			}
		}

		return event, m
	})
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			if event.Modifiers() == tcell.ModCtrl {
				text := readInputString(*textArea)
				if text == "" {
					text = "nil"
				}
				target := trans.TranslateText(text)
				textView.Clear()
				textView.Write([]byte(target))
			}
		case tcell.KeyCtrlP:
			text, _ := clipboard.ReadAll()
			target := trans.TranslateText(text)
			textView.Clear()
			textArea.SetText(text, true)
			textView.Write([]byte(target))
		case tcell.KeyCtrlD:
			textView.Clear()
			cmd := exec.Command("gnome-screenshot", "-a", "--file="+imgFile)
			_, err := cmd.CombinedOutput()
			if err != nil {
				textView.Write([]byte(err.Error()))
			}
			sor, targ := trans.TranslateImg(imgFile)
			textArea.SetText(sor, true)
			textView.Write([]byte(targ))
		case tcell.KeyCtrlY:
			v := app.GetFocus()
			if ta, ok := v.(*tview.TextArea); ok {
				clipboard.WriteAll(ta.GetText())
			} else if tv, ok := v.(*tview.TextView); ok {
				clipboard.WriteAll(tv.GetText(true))
			} else {
				textView.Clear()
				textView.Write([]byte("get focus not in right way"))
			}
		default:
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
