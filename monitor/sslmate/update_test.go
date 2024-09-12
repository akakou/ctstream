package sslmate

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	domain := "test.example.com"

	stream, _ := DefaultCTsStream([]string{}, context.Background())

	t.Run("TestAdd", func(t *testing.T) {
		const domain = "test.example.com"
		_, i, err := AddByDomain(domain, context.Background(), stream)
		assert.NoError(t, err)
		assert.Equal(t, i, 0)
		assert.Equal(t, len(stream.Streams), 1)
	})

	t.Run("TestDel", func(t *testing.T) {
		_, i, err := DelByDomain(domain, context.Background(), stream)
		assert.NoError(t, err)
		assert.Equal(t, i, 0)
		assert.Equal(t, len(stream.Streams), 0)
	})

}
