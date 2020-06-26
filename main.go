package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	winWidth  = 800
	winHeight = 600
)

var pixels = make([]byte, winWidth*winHeight*4)

type rgba struct {
	r, g, b, a byte
}

type position struct {
	x, y int
}

func drawPixel(x, y int, c rgba) {
	index := (x + winWidth*y) * 4
	if index < len(pixels) && index > 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
		pixels[index+3] = c.a
	}
}

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, _ := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	defer window.Destroy()

	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer renderer.Destroy()

	texture, _ := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	defer texture.Destroy()

	for {
		frameStart := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		texture.Update(nil, pixels, winWidth*4)
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		elapsedTime := uint32(time.Since(frameStart).Milliseconds())
		if elapsedTime < 10 {
			sdl.Delay(10 - elapsedTime)
		}
	}
}
