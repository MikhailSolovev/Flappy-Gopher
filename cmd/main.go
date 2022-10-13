package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"log"
	"time"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		log.Fatalf("could not initialize TTF: %v", err)
	}
	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("could not create window: %v", err)
	}
	sdl.PumpEvents()
	defer w.Destroy()

	if err := drawTitle(r, "Flappy gopher"); err != nil {
		log.Fatalf("could not draw title: %v", err)
	}

	time.Sleep(5 * time.Second)

	if err := drawBackground(r); err != nil {
		log.Fatalf("could not draw background: %v", err)
	}

	time.Sleep(1000 * time.Second)
}

func drawBackground(r *sdl.Renderer) error {
	r.Clear()

	t, err := img.LoadTexture(r, "./pkg/imgs/background.png")
	if err != nil {
		return fmt.Errorf("could not load background image: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	r.Present()

	return nil
}

func drawTitle(r *sdl.Renderer, title string) error {
	r.Clear()

	f, err := ttf.OpenFont("./pkg/fonts/Silkscreen-Regular.ttf", 20)
	if err != nil {
		return fmt.Errorf("could not load font: %v", err)
	}
	defer f.Close()

	s, err := f.RenderUTF8Solid(title, sdl.Color{
		R: 255,
		G: 100,
		B: 0,
		A: 255,
	})
	if err != nil {
		return fmt.Errorf("could not render title: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	r.Present()

	return nil
}
