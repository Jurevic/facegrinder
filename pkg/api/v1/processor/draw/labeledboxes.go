package draw

import (
	"gocv.io/x/gocv"
	"image/color"
	"image"
	"errors"
	"fmt"
)

type LabeledBox struct {
	Rectangle image.Rectangle
	Label string
}

type LabeledBoxes struct {
	Key string
	BoxThickness int
	FontThickness int
	FontSize float64
	Color color.RGBA

	context *map[string]interface{}
}

func (o *LabeledBoxes) Init(params map[string]interface{}) (err error) {
	o.Key = params["key"].(string)

	FontSize := params["font_size"]
	switch FontSize.(type) {
	case int:
		o.FontSize = float64(FontSize.(int))
	case float64:
		o.FontSize = FontSize.(float64)
	default:
		return errors.New("font size is neither int nor float")
	}

	o.FontThickness = params["font_thickness"].(int)
	o.FontThickness = params["box_thickness"].(int)

	var R, G, B, A uint8
	fmt.Sscanf(params["color"].(string), "R:%d G:%d B:%d A:%d", &R, &G, &B, &A)
	o.Color = color.RGBA{R, G, B, A}

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
		//if matches[i] < 0 {
		//	frameColor = color.RGBA{R: 0, G: 255, B: 0, A: 0}
		//} else {
		//	frameColor = color.RGBA{R: 255, G: 0, B: 0, A: 0}
		//	text = strconv.Itoa(matches[i])
		//}

		// Scale back rectangle
		//if sc != 1 {
		//	// Scale back rectangle
		//	rectangle = image.Rectangle{
		//		Max: image.Point{X: sc * faces[i].Rectangle.Max.X, Y: sc * faces[i].Rectangle.Max.Y},
		//		Min: image.Point{X: sc * faces[i].Rectangle.Min.X, Y: sc * faces[i].Rectangle.Min.Y}}
		//} else {
		//	rectangle = faces[i].Rectangle
		//}

		// Draw rectangle around face
		gocv.Rectangle(frame, boxes[i].Rectangle, o.Color, o.BoxThickness)

		// Draw a label
		gocv.PutText(frame, boxes[i].Label, boxes[i].Rectangle.Min, gocv.FontHersheyDuplex, o.FontSize, o.Color, o.FontThickness)
	}

	return
}
