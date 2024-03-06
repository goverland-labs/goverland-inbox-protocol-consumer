package feed

import (
	"context"
	"fmt"

	"github.com/goverland-labs/goverland-platform-events/events/inbox"
)

type Publisher interface {
	PublishJSON(ctx context.Context, subject string, obj any) error
}

type Service struct {
	events Publisher
}

func NewService(p Publisher) *Service {
	return &Service{
		events: p,
	}
}

func (s *Service) Handle(ctx context.Context, payload inbox.FeedPayload) error {
	if err := s.events.PublishJSON(ctx, inbox.SubjectFeedUpdated, payload); err != nil {
		return fmt.Errorf("publish feed event: %w", err)
	}

	return nil
}
