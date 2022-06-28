package main

import (
	"clw/styles"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

// const blockCh = '█' // this char was surprisingly thin on a monospaced font
const blockCh = ' '

var empt = []rune{blockCh}

// TODO quick manual page

func main() {

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	scr, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = scr.Init(); err != nil {
		panic(err)
	}

	defaultStyle := styles.AllBlack()
	scr.SetStyle(defaultStyle)
	scr.Clear()

	qsig := make(chan struct{})
	keys := make(chan byte)
	showOutput := false
	changed := true

	go func() {
		defer close(keys)
		defer close(qsig)
		for {
			ev := scr.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {

				case tcell.KeyRune:
					r := toLower(ev.Rune())
					switch r {
					case 'h':
						keys <- 0
					}

				case tcell.KeyLeft:
					keys <- 0

				case tcell.KeyDown:
					keys <- 1

				case tcell.KeyEscape, tcell.KeyCtrlC:
					return

				case tcell.KeyEnter:
					showOutput = true
					return

				case tcell.KeyCtrlL:
					scr.Sync()
				}
			case *tcell.EventResize:
				if !termFits(scr) {
					invalidSize(scr)
				}
				scr.Sync()
			}

			changed = true
		}
	}()

	var tkr byte = 0

	showColors(scr)

loop:
	for {
		select {
		case <-qsig:
			break loop
		case key := <-keys:
			fmt.Println(key)

		case <-time.After(time.Millisecond * 100):
			// TODO blink cursor
			if tkr%2 == 0 {

			} else {

			}
		}

		tkr++
		if changed {
			scr.Show()
			changed = false
		}
	}

	scr.Fini()
	if showOutput {
		fmt.Println("some kinda output")
	}
}

func toLower(r rune) rune {
	if r < 97 {
		return r + 32
	}

	return r
}

func termFits(scr tcell.Screen) bool {
	w, h := scr.Size()
	return w > 40 && h > 20
}

var wrn = []rune("term screen not big enough")

func invalidSize(scr tcell.Screen) {
	scr.Clear()
	for i := 0; i < len(wrn); i++ {
		scr.SetContent(i, 0, wrn[i], empt, styles.RedFg())
	}
}
