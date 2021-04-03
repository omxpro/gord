package tviewutil

import "github.com/cainy-a/gord/tview"

func FocusNextIfPossible(direction tview.FocusDirection, app *tview.Application, focused tview.Primitive) {
	if focused == nil {
		return
	}

	focusNext := focused.NextFocusableComponent(direction)
	if focusNext != nil {
		app.SetFocus(focusNext)
	}
}
