package processor

import (
	"github.com/jurevic/facegrinder/pkg/api/v1/processor/stats"
	"reflect"
	"strings"
)

// Register your processor here
var ProcessorsList = []ListEntry{
	{Key: "stats_fps", Name: "FPS counter", Processor: new(stats.Fps)},
}

type ListEntry struct {
	Key       string
	Name      string
	Processor interface{}
}

var ProcessorsMap map[string]MapEntry

type MapEntry struct {
	Name        string            `json:"name"`
	Params      interface{}       `json:"params"`
	Types       map[string]string `json:"types"`
	IsReader    bool              `json:"is_reader"`
	IsProcessor bool              `json:"is_processor"`
}

func InitProcessorsMap() {
	ProcessorsMap = make(map[string]MapEntry)

	for i := range ProcessorsList {
		key := ProcessorsList[i].Key
		name := ProcessorsList[i].Name
		processor := ProcessorsList[i].Processor

		// Load defaults
		setDefaults(processor)

		// Get types listing
		types := getTypes(processor)

		e := MapEntry{
			Name:        name,
			Params:      processor,
			Types:       types,
			IsReader:    isFrameReader(processor),
			IsProcessor: isFrameProcessor(processor),
		}

		ProcessorsMap[key] = e
	}
}

func setDefaults(i interface{}) {
	o, ok := i.(Defaulter)
	if ok {
		err := o.Default()
		if err != nil {
			panic("could not set defaults")
		}
	}
}

func isFrameReader(i interface{}) bool {
	_, ok := i.(FrameReader)
	if !ok {
		return false
	}
	return true
}

func isFrameProcessor(i interface{}) bool {
	_, ok := i.(FrameProcessor)
	if !ok {
		return false
	}
	return true
}

func getTypes(i interface{}) map[string]string {
	v := reflect.ValueOf(i)

	// If we have pointer take referenced object
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Read json exported vars and mark their JsTypes
	t := make(map[string]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		jsonTag := v.Type().Field(i).Tag.Get("json")
		if jsonTag != "" {
			fieldName := ""
			if strings.Contains(jsonTag, ",") {
				idx := strings.Index(jsonTag, ",")
				fieldName = jsonTag[:idx]
			} else {
				fieldName = jsonTag
			}

			if fieldName == "" {
				fieldName = v.Type().Field(i).Name
			}

			fieldType := v.Field(i).Type()
			t[fieldName] = getTypeFromMapping(fieldType)
		}
	}

	return t
}

var customTypeMapping = map[string]string{
	"time.Time":   "date-time",
	"color.RGBA":  "color",
	"image.Point": "point",
}

var goTypeMapping = map[reflect.Kind]string{
	reflect.Bool:    "boolean",
	reflect.Int:     "integer",
	reflect.Int8:    "integer",
	reflect.Int16:   "integer",
	reflect.Int32:   "integer",
	reflect.Int64:   "integer",
	reflect.Uint:    "integer",
	reflect.Uint8:   "integer",
	reflect.Uint16:  "integer",
	reflect.Uint32:  "integer",
	reflect.Uint64:  "integer",
	reflect.Float32: "number",
	reflect.Float64: "number",
	reflect.String:  "string",
	reflect.Slice:   "array",
	reflect.Struct:  "object",
	reflect.Map:     "object",
}

// Get JsType from go reflect type
func getTypeFromMapping(t reflect.Type) string {
	if v, ok := customTypeMapping[t.String()]; ok {
		return v
	}

	if v, ok := goTypeMapping[t.Kind()]; ok {
		return v
	}

	return ""
}
