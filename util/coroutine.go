package util

func Go(cb func()) {
	go Protect(cb)
}

func Protect(g func()) {
	defer HandlePanic()
	g()
}
