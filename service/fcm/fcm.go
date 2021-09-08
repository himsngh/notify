package fcm

import (
	"context"
	"github.com/appleboy/go-fcm"
	"github.com/pkg/errors"
)

type fcmClient interface {
	Send(msg *fcm.Message) (*fcm.Response, error)
}

type Service struct {
	client       fcmClient
	deviceTokens []string
}

func NewService(apiKey string) (*Service, error) {
	cli, err := fcm.NewClient(apiKey)
	if err != nil {
		return nil, err
	}

	return &Service{
		client: cli,
	}, nil
}

func (s *Service) AddReceivers(deviceToken ...string) {
	s.deviceTokens = append(s.deviceTokens, deviceToken...)
}

func (s *Service) Send(ctx context.Context, subject, message string) error {

	for _, token := range s.deviceTokens {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			res, err := s.client.Send(&fcm.Message{
				To: token,
				Notification: &fcm.Notification{
					Title: subject,
					Body:  message,
				},
				Apns: map[string]interface{}{},
			})
			if err != nil || res.Error != nil {
				return errors.Wrapf(err, "failed to send message to WhatsApp contact '%s'", token)
			}
		}
	}
	return nil
}
