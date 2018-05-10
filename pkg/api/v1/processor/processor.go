package processor

import (
	"fmt"
	"github.com/Kagami/go-face"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"strconv"
	"time"
	"os/exec"
	"github.com/jurevic/facegrinder/pkg/datastore"
	apiFace "github.com/jurevic/facegrinder/pkg/api/v1/face"
)

type Initializer interface {
	Init(map[string]interface{}) error
}

type ContextInitializer interface {
	InitCtx(*map[string]interface{}) error
}

type Closer interface {
	Close() error
}

type FrameReader interface {
	Read() (*gocv.Mat, error)
}

type FrameProcessor interface {
	Process(*gocv.Mat) error
}

type ProcessingChain struct {
	Input FrameReader
	Processors []FrameProcessor
	ChainContext map[string]interface{}
}

func (o ProcessingChain) Init(params map[string]interface{}) (err error) {
	o.ChainContext = make(map[string]interface{})

	ini, ok := o.Input.(Initializer)
	if ok {
		err = ini.Init(params)
		if err != nil {
			return err
		}
	}

	for i := range o.Processors{
		ini, ok = o.Processors[i].(Initializer)
		if ok {
			err = ini.Init(params)
			if err != nil {
				return err
			}
		}

		ctxini, ok := o.Processors[i].(ContextInitializer)
		if ok {
			ctxini.InitCtx(&o.ChainContext)
			if err != nil {
				return err
			}
		}
	}

	return
}

func (o ProcessingChain) Close() (err error) {
	ini, ok := o.Input.(Closer)
	if ok {
		err = ini.Close()
		if err != nil {
			return err
		}
	}

	for i := range o.Processors{
		ini, ok = o.Processors[i].(Closer)
		if ok {
			err = ini.Close()
			if err != nil {
				return err
			}
		}
	}

	return
}

func (o ProcessingChain) Run() (err error) {
	for {
		frame, err := o.Input.Read()
		if err != nil {
			return err
		}

		for i := range o.Processors{
			err = o.Processors[i].Process(frame)
			if err != nil {
				return err
			}
		}
	}

	return
}

var faceRec *face.Recognizer

var rectangle image.Rectangle
var frameColor color.RGBA
var processThisFrame int
var faces []face.Face
var matches []int

func InitKnownFaces(userId int) {
	db := datastore.DBConn()

	var faces []apiFace.Face
	err := db.Model(&faces).Where("owner_id = ?", userId).Select()
	if err != nil {
		return
	}

	var descriptors []face.Descriptor
	for i := range faces {
		knownFace, err := faceRec.RecognizeSingleFile(faces[i].Url)
		if err != nil {
			fmt.Println(err)
			return
		}
		descriptors = append(descriptors, knownFace.Descriptor)
	}

	faceRec.SetSamples(descriptors)
}

func InitFacePredictor() {
	rec, err := face.NewRecognizer("models", 0.6, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	faceRec = rec

	processThisFrame = 0
}

func ProcessFromCam() error {
	// Camera
	cam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer cam.Close()

	proc := exec.Command("ffmpeg",
		"-y",
		"-f", "rawvideo",
		"-vcodec", "rawvideo",
		"-pix_fmt", "bgr24",
		"-video_size", "640x480",
		"-i", "-",
		"-c:v", "libx264",
		"-pix_fmt", "yuv420p",
		"-preset", "ultrafast",
		"-f", "flv",
		"rtmp://0.0.0.0/1?key=password")
	pipe, _ := proc.StdinPipe()
	defer pipe.Close()
	err = proc.Start()

	// Open view window
	window := gocv.NewWindow("Test")
	defer window.Close()

	// Init frame
	frame := gocv.NewMat()
	defer frame.Close()

	var ctime time.Time
	ltime := time.Now()
	frames := 0

	for {
		ctime = time.Now()
		frames++
		if ctime.Sub(ltime).Seconds() >= 1 {
			fmt.Println(frames, "fps")
			frames = 0
			ltime = ltime.Add(time.Second)
		}

		// Grab a single frame of video
		if !cam.Read(&frame) {
			break
		}

		// Process frame
		ProcessFrame(&frame, 1, 2)

		// Output frame
		pipe.Write(frame.ToBytes())

		// show the image in the window, and wait 1 millisecond
		window.IMShow(frame)
	}

	return nil
}

func ProcessFromRtmpChannel() error {
	// input RTMP stream
	stream, err := gocv.VideoCaptureFile("rtmp://localhost/1")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer stream.Close()

	//output RTMP stream
	//params := []string{"-y",
	//	"-f", "rawvideo",
	//	"-vcodec", "rawvideo",
	//	"-pix_fmt", "bgr24",
	//	"-s", "1280x720",
	//	"-i", "-",
	//	"-c:v", "libx264",
	//	"-pix_fmt", "yuv420p",
	//	"-preset", "ultrafast",
	//	"-f", "flv",
	//	"rtmp://0.0.0.0/2?key=password"}
	//
	//proc := exec.Command("ffmpeg", strings.Join(params, " "))
	//pipe, _ := proc.StdinPipe()
	//defer pipe.Close()
	//proc.Start()

	// Reduce buffer size as large buffer introduces large delay
	// stream.Set(gocv.VideoCaptureBufferSize, 5)

	// Open view window
	window := gocv.NewWindow("Test")
	defer window.Close()

	// Init frame
	frame := gocv.NewMat()
	defer frame.Close()

	var ctime time.Time
	ltime := time.Now()
	frames := 0

	for {
		ctime = time.Now()
		frames++
		if ctime.Sub(ltime).Seconds() >= 1 {
			fmt.Println(frames, "fps")
			frames = 0
			ltime = ltime.Add(time.Second)
		}

		// Grab a single frame of video
		if !stream.Read(&frame) {
			break
		}

		// Process frame
		ProcessFrame(&frame, 1, 30)

		// Output frame
		// pipe.Write(frame.ToBytes())

		// show the image in the window, and wait 1 millisecond
		window.IMShow(frame)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

	return nil
}

func ProcessFrame(frame *gocv.Mat, sc, pc int) error {
	// Clone frame for processing
	pframe := frame.Clone()
	defer pframe.Close()

	// Resize frame to sc
	if sc != 1 {
		fx := 1.0 / float64(sc)
		fy := fx
		gocv.Resize(pframe, &pframe, image.Point{}, fx, fy, gocv.InterpolationLinear)
	}

	// Convert the image from BGR to RGB color
	gocv.CvtColor(pframe, &pframe, gocv.ColorRGBAToBGR)

	// Processing every other frame
	if processThisFrame == pc {
		// Find all the face locations and their encodings in the current frame
		faces = getFaces(frame)

		// Classify encodings
		matches = classifyFaces(faces)
		processThisFrame = 0
	}
	processThisFrame++

	// Display the results
	for i := range faces {
		text := ""

		// Name and color
		if matches[i] < 0 {
			frameColor = color.RGBA{R: 0, G: 255, B: 0, A: 0}
		} else {
			frameColor = color.RGBA{R: 255, G: 0, B: 0, A: 0}
			text = strconv.Itoa(matches[i])
		}

		// Scale back rectangle
		if sc != 1 {
			// Scale back rectangle
			rectangle = image.Rectangle{
				Max: image.Point{X: sc * faces[i].Rectangle.Max.X, Y: sc * faces[i].Rectangle.Max.Y},
				Min: image.Point{X: sc * faces[i].Rectangle.Min.X, Y: sc * faces[i].Rectangle.Min.Y}}
		} else {
			rectangle = faces[i].Rectangle
		}

		// Draw rectangle around face
		gocv.Rectangle(frame, faces[i].Rectangle, frameColor, 3)

		// Draw a label
		gocv.PutText(frame, text, rectangle.Min, gocv.FontHersheyDuplex, 2, frameColor, 1)
	}

	return nil
}

func getFaces(mat *gocv.Mat) []face.Face {
	faces, err := faceRec.RecognizeMat(mat.ToBytes(), mat.Rows(), mat.Cols())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return faces
}

func classifyFaces(faces []face.Face) []int {
	var matches []int

	for i := range faces {
		matches = append(
			matches,
			faceRec.Classify(faces[i].Descriptor))
	}

	return matches
}
