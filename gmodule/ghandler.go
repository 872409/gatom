package gmodule

type ModuleHandler interface {
	RunHandles(handlers ...HandlerFunc) error
}

func NewModuleHandler() ModuleHandler {
	return &BaseModuleHandler{}
}

type BaseModuleHandler struct {
}

func (receiver *BaseModuleHandler) RunHandles(handlers ...HandlerFunc) error {
	return RunHandles(handlers...)
}

//
// func Warp(handlers ...HandlerFunc) HandlerFunc {
// 	return func() error {
// 		return RunHandles(handlers...)
// 	}
// }
//
// func (receiver HandlerFunc) Warp(handlers ...HandlerFunc) HandlerFunc {
// 	return Warp(handlers...)
// }

type HandlerFunc func() error

func RunHandles(handlers ...HandlerFunc) error {

	for _, handler := range handlers {
		if err := handler(); err != nil {
			return err
		}
	}

	return nil
}

type Runner func(p1 interface{}, p2 interface{}) (interface{}, error)

func Run(param interface{}, runners ...Runner) (interface{}, error) {

	var lastP interface{}
	var err error

	for _, runner := range runners {
		lastP, err = runner(param, lastP)
		if err != nil {
			return lastP, err
		}
	}

	return lastP, nil
}
