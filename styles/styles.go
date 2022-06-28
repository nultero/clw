package styles

import "github.com/gdamore/tcell/v2"

func AllBlack() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorBlack)
}

func RedFg() tcell.Style {
	return AllBlack().Foreground(tcell.ColorRed)
}
