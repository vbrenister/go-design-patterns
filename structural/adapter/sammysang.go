package adapter


type SammysangTV struct {
	currentChan int
	currentVol int
	tvOn bool
}

func (tv *SammysangTV) getVolume() int {
	return tv.currentVol
}

func (tv *SammysangTV) setVolume(vol int)  {
	tv.currentVol = vol
}

func (tv *SammysangTV) getChannel() int {
	return tv.currentChan
}

func (tv *SammysangTV) setChannel(channel int)  {
	tv.currentChan = channel
}

func (tv *SammysangTV) setOnState(tvOn bool)  {
	tv.tvOn = tvOn
}