package core

import (
	"time"
)

type ConcurrentCTsStream[T CtStream] struct {
	Streams []T
	Sleep   time.Duration
}

func NewConcurrentCTsStream[T CtStream](streams []T, sleep time.Duration) (*ConcurrentCTsStream[T], error) {
	return &ConcurrentCTsStream[T]{
		Streams: streams,
		Sleep:   sleep,
	}, nil
}

func (stream *ConcurrentCTsStream[T]) Init() error {
	for _, s := range stream.Streams {
		err := s.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (stream *ConcurrentCTsStream[T]) Start(callback Callback) {
	go stream.start(callback)
}

func (stream *ConcurrentCTsStream[T]) start(callback Callback) {
	for {
		stream.Next(callback)
	}
}

func (stream *ConcurrentCTsStream[T]) Next(callback Callback) {
	for _, s := range stream.Streams {
		go s.Next(callback)
		time.Sleep(stream.Sleep)
	}
}

func (stream ConcurrentCTsStream[T]) Await() {
	for _, s := range stream.Streams {
		time.Sleep(stream.Sleep)
		s.Await()
	}
}

func (stream *ConcurrentCTsStream[T]) Run(callback Callback) {
	stream.Start(callback)
	stream.Await()
}

func (stream *ConcurrentCTsStream[T]) Stop() {
	for _, s := range stream.Streams {
		s.Stop()
	}
}
