package core

import (
	"time"
)

type SyncCTsStream[T CtStream] struct {
	Streams []T
	Sleep   time.Duration
}

func NewSyncCTsStream[T CtStream](streams []T, sleep time.Duration) (*SyncCTsStream[T], error) {
	return &SyncCTsStream[T]{
		Streams: streams,
		Sleep:   sleep,
	}, nil
}

func (stream *SyncCTsStream[T]) Init() error {
	for _, s := range stream.Streams {
		err := s.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (stream *SyncCTsStream[T]) Start(callback Callback) {
	for {
		stream.Next(callback)
	}
}

func (stream *SyncCTsStream[T]) Next(callback Callback) {
	for _, s := range stream.Streams {
		s.Next(callback)
		time.Sleep(stream.Sleep)
	}
}

func (stream SyncCTsStream[T]) Await() {
	for _, s := range stream.Streams {
		time.Sleep(stream.Sleep)
		s.Await()
	}
}

func (stream *SyncCTsStream[T]) Run(callback Callback) {
	stream.Start(callback)
	stream.Await()
}

func (stream *SyncCTsStream[T]) Stop() {
	for _, s := range stream.Streams {
		s.Stop()
	}
}
