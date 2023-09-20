package main

import "fmt"

// Интерфейс состояний
type State interface {
	doAction(context *Context)
}

// Реализация состояния 1
type State1 struct {
}

func (s *State1) doAction(context *Context) {
	fmt.Println("Выполняется действие в состоянии 1")
	context.setState(&State2{})
}

// Реализация состояния 2
type State2 struct {
}

func (s *State2) doAction(context *Context) {
	fmt.Println("Выполняется действие в состоянии 2")
	context.setState(&State1{})
}

// Контекст
type Context struct {
	state State
}

func (c *Context) setState(state State) {
	c.state = state
}

func (c *Context) doAction() {
	c.state.doAction(c)
}

func main() {
	context := &Context{state: &State1{}}

	context.doAction()
	context.doAction()
}
