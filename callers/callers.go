package callers

import (
	"fmt"
	"runtime"
)

type Callers []uintptr

func NewCallers(skip int, depth int) *Callers {
	pcs := make([]uintptr, depth)
	numFrames := runtime.Callers(skip, pcs)
	pcs = pcs[:numFrames]
	return (*Callers)(&pcs)
}
func (cs *Callers) Frames() *runtime.Frames {
	return runtime.CallersFrames(*cs)
}

func (cs *Callers) FirstFrame() *runtime.Frame {
	frames := cs.Frames()
	ret, _ := frames.Next()
	return &ret
}

func (cs *Callers) FramesSlice() (ret []*runtime.Frame) {
	frames := cs.Frames()
	ret = make([]*runtime.Frame, 0, len(*cs))
	for {
		frame, more := frames.Next()
		ret = append(ret, &frame)
		if !more {
			break
		}
	}
	return ret
}

func (cs *Callers) Lines(linerFunc func(*runtime.Frame) string) (ret []string) {
	if linerFunc == nil {
		linerFunc = func(frame *runtime.Frame) string {
			return fmt.Sprintf("%s at %s:%d", frame.Function, frame.File, frame.Line)
		}
	}
	frames := cs.Frames()
	ret = make([]string, 0, len(*cs))
	for {
		frame, more := frames.Next()
		ret = append(ret, linerFunc(&frame))
		if !more {
			break
		}
	}

	return ret
}
