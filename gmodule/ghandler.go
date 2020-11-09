package gmodule

// type IGHandler interface {
// 	Handle(handlers ...HandlerFunc) error
// }

type HandlerFunc func() (err error)

// type GHandler struct {
// 	IGHandler
// }

func RunHandles(handlers ...HandlerFunc) error {

	for _, handler := range handlers {
		if err := handler(); err != nil {
			return err
		}
	}

	return nil
}
