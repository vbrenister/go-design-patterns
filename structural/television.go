package structural


type television interface {
	volumeUp()
	volumeDown()
	channelUp()
	channelDown()
	turnOn()
	turnOff()
	goToChannel(ch int)
}