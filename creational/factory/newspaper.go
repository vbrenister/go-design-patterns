package factory

import "fmt"


type newspaper struct {
	publication
}

func (n newspaper) String() string {
	return fmt.Sprintf("This is a newspaper name: %s", n.name)
}

func createNewspaper(name string, pages int, publisher string) iPublication {
	return &newspaper{
		publication: publication{
			name: name,
			pages: pages,
			publisher: publisher,
		},
	}
} 