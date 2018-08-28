package lorem_grpc

import (
	"context"
	"errors"
	"strings"

	gl "github.com/drhodes/golorem"
)

var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence or paragraph.")
)

// Service interface
type Service interface {
	// generate word with at least min letters and at most max letters
	Lorem(ctx context.Context, requestType string, min, max int) (string, error)
}

// LoremService struct
type LoremService struct {
}

// Lorem Service function
func (LoremService) Lorem(_ context.Context, requestType string, min, max int) (string, error) {
	var result string
	var err error
	if strings.EqualFold(requestType, "Word") {
		result = gl.Word(min, max)
	} else if strings.EqualFold(requestType, "Sentence") {
		result = gl.Sentence(min, max)
	} else if strings.EqualFold(requestType, "Paragraph") {
		result = gl.Paragraph(min, max)
	} else {
		err = ErrRequestTypeNotFound
	}
	return result, err
}
