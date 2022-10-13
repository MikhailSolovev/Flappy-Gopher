package internal

import (
	"context"
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type Scene struct {
	time  int
	bg    *sdl.Texture
	birds []*sdl.Texture
}

func NewScene(r *sdl.Renderer) (*Scene, error) {
	bg, err := img.LoadTexture(r, "./pkg/imgs/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}

	var birds []*sdl.Texture
	for i := 1; i < 5; i++ {
		path := fmt.Sprintf("./pkg/imgs/bird_frame_%d.png", i)
		bird, err := img.LoadTexture(r, path)
		if err != nil {
			return nil, fmt.Errorf("could not load bird: %v", err)
		}

		birds = append(birds, bird)
	}

	return &Scene{bg: bg, birds: birds}, nil
}

func (s *Scene) Run(ctx context.Context, r *sdl.Renderer) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)
		for range time.Tick(10 * time.Millisecond) {
			select {
			case <-ctx.Done():
				return
			default:
				if err := s.Paint(r); err != nil {
					errc <- err
				}
			}
		}
	}()

	return errc
}

func (s *Scene) Paint(r *sdl.Renderer) error {
	s.time++
	r.Clear()

	if err := r.Copy(s.bg, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	rect := &sdl.Rect{X: 10, Y: 300 - 43/2, W: 50, H: 43}
	idx := s.time / 100 % len(s.birds)
	if err := r.Copy(s.birds[idx], nil, rect); err != nil {
		return fmt.Errorf("could not copy bird: %v", err)
	}

	r.Present()

	return nil
}

func (s *Scene) Destroy(r *sdl.Renderer) {
	s.bg.Destroy()
}
