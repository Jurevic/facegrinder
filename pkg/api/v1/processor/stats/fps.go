package stats

import (
	"gocv.io/x/gocv"
	"time"
	"image"
	"image/color"
	"fmt"
	"errors"
)

type Fps struct {
	FontThickness int
	RefreshInterval float64
	FontSize float64
	BasePoint image.Point
	Color color.RGBA

	t time.Time
	lt time.Time
	frames int
	text string
}

func (o *Fps) Init(params map[string]interface{}) (err error) {
	o.BasePoint = image.Point{X: params["x"].(int), Y: params["y"].(int)}

	RefreshInterval := params["refresh_interval"]
	switch RefreshInterval.(type) {
	case int:
		o.RefreshInterval = float64(RefreshInterval.(int))
	case float64:
		o.RefreshInterval = RefreshInterval.(float64)
	default:
		return errors.New("scale is neither int nor float")
	}

	FontSize := params["font_size"]
	switch FontSize.(type) {
	case int:
		o.FontSize = float64(FontSize.(int))
	case float64:
		o.FontSize = FontSize.(float64)
	default:
		return errors.New("font_size is neither int nor float")
	}

	o.FontThickness = params["font_thickness"].(int)

	var R, G, B, A uint8
	fmt.Sscanf(params["color"].(string), "R:%d G:%d B:%d A:%d", &R, &G, &B, &A)
	o.Color = color.RGBA{R, G, B, A}

	o.t = time.Now()
	o.frames = 0

	return
}

func (o *Fps) Process(frame *gocv.Mat) (err error) {
	o.t = time.Now()
	o.frames++
	if o.t.Sub(o.lt).Seconds() >= o.RefreshInterval {
		o.text = fmt.Sprintf("fps: %d", o.frames)
		o.frames = 0
		o.lt = o.t
	}

	gocv.PutText(frame, o.text, o.BasePoint, gocv.FontHersheyDuplex, o.FontSize, o.Color, o.FontThickness)

	return
}
