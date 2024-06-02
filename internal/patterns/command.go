/* Command is behavioral design pattern that converts requests or simple operations into objects.

The conversion allows deferred or remote execution of commands, storing command history, etc. */

package patterns

import "fmt"

type Command interface {
	execute()
}

type Device interface {
	on()
	off()
}

// Invoker
type Button struct {
	Command Command
}

func (button *Button) Press() {
	button.Command.execute()
}

// Receiver
type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

//commands

type OnCommand struct {
	Device Device
}

func (c *OnCommand) execute() {
	c.Device.on()
}

type OffCommand struct {
	Device Device
}

func (c *OffCommand) execute() {
	c.Device.off()
}
