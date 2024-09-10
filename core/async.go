package core

import (
	"time"
)

type AsyncCTsStream[T CtStream] struct {
	Streams []T
	Sleep   time.Duration
}

func NewAsyncCTsStream[T CtStream](streams []T, sleep time.Duration) (*AsyncCTsStream[T], error) {
	return &AsyncCTsStream[T]{
		Streams: streams,
		Sleep:   sleep,
	}, nil
}

func (stream *AsyncCTsStream[T]) Init() error {
	for _, s := range stream.Streams {
		err := s.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (stream *AsyncCTsStream[T]) Start(callback Callback) {
	for _, s := range stream.Streams {
		go s.Run(callback)
	}
}

func (stream *AsyncCTsStream[T]) Next(callback Callback) {
}

func (stream AsyncCTsStream[T]) Await() {
	for _, s := range stream.Streams {
		time.Sleep(stream.Sleep)
		s.Await()
	}
}

func (stream *AsyncCTsStream[T]) Run(callback Callback) {
	stream.Start(callback)
	stream.Await()
}

func (stream *AsyncCTsStream[T]) Stop() {
	for _, s := range stream.Streams {
		s.Stop()
	}
}
