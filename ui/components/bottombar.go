package components

import (
	"sync"

	"github.com/cainy-a/gord/config"
	"github.com/cainy-a/gord/tview"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"github.com/rivo/uniseg"
)

// BottomBar custom simple component to render static information at the bottom
// of the application.
type BottomBar struct {
	*sync.Mutex
	*tview.Box
	items []*bottomBarItem
}

type bottomBarItem struct {
	content string
}

// Draw draws this primitive onto the screen. Implementers can call the
// screen's ShowCursor() function but should only do so when they have focus.
// (They will need to keep track of this themselves.)
func (b *BottomBar) Draw(screen tcell.Screen) bool {
	b.Lock()
	defer b.Unlock()
	hasDrawn := b.Box.Draw(screen)
	if !hasDrawn {
		return false
	}

	if len(b.items) == 0 {
		//True, as we've already drawn.
		return true
	}

	style := tcell.StyleDefault.
		//Background(config.GetTheme().PrimitiveBackgroundColor).
		Foreground(config.GetTheme().PrimaryTextColor).
		Reverse(true)

	xPos, yPos, _, _ := b.GetInnerRect()
	for _, item := range b.items {
		gr := uniseg.NewGraphemes(item.content)
		for gr.Next() {
			r := gr.Runes()
			width := runewidth.StringWidth(gr.Str())
			var comb []rune
			if len(r) > 1 {
				comb = r[1:]
			}

			screen.SetContent(xPos, yPos, r[0], comb, style)
			xPos += width
		}

		//Spacing between items
		xPos++
	}

	return true
}

// AddItem adds a new item to the right side of the already existing items.
func (b *BottomBar) AddItem(text string) {
	b.Lock()
	defer b.Unlock()
	b.items = append(b.items, &bottomBarItem{text})
}

func (b *BottomBar) InsertItemAtStart(text string) {
	b.Lock()
	defer b.Unlock()
	originalItems := b.items
	b.items = []*bottomBarItem{{content: text}}
	for i := range originalItems {
		b.items = append(b.items, originalItems[i])
	}
}

func (b *BottomBar) RemoveItemAtIndex(index int) {
	b.Lock()
	defer b.Unlock()
	originalItems := b.items
	b.items = []*bottomBarItem{}
	for i := range originalItems {
		if i != index {
			b.items = append(b.items, originalItems[i])
		}
	}
}

// NewBottomBar creates a new bar to be put at the bottom aplication.
// It contains static information and hints.
func NewBottomBar() *BottomBar {
	bottomBar := &BottomBar{
		Mutex: &sync.Mutex{},
		Box:   tview.NewBox(),
	}
	bottomBar.SetBorder(false)

	return bottomBar
}
