package gcom

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/872409/gatom/util"
)
var (
	// ErrConcurrencyComponentTimeout 并发组件业务超时
	ErrConcurrencyComponentTimeout = errors.New("Concurrency Component Timeout")
)



// Component 组件接口
type Component interface {
	// 添加一个子组件
	Mount(c Component, components ...Component) error
	// 移除一个子组件
	Remove(c Component) error
	// 执行当前组件业务:`BusinessLogicDo`和执行子组件:`ChildsDo`
	// ctx 业务上下文
	// currentConponent 当前组件
	// wg 父组件的waitgroup对象
	Do(ctx *Context, currentConponent Component, wg *sync.WaitGroup) error
	// 执行当前组件业务逻辑
	// resChan 回写当前组件业务执行结果的channel
	BusinessLogicDo(resChan chan interface{}) error
	// 执行子组件
	ChildsDo(ctx *Context) error
}

// BaseComponent 基础组件
// 实现Add:添加一个子组件
// 实现Remove:移除一个子组件
type BaseComponent struct {
	// 子组件列表
	ChildComponents []Component
}

// Mount 挂载一个子组件
func (bc *BaseComponent) Mount(c Component, components ...Component) (err error) {
	bc.ChildComponents = append(bc.ChildComponents, c)
	if len(components) == 0 {
		return
	}
	bc.ChildComponents = append(bc.ChildComponents, components...)
	return
}

// Remove 移除一个子组件
func (bc *BaseComponent) Remove(c Component) (err error) {
	if len(bc.ChildComponents) == 0 {
		return
	}
	for k, childComponent := range bc.ChildComponents {
		if c == childComponent {
			fmt.Println(util.RunFuncName(), "移除:", reflect.TypeOf(childComponent))
			bc.ChildComponents = append(bc.ChildComponents[:k], bc.ChildComponents[k+1:]...)
		}
	}
	return
}

// Do 执行子组件
// ctx 业务上下文
// currentConponent 当前组件
// wg 父组件的waitgroup对象
func (bc *BaseComponent) Do(ctx *Context, currentConponent Component, wg *sync.WaitGroup) (err error) {
	//执行当前组件业务代码
	err = currentConponent.BusinessLogicDo(nil)
	if err != nil {
		return err
	}
	// 执行子组件
	return currentConponent.ChildsDo(ctx)
}

// BusinessLogicDo 当前组件业务逻辑代码填充处
func (bc *BaseComponent) BusinessLogicDo(resChan chan interface{}) (err error) {
	// do nothing
	return
}

// ChildsDo 执行子组件
func (bc *BaseComponent) ChildsDo(ctx *Context) (err error) {
	// 执行子组件
	for _, childComponent := range bc.ChildComponents {
		if err = childComponent.Do(ctx, childComponent, nil); err != nil {
			return err
		}
	}
	return
}

// BaseConcurrencyComponent 并发基础组件
type BaseConcurrencyComponent struct {
	// 合成复用基础组件
	BaseComponent
	// 当前组件是否有并发子组件
	HasChildConcurrencyComponents bool
	// 并发子组件列表
	ChildConcurrencyComponents []Component
	// wg 对象
	*sync.WaitGroup
	// 当前组件业务执行结果channel
	logicResChan chan interface{}
	// 当前组件执行过程中的错误信息
	Err error
}

// Remove 移除一个子组件
func (bc *BaseConcurrencyComponent) Remove(c Component) (err error) {
	if len(bc.ChildComponents) == 0 {
		return
	}
	for k, childComponent := range bc.ChildComponents {
		if c == childComponent {
			fmt.Println(util.RunFuncName(), "移除:", reflect.TypeOf(childComponent))
			bc.ChildComponents = append(bc.ChildComponents[:k], bc.ChildComponents[k+1:]...)
		}
	}
	for k, childComponent := range bc.ChildConcurrencyComponents {
		if c == childComponent {
			fmt.Println(util.RunFuncName(), "移除:", reflect.TypeOf(childComponent))
			bc.ChildConcurrencyComponents = append(bc.ChildComponents[:k], bc.ChildComponents[k+1:]...)
		}
	}
	return
}

// MountConcurrency 挂载一个并发子组件
func (bc *BaseConcurrencyComponent) MountConcurrency(c Component, components ...Component) (err error) {
	bc.HasChildConcurrencyComponents = true
	bc.ChildConcurrencyComponents = append(bc.ChildConcurrencyComponents, c)
	if len(components) == 0 {
		return
	}
	bc.ChildConcurrencyComponents = append(bc.ChildConcurrencyComponents, components...)
	return
}

// ChildsDo 执行子组件
func (bc *BaseConcurrencyComponent) ChildsDo(ctx *Context) (err error) {
	if bc.WaitGroup == nil {
		bc.WaitGroup = &sync.WaitGroup{}
	}
	// 执行并发子组件
	for _, childComponent := range bc.ChildConcurrencyComponents {
		bc.WaitGroup.Add(1)
		go childComponent.Do(ctx, childComponent, bc.WaitGroup)
	}
	// 执行子组件
	for _, childComponent := range bc.ChildComponents {
		if err = childComponent.Do(ctx, childComponent, nil); err != nil {
			return err
		}
	}
	if bc.HasChildConcurrencyComponents {
		// 等待并发组件执行结果
		bc.WaitGroup.Wait()
	}
	return
}

// Do 执行子组件
// ctx 业务上下文
// currentConponent 当前组件
// wg 父组件的waitgroup对象
func (bc *BaseConcurrencyComponent) Do(ctx *Context, currentConponent Component, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	// 初始化并发子组件channel
	if bc.logicResChan == nil {
		bc.logicResChan = make(chan interface{}, 1)
	}

	go currentConponent.BusinessLogicDo(bc.logicResChan)

	select {
	// 等待业务执行结果
	case <-bc.logicResChan:
		// 业务执行结果
		fmt.Println(util.RunFuncName(), "bc.BusinessLogicDo wait.done...")
		break
	// 超时等待
	case <-ctx.TimeoutCtx.Done():
		// 超时退出
		fmt.Println(util.RunFuncName(), "bc.BusinessLogicDo timeout...")
		bc.Err = ErrConcurrencyComponentTimeout
		break
	}
	// 执行子组件
	err = currentConponent.ChildsDo(ctx)
	return
}