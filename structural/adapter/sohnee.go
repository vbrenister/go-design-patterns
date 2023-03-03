package adapter


type SohneeTV struct {
	vol int
	channel int
	isOn bool
}


func (tv *SohneeTV) turnOn()  {
	tv.isOn = true
}

func (tv *SohneeTV) turnOff()  {
	tv.isOn = false
}

func (tv *SohneeTV) volumeUp()  {
	tv.vol++
}

func (tv *SohneeTV) volumeDown() {
	tv.vol--
}

func (tv *SohneeTV) channelUp() {
	tv.channel++
}

func (tv *SohneeTV) channelDown() {
	tv.channel--
}

func (st *SohneeTV) goToChannel(ch int) {
	st.channel = ch
}