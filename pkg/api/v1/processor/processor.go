package processor

import (
	"gocv.io/x/gocv"
	"github.com/mohae/deepcopy"
	"errors"
)

type Initializer interface {
	Init(map[string]interface{}) error
}

type ContextInitializer interface {
	InitCtx(*map[string]interface{}) error
}

type Defaulter interface {
	Default() error
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
	ProcessingNodes []FrameProcessor
	ChainContext map[string]interface{}
}

func (o *ProcessingChain) Init(p *Processor) (err error) {
	// Create context
	o.ChainContext = make(map[string]interface{})

	// Create nodes
	o.ProcessingNodes = make([]FrameProcessor, len(p.Nodes))

	for i := range p.Nodes {
		// Load defaults
		entry := ProcessorsMap[p.Nodes[i].Key]
		newNode, ok := deepcopy.Copy(entry.Default).(FrameProcessor)
		if !ok {
			err = errors.New("node does not implement frame processor interface")
			return err
		}

		// Initialize node
		err = initNode(newNode, p.Nodes[i].Params)
		if err != nil {
			return err
		}

		// Init context
		err = initNodeCtx(newNode, &o.ChainContext)
		if err != nil {
			return err
		}

		// Set node
		o.ProcessingNodes[i] = newNode
	}

	return
}

func (o *ProcessingChain) Close() (err error) {
	for i := range o.ProcessingNodes{
		ini, ok := o.ProcessingNodes[i].(Closer)
		if ok {
			err = ini.Close()
			if err != nil {
				return err
			}
		}
	}

	return
}

func (o *ProcessingChain) Run() (err error) {
	var frame *gocv.Mat

	for {
		for i := range o.ProcessingNodes {
			// Returns frame if implements reader
			rframe, err := processRead(o.ProcessingNodes[i])
			if err != nil {
				return err
			}
			if rframe != nil {
				frame = rframe
			}

			err = o.ProcessingNodes[i].Process(frame)
			if err != nil {
				return err
			}
		}
	}

	return
}

func initNode(n interface{}, params map[string]interface{}) (err error) {
	ini, ok := n.(Initializer)
	if ok {
		err = ini.Init(params)
		if err != nil {
			return err
		}
	}

	return
}

func initNodeCtx(n interface{}, ref *map[string]interface{}) (err error) {
	ini, ok := n.(ContextInitializer)
	if ok {
		err = ini.InitCtx(ref)
		if err != nil {
			return err
		}
	}

	return
}


func processRead(n interface{}) (mat *gocv.Mat, err error) {
	p, ok := n.(FrameReader)
	if ok {
		mat, err = p.Read()
		if err != nil {
			return nil, err
		}
	}

	return
}
