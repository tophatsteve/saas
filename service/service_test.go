package saas

import (
	"testing"
	"github.com/stretchr/testify/assert"	
	"golang.org/x/net/context"
)

func TestGenerateReturnsExpectedMessage(t *testing.T) {
	s := NewService([]string{"test message"})
	m, _ := s.Generate(context.Background(), &Empty{})
	assert.Equal(t, 
		"test message",
		m.Value, 
		"received message is not correct",
	)
}