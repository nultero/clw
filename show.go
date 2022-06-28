package main

import (
	"clw/styles"

	"github.com/gdamore/tcell/v2"
)

const ttl byte = 255

func showColors(scr tcell.Screen) {
	w, h := scr.Size()
	cm := styles.ColorMap
	var idx byte = 0

mainloop:
	for j := 0; j < h; j++ {
		if j%2 == 0 {
			continue
		}

	subloop:
		for i := 0; i < w; i++ {
			rgb := cm[idx]
			col := tcell.NewRGBColor(rgb[0], rgb[1], rgb[2])
			// tmpSty := styles.AllBlack().Foreground(col)
			tmpSty := styles.AllBlack().Background(col)

			scr.SetContent(
				i, j, blockCh, empt, tmpSty,
			)

			idx++
			if idx == 16 || idx == 232 {
				break subloop

			} else if cm.NthBreak(idx) {
				break subloop

			} else if idx == 255 {
				break mainloop
			}
		}
	}

}
