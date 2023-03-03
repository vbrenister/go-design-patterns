package observer

import "fmt"

type observer interface {
	onUpdate(data string)
}

type DataListener struct {
	Name string
}

func (dl *DataListener) onUpdate(data string)  {
	fmt.Println("Listener:", dl.Name, "god data changed", data)
}