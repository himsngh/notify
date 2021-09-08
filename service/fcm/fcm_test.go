package fcm

import (
	"context"
	"github.com/appleboy/go-fcm"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddReceivers(t *testing.T) {
	assert := require.New(t)

	svc := &Service{
		deviceTokens: []string{},
	}
	tokens := []string{"Token1", "Token2", "Token3"}
	svc.AddReceivers(tokens...)

	assert.Equal(svc.deviceTokens, tokens)
}

func TestSend(t *testing.T) {
	assert := require.New(t)

	svc := &Service{
		deviceTokens: nil,
	}

	mockClient := new(mockFcmClient)
	mockClient.On("Send", &fcm.Message{
		To: "Token1",
		Notification: &fcm.Notification{
			Title: "Subject",
			Body:  "Body",
		},
	}).Return(nil, errors.New("some error"))
	svc.client = mockClient
	svc.AddReceivers("Token1")
	ctx := context.Background()
	err := svc.Send(ctx, "Subject", "Body")
	assert.NotNil(err)
	mockClient.AssertExpectations(t)
}
