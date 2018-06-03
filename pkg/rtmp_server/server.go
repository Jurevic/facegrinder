package rtmp_server

import (
	"github.com/jurevic/facegrinder/pkg/api/v1/channel"
	"github.com/nareix/joy4/format"
	"github.com/nareix/joy4/format/rtmp"
)

func Init() {
	format.RegisterAll()

	go startRtmpServer()
}

func startRtmpServer() {
	server := &rtmp.Server{}

	server.HandlePlay = channel.RtmpPlayHandler()
	server.HandlePublish = channel.RtmpPublishHandler()

	server.ListenAndServe()
}
