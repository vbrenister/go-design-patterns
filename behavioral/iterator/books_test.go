package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


var lib *Library = &Library{
	Collection: []Book{
		{
			Name:      "War and Peace",
			Author:    "Leo Tolstoy",
			PageCount: 864,
			Type:      HardCover,
		},
		{
			Name:      "Crime and Punishment",
			Author:    "Leo Tolstoy",
			PageCount: 1225,
			Type:      SoftCover,
		},
		{
			Name:      "Brave New World",
			Author:    "Aldous Huxley",
			PageCount: 325,
			Type:      PaperBack,
		},
		{
			Name:      "Catcher in the Rye",
			Author:    "J.D. Salinger",
			PageCount: 206,
			Type:      HardCover,
		},
		{
			Name:      "To Kill a Mockingbird",
			Author:    "Harper Lee",
			PageCount: 399,
			Type:      PaperBack,
		},
		{
			Name:      "The Grapes of Wrath",
			Author:    "John Steinbeck",
			PageCount: 464,
			Type:      HardCover,
		},
		{
			Name:      "Wuthering Heights",
			Author:    "Emily Bronte",
			PageCount: 288,
			Type:      EBook,
		},
	},
}


func TestShouldIterate(t *testing.T) {
	iterator := lib.createIterator()

	assert.True(t, iterator.hasNext())
	assert.Equal(t, iterator.next(), &lib.Collection[0])
}