package gmodule

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func GoRunError(timeoutMillisecond time.Duration, fun ...HandlerFunc) error {
	err := make(chan error)

	wg := &sync.WaitGroup{}

	go func() {
		for _, f := range fun {
			wg.Add(1)
			go func(wg *sync.WaitGroup, err chan error, fn HandlerFunc) {
				defer wg.Done()
				e := fn()
				if e != nil {
					err <- e
				}
			}(wg, err, f)
		}
		wg.Wait()
		close(err)
	}()

	select {
	case e := <-err:
		fmt.Println("e:", e)
		return e
	case <-time.After(timeoutMillisecond * time.Millisecond):
		fmt.Println("timed out..")
		return errors.New("timed out")
	}
}
func goRun(timeout time.Duration, fun ...HandlerFunc) error {
	err := make(chan error)

	wg := &sync.WaitGroup{}

	go func() {
		for _, f := range fun {
			wg.Add(1)
			go func(wg *sync.WaitGroup, err chan error, fn HandlerFunc) {
				defer wg.Done()
				e := fn()
				if e != nil {
					err <- e
				}
			}(wg, err, f)
		}
		wg.Wait()
		close(err)
	}()

	select {
	case e := <-err:
		fmt.Println("e:", e)
		return e
	case <-time.After(timeout * time.Millisecond):
		fmt.Println("timed out..")
		return errors.New("timed out")
	}
}

//
// func D2() error {
//
// 	return RunChan(2000, func() error {
// 		return GDo2(12)
// 	}, func() error {
// 		return GDo2(20)
// 	}, func() error {
// 		return GDo2(1)
// 	}, func() error {
// 		return GDo2(1)
// 	}, func() error {
// 		return GDo2(1)
// 	}, func() error {
// 		return GDo2(1)
// 	}, func() error {
// 		return GDo2(1)
// 	}, func() error {
// 		return GDo2(1)
// 	})
//
// }

//
// func D2() error {
//
// 	return RunChan(2000, func(wg *sync.WaitGroup, err chan error) {
// 		defer wg.Done()
// 		e := GDo2(1)
// 		if e != nil {
// 			err <- e
// 		}
// 	}, func(wg *sync.WaitGroup, err chan error) {
// 		defer wg.Done()
// 		e := GDo2(20)
// 		if e != nil {
// 			err <- e
// 		}
// 	})
// }
//
// func GDo2(index int) error {
// 	defer fmt.Println("do some action", ":", index, " done")
// 	fmt.Println("do some action", ":", index, " begin")
// 	if index == 20 {
// 		return errors.New(fmt.Sprintf("error index: %d", index))
// 	}
// 	return nil
// }
//
// func GDo(wg *sync.WaitGroup, err chan error, index int) {
//
// 	defer wg.Done()
// 	defer func(index int) {
// 		fmt.Println("do some action done ", ":", index)
// 	}(index)
//
// 	fmt.Println("do some action", ":", index)
// 	if index == 20 {
// 		err <- errors.New(fmt.Sprintf("error index: %v", index))
// 	}
// }
//
// func D() error {
//
// 	err := make(chan error)
// 	// succeed := make(chan int)
//
// 	wg := &sync.WaitGroup{}
//
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			wg.Add(1)
// 			(func(index int) {
// 				go GDo(wg, err, index)
// 			})(i)
// 		}
// 		wg.Wait()
// 		close(err)
// 	}()
//
// 	select {
// 	case e := <-err:
// 		fmt.Println("e:", e)
// 		// return e
// 	case <-time.After(20 * time.Second):
// 		fmt.Println("timed out..")
// 		return errors.New("timed out..")
// 	}
// 	return nil
// }
