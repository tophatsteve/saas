package iaas

import (
	"golang.org/x/net/context"
)

type Service interface {
	Generate(ctx context.Context, in *Empty) (*GenerateResponse, error)
}

type server struct{}

func NewService() Service {
	return server{}
}

func (s server) Generate(ctx context.Context, in *Empty) (*GenerateResponse, error) {
	resp := GenerateResponse{
		Value: generateMessage(),
	}
	return &resp, nil
}

func generateMessage() string {
	return "Your mother was a hamster, and your father smelt of elderberries."
}