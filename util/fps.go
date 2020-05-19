package util

import (
	"time"
)

type fps struct {
	lastTime    time.Time
	accumulated float64 // ms
	fps         int
}

// FPS interface
type FPS interface {
	GetAccumulated() float64
	Reset()
	GetDeltaT() float64
	SetFPS() FPS
	GetFPS() int
}

// NewFps gets a new FPS struct
func NewFps() FPS {
	return &fps{
		lastTime:    time.Now(),
		accumulated: 0,
		fps:         0,
	}
}

func (f *fps) GetDeltaT() float64 {
	return time.Since(f.lastTime).Seconds()
}

func (f *fps) GetAccumulated() float64 {
	f.accumulated += f.GetDeltaT()
	f.lastTime = time.Now()
	return f.accumulated
}

func (f *fps) Reset() {
	f.accumulated = 0
	f.lastTime = time.Now()
}

func (f *fps) SetFPS() FPS {
	dt := f.GetDeltaT()
	f.fps = int(1.0 / dt)
	return f
}

func (f *fps) GetFPS() int {
	return f.fps
}
