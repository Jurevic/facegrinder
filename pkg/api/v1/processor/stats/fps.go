package stats

import (
	"fmt"
	"github.com/mcuadros/go-defaults"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"time"
)

type Fps struct {
	FontThickness   int         `json:"font_thickness"`
	FontSize        float64     `json:"font_size"`
	FontColor       color.RGBA  `json:"font_color"`
	BasePoint       image.Point `json:"base_point"`
	RefreshInterval float64     `json:"refresh_interval"`

	t      time.Time
	lt     time.Time
	frames int
	text   string
}

func (o *Fps) Default() (err error) {
	o.FontThickness = 1
	o.FontSize = 1
	o.FontColor = color.RGBA{R: 0, G: 0, B: 0, A: 0}
	o.BasePoint = image.Point{X: 10, Y: 10}
	o.RefreshInterval = 1

	return
}

func (o *Fps) Init(params map[string]interface{}) (err error) {
	defaults.SetDefaults(o)

	o.t = time.Now()
	o.frames = 0

	return
}

func (o *Fps) Process(frame *gocv.Mat) (err error) {
	o.t = time.Now()
	o.frames++
	if o.t.Sub(o.lt).Seconds() >= o.RefreshInterval {
		o.text = fmt.Sprintf("fps: %d", o.frames/int(o.RefreshInterval))
		o.frames = 0
		o.lt = o.t
	}

	gocv.PutText(frame, o.text, o.BasePoint, gocv.FontHersheyDuplex, o.FontSize, o.FontColor, o.FontThickness)

	return
}
