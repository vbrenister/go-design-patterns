package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	title    = "Test Title"
	subtitle = "Test Sub Title"
	message  = "Test message"
	icon     = "Test icon"
	image    = "Test image"
	pri      = 3
	notType  = "test not type"
)

func TestValidNotificationBuilder(t *testing.T) {
	builder := newNotificationBuilder()

	builder.SetTitle(title)
	builder.SetSubtitle(subtitle)
	builder.SetMessage(message)
	builder.SetImage(image)
	builder.SetIcon(icon)
	builder.SetType(notType)
	builder.SetPriority(pri)

	notification, err := builder.Build()

	if err != nil {
		t.Errorf("Failed to build notification. %v", err)
	}

	assert.Equal(t, title, notification.title)
	assert.Equal(t, subtitle, notification.subtitle)
	assert.Equal(t, message, notification.message)
	assert.Equal(t, icon, notification.icon)
	assert.Equal(t, image, notification.image)
	assert.Equal(t, pri, notification.priority)
	assert.Equal(t, notType, notification.notType)
}

func TestInvalidNotificationBuilder(t *testing.T) {
	builder := newNotificationBuilder()

	builder.SetTitle(title)
	builder.SetSubtitle(subtitle)
	builder.SetMessage(message)
	builder.SetImage(image)
	builder.SetIcon(icon)
	builder.SetType(notType)
	builder.SetPriority(10)

	_, err := builder.Build()

	if assert.Error(t, err) {
		assert.EqualError(t, err, "priority can't be higher than 5")
	}

}
