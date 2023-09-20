package main

import "fmt"

// Интерфейс команды
type Command interface {
	Execute()
}

// Реализация конкретной команды
type PrintCommand struct {
	message string
}

func (c *PrintCommand) Execute() {
	fmt.Println(c.message)
}

// Объект, управляющий командами
type Invoker struct {
	commands []Command
}

// Добавление команды в очередь
func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

// Выполнение всех команд в очереди
func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func main() {
	invoker := &Invoker{}

	// Создание команды на основе PrintCommand и добавление в очередь
	invoker.AddCommand(&PrintCommand{message: "Hello, world!"})

	// Выполнение всех команд в очереди
	invoker.ExecuteCommands()
}
