//+build !windows

package tviewutil

import (
	tview "github.com/gord-project/gview"
)

// Escape delegates to tview escape, optionally doing additional escaping.
func Escape(text string) string {
	return tview.Escape(text)
}
