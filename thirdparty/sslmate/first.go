package sslmate

import (
	"strconv"

	"github.com/akakou/ctstream/core"
)

func GetFirst(CTsStream *core.ConcurrentCTsStream[*core.CTStream[*SSLMateCTClient]]) int {
	var first int

	for _, stream := range CTsStream.Streams {
		tmp, _ := strconv.Atoi(stream.Client.First)

		if tmp > first {
			first = tmp
		}
	}

	return first
}

func SetFirst(first int, CTsStream *core.ConcurrentCTsStream[*core.CTStream[*SSLMateCTClient]]) {
	for _, stream := range CTsStream.Streams {
		stream.Client.First = strconv.Itoa(first)
	}
}
