package draw

import (
	"errors"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

type LabeledBox struct {
	Rectangle image.Rectangle
	Label     string
}

type LabeledBoxes struct {
	Key           string     `json:"key"`
	BoxThickness  int        `json:"box_thickness"`
	FontThickness int        `json:"font_thickness"`
	FontSize      float64    `json:"font_size"`
	Color         color.RGBA `json:"color"`
	MatchColor    color.RGBA `json:"match_color"`
	Scale         float64    `json:"scale"`

	context *map[string]interface{}
}

func (o *LabeledBoxes) Default() (err error) {
	o.Key = "faces"
	o.BoxThickness = 1
	o.FontThickness = 1
	o.FontSize = 1
	o.Color = color.RGBA{R: 0, G: 0, B: 0, A: 0}
	o.MatchColor = color.RGBA{R: 255, G: 0, B: 0, A: 0}
	o.Scale = 1

	return
}

func (o *LabeledBoxes) Init(params map[string]interface{}) (err error) {
	if key, ok := params["key"]; ok {
		o.Key, ok = key.(string)
		if !ok {
			err = errors.New("key is not string type")
			return err
		}
	}
	if boxThickness, ok := params["box_thickness"]; ok {
		o.BoxThickness, ok = boxThickness.(int)
		if !ok {
			err = errors.New("box thickness is not int type")
			return err
		}
	}
	if fontThickness, ok := params["font_thickness"]; ok {
		o.FontThickness, ok = fontThickness.(int)
		if !ok {
			err = errors.New("font thickness is not int type")
			return err
		}
	}
	if fontSize, ok := params["font_size"]; ok {
		o.FontSize, ok = fontSize.(float64)
		if !ok {
			err = errors.New("font size is not float type")
			return err
		}
	}
	if myColor, ok := params["color"]; ok {
		o.Color, ok = myColor.(color.RGBA)
		if !ok {
			err = errors.New("color is not color.RGBA type")
			return err
		}
	}
	if scale, ok := params["scale"]; ok {
		o.Scale, ok = scale.(float64)
		if !ok {
			err = errors.New("scale is not int type")
			return err
		}
	}

	return
}

func (o *LabeledBoxes) InitCtx(ref *map[string]interface{}) (err error) {
	o.context = ref

	return
}

func (o *LabeledBoxes) Process(frame *gocv.Mat) (err error) {
	ctx := *o.context

	boxesPtr, ok := ctx[o.Key].(*[]LabeledBox)
	if !ok {
		err = errors.New(fmt.Sprintf("could not find labeled boxes in chain context given key %s", o.Key))
		return err
	}

	boxes := *boxesPtr

	for i := range boxes {
		// Name and color
		boxColor := o.Color
		label := ""
		if boxes[i].Label != "" {
			boxColor = o.MatchColor
			label = boxes[i].Label
		}

		rectangle := boxes[i].Rectangle

		// Rescaling
		if o.Scale != 1 {
			// Scale rectangle
			rectangle.Max.X = int(o.Scale * float64(rectangle.Max.X))
			rectangle.Max.Y = int(o.Scale * float64(rectangle.Max.Y))
			rectangle.Min.X = int(o.Scale * float64(rectangle.Min.X))
			rectangle.Min.Y = int(o.Scale * float64(rectangle.Min.Y))
		}

		// Draw rectangle around face
		gocv.Rectangle(frame, rectangle, boxColor, o.BoxThickness)

		// Draw a label
		gocv.PutText(frame, label, rectangle.Min, gocv.FontHersheyDuplex, o.FontSize, boxColor, o.FontThickness)
	}

	return
}
