package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const coldef = termbox.ColorDefault

func drawBox(x, y int) {
	termbox.Clear(coldef, coldef)
	termbox.SetCell(x, y, '┏', coldef, coldef)
	termbox.SetCell(x+1, y, '┓', coldef, coldef)
	termbox.SetCell(x, y+1, '┗', coldef, coldef)
	termbox.SetCell(x+1, y+1, '┛', coldef, coldef)
	termbox.Flush()
}

func timerLoop(tch chan bool) {
	for {
		tch <- true
		time.Sleep(time.Millisecond * 100)
	}
}

func keyEventLoop(kch chan termbox.Key) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			kch <- ev.Key
		default:
		}
	}
}

func controller(kch chan termbox.Key, tch chan bool) {
	for {
		select {
		case key := <-kch: // key event
			switch key {
			case termbox.KeyEsc:
				return
			}
		case <-tch: // timer event
			w, h := termbox.Size()
			drawBox(rand.Intn(w), rand.Intn(h))
			break
		default:
			break
		}
	}
}

func main() {
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	kch := make(chan termbox.Key)
	tch := make(chan bool)
	go keyEventLoop(kch)
	go timerLoop(tch)
	controller(kch, tch)
}
