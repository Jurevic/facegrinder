package channel

import (
	"fmt"
	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/rtmp"
	"strings"
)

func RtmpPublishHandler() func(conn *rtmp.Conn) {
	return func(conn *rtmp.Conn) {
		streams, _ := conn.Streams()

		id := strings.Trim(conn.URL.Path, "/")
		key := conn.URL.Query().Get("key")

		chb.Lock()
		ch := chb.channels[id]
		chb.Unlock()

		if ch == nil {
			fmt.Println("Specified channel does not exist")
			return
		}

		if key != ch.Key {
			fmt.Println("Channel key does not match, stream denied")
			return
		}

		ch.Que = pubsub.NewQueue()
		ch.Que.WriteHeader(streams)

		// Stream
		avutil.CopyPackets(ch.Que, conn)

		// On closed stream close channel
		chb.Lock()
		delete(chb.channels, id)
		chb.Unlock()

		ch.Que.Close()
	}
}

func RtmpPlayHandler() func(conn *rtmp.Conn) {
	return func(conn *rtmp.Conn) {
		id := strings.Trim(conn.URL.Path, "/")

		chb.RLock()
		ch := chb.channels[id]
		chb.RUnlock()

		if ch != nil {
			cursor := ch.Que.Latest()
			err := avutil.CopyFile(conn, cursor)
			fmt.Println(err)
		}
	}
}
