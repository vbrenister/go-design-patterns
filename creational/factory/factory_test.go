package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestMagazineIsCreated(t *testing.T) {
	mgz, err := newPublication("magazine", "Rolling Stones", 100, "New York")

	if err != nil {
		t.Error("new magazine should be created")
	}

	if mag, ok := mgz.(*magazine);  !ok {
		t.Error("new magazine should be created")
	} else {
		assert.Equal(t, "This is a magazine name: Rolling Stones", mag.String())
	}
}

func TestNewsPaperIsCreated(t *testing.T) {
	mgz, err := newPublication("newspaper", "Times", 100, "New York")

	if err != nil {
		t.Error("new newspaper should be created")
	}

	if mag, ok := mgz.(*newspaper);  !ok {
		t.Error("new newspaper should be created")
	} else {
		assert.Equal(t, "This is a newspaper name: Times", mag.String())
	}
}

func TestErrorWhenUnsupportedType(t *testing.T) {
	_, err := newPublication("comic", "Batman", 200, "Washington")

	if assert.Error(t, err) {
		assert.EqualError(t, err, "no such publication type")
	}
}