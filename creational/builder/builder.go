package builder

import "fmt"

type NotificaitonBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificaitonBuilder {
	return &NotificaitonBuilder{}
}

func (b *NotificaitonBuilder) SetTitle(title string) {
	b.Title = title
}

func (b *NotificaitonBuilder) SetSubtitle(subtitle string) {
	b.SubTitle = subtitle
}

func (b *NotificaitonBuilder) SetMessage(message string) {
	b.Message = message
}

func (b *NotificaitonBuilder) SetImage(image string) {
	b.Image = image
}

func (b *NotificaitonBuilder) SetIcon(icon string) {
	b.Icon = icon
}

func (b *NotificaitonBuilder) SetPriority(pri int) {
	b.Priority = pri
}

func (b *NotificaitonBuilder) SetType(notType string) {
	b.NotType = notType
}

func (b *NotificaitonBuilder) Build() (*Notificaiton, error) {
	if b.Priority > 5 {
		return nil, fmt.Errorf("priority can't be higher than 5")
	}

	return &Notificaiton{
		title:    b.Title,
		subtitle: b.SubTitle,
		message:  b.Message,
		image:    b.Image,
		icon:     b.Icon,
		priority: b.Priority,
		notType:  b.NotType,
	}, nil
}
