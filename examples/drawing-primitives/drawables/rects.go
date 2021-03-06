package drawables

import (
	"math/rand"

	"github.com/forestgiant/eff"
)

type Rects struct {
	rects       []eff.Rect
	initialized bool
}

func (r *Rects) randomRects(count int, maxX int, maxY int) []eff.Rect {
	var rects []eff.Rect
	for i := 0; i < count; i++ {
		r := eff.Rect{X: rand.Intn(maxX), Y: rand.Intn(maxY), W: rand.Intn(maxX / 2), H: rand.Intn(maxY / 2)}
		rects = append(rects, r)
	}

	return rects
}

func (r *Rects) Init(canvas eff.Canvas) {
	numRects := (canvas.Width() * canvas.Height()) / 1000
	r.rects = r.randomRects(numRects, canvas.Width(), canvas.Height())
	r.initialized = true
}

func (r *Rects) Draw(canvas eff.Canvas) {
	canvas.DrawRects(r.rects, eff.RandomColor())
}

func (r *Rects) Update(canvas eff.Canvas) {
	updateRandomRects := func() {
		for i := range r.rects {
			r.rects[i].X = rand.Intn(canvas.Width())
			r.rects[i].Y = rand.Intn(canvas.Height())
			r.rects[i].W = rand.Intn(canvas.Width() / 2)
			r.rects[i].H = rand.Intn(canvas.Height() / 2)
		}
	}

	updateRandomRects()
}

func (r *Rects) Initialized() bool {
	return r.initialized
}
