package channel

import (
	"github.com/nareix/joy4/av/pubsub"
	"sync"
)

type Channel struct {
	Id      string        `json:"id"`
	OwnerId string        `json:"-"`
	Name    string        `json:"name"`
	Key     string        `json:"key"`
	Que     *pubsub.Queue `json:"-"`
}

type channelBuffer struct {
	sync.RWMutex
	channels map[string]*Channel
}

var chb = newChannelBuffer()

// Channel buffer constructor
func newChannelBuffer() *channelBuffer {
	var chb channelBuffer
	chb.channels = make(map[string]*Channel)
	return &chb
}

func GetChannelById(id string) *Channel {
	chb.RLock()
	ch := chb.channels[id]
	chb.RUnlock()

	return ch
}
