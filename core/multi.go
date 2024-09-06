package core

import (
	"fmt"
	"time"
)

type CTsStream[T CtStream] struct {
	Streams []T
	Sleep   time.Duration
}

func NewCTsStream[T CtStream](streams []T, sleep time.Duration) (*CTsStream[T], error) {
	return &CTsStream[T]{
		Streams: streams,
		Sleep:   sleep,
	}, nil
}

func (stream *CTsStream[T]) Init() error {
	for _, s := range stream.Streams {
		err := s.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (stream *CTsStream[T]) Start(callback Callback) {
	for _, s := range stream.Streams {
		go s.Run(callback)
	}
}

func (stream CTsStream[T]) Await() {
	for _, s := range stream.Streams {
		time.Sleep(stream.Sleep)
		s.Await()
	}
}

func (stream *CTsStream[T]) Run(callback Callback) {
	stream.Start(callback)
	stream.Await()
}

func (stream *CTsStream[T]) Stop() {
	for _, s := range stream.Streams {
		fmt.Printf("1\n")
		s.Stop()
	}
}
