package adapter

import (
	"testing"
)


func TestSohnee(t *testing.T) {
	var tv television = &SohneeTV{}

	tv.channelUp()
	tv.volumeUp()
	tv.volumeUp()
	tv.turnOn()
}

func TestShouldAdapt(t *testing.T) {
	var tv television = &sammysangAdapter{
		sstv: &SammysangTV{
		},
	}

	tv.channelUp()
	tv.volumeUp()
	tv.volumeUp()
	tv.turnOn()
}