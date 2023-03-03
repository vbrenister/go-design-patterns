package structural

type sammysangAdapter struct {
	sstv *SammysangTV
}


func (tv *sammysangAdapter) turnOn()  {
	tv.sstv.tvOn = true
}

func (tv *sammysangAdapter) turnOff()  {
	tv.sstv.tvOn = false
}

func (tv *sammysangAdapter) volumeUp()  {
	tv.sstv.currentVol++
}

func (tv *sammysangAdapter) volumeDown() {
	tv.sstv.currentVol--
}

func (tv *sammysangAdapter) channelUp() {
	tv.sstv.currentChan++
}

func (tv *sammysangAdapter) channelDown() {
	tv.sstv.currentChan--
}

func (st *sammysangAdapter) goToChannel(ch int) {
	st.sstv.currentChan = ch
}