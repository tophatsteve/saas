package iaas

import (
	"golang.org/x/net/context"
	"math/rand"
	"time"
)

type Service interface {
	Generate(ctx context.Context, in *Empty) (*GenerateResponse, error)
}

type server struct{
	messages []string
}

func NewService(m []string) Service {
	rand.Seed(time.Now().UnixNano())
	return server{messages: m}
}

func (s server) Generate(ctx context.Context, in *Empty) (*GenerateResponse, error) {
	resp := GenerateResponse{
		Value: s.generateMessage(),
	}
	return &resp, nil
}

func (s server)generateMessage() string {
	i := rand.Intn(len(s.messages))
	return s.messages[i]
}