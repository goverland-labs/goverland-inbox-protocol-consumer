package feed

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/goverland-labs/platform-events/events/inbox"
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

func (s *Service) Handle(ctx context.Context, item Item) error {
	switch item.Type {
	case TypeDao:
		return s.handleDao(ctx, item)
	case TypeProposal:
		return s.handleProposal(ctx, item)
	default:
		return errors.New("unsupported type")
	}
}

func (s *Service) handleDao(ctx context.Context, item Item) error {
	var dao Dao
	err := json.Unmarshal(item.Snapshot, &dao)
	if err != nil {
		return fmt.Errorf("unmarshal snapshot to dao object: %w", err)
	}

	subject, err := getDaoSubject(item.Action)
	if err != nil {
		return fmt.Errorf("get subject: %w", err)
	}

	if err = s.events.PublishJSON(ctx, subject, convertDaoToInbox(dao)); err != nil {
		return fmt.Errorf("publish dao event: %w", err)
	}

	return nil
}

func getDaoSubject(action string) (string, error) {
	switch action {
	case ActionCreated:
		return inbox.SubjectDaoCreated, nil
	case ActionUpdated:
		return inbox.SubjectDaoUpdated, nil
	default:
		return "", errors.New("unsupported action")
	}
}

func convertDaoToInbox(dao Dao) inbox.DaoPayload {
	return inbox.DaoPayload{
		ID:         dao.ID,
		Name:       dao.Name,
		Private:    dao.Private,
		About:      dao.About,
		Avatar:     dao.Avatar,
		Terms:      dao.Terms,
		Location:   dao.Location,
		Website:    dao.Website,
		Twitter:    dao.Twitter,
		Github:     dao.Github,
		Coingecko:  dao.Coingecko,
		Email:      dao.Email,
		Network:    dao.Network,
		Symbol:     dao.Symbol,
		Skin:       dao.Skin,
		Domain:     dao.Domain,
		Strategies: convertStrategiesToInbox(dao.Strategies),
		Admins:     dao.Admins,
		Members:    dao.Members,
		Moderators: dao.Moderators,
		Voting: inbox.VotingPayload{
			Delay:       dao.Voting.Delay,
			Period:      dao.Voting.Period,
			Type:        dao.Voting.Type,
			Quorum:      dao.Voting.Quorum,
			Blind:       dao.Voting.Blind,
			HideAbstain: dao.Voting.HideAbstain,
			Privacy:     dao.Voting.Privacy,
			Aliased:     dao.Voting.Aliased,
		},
		Categories:     dao.Categories,
		Treasures:      convertTreasuresToInbox(dao.Treasures),
		FollowersCount: dao.FollowersCount,
		ProposalsCount: dao.ProposalsCount,
		Guidelines:     dao.Guidelines,
		Template:       dao.Template,
		ParentID:       dao.ParentID,
	}
}

func convertStrategiesToInbox(data []Strategy) []inbox.StrategyPayload {
	res := make([]inbox.StrategyPayload, len(data))
	for i, info := range data {
		res[i] = inbox.StrategyPayload{
			Name:    info.Name,
			Network: info.Network,
		}
	}

	return res
}

func convertTreasuresToInbox(data []Treasury) []inbox.TreasuryPayload {
	res := make([]inbox.TreasuryPayload, len(data))
	for i, info := range data {
		res[i] = inbox.TreasuryPayload{
			Name:    info.Name,
			Address: info.Address,
			Network: info.Network,
		}
	}

	return res
}

func (s *Service) handleProposal(ctx context.Context, item Item) error {
	var pd Proposal
	err := json.Unmarshal(item.Snapshot, &pd)
	if err != nil {
		return fmt.Errorf("unmarshal snapshot to proposal object: %w", err)
	}

	subject, err := getProposalSubject(item.Action)
	if err != nil {
		return fmt.Errorf("get subject: %w", err)
	}

	if err = s.events.PublishJSON(ctx, subject, convertProposalToInbox(pd)); err != nil {
		return fmt.Errorf("publish proposal event: %w", err)
	}

	return nil
}

func convertProposalToInbox(data Proposal) inbox.ProposalPayload {
	return inbox.ProposalPayload{
		ID:            data.ID,
		Ipfs:          data.Ipfs,
		Author:        data.Author,
		Created:       data.Created,
		DaoID:         data.DaoID,
		Network:       data.Network,
		Symbol:        data.Symbol,
		Type:          data.Type,
		Strategies:    convertStrategiesToInbox(data.Strategies),
		Title:         data.Title,
		Body:          data.Body,
		Discussion:    data.Discussion,
		Choices:       data.Choices,
		Start:         data.Start,
		End:           data.End,
		Quorum:        data.Quorum,
		Privacy:       data.Privacy,
		Snapshot:      data.Snapshot,
		State:         data.State,
		Link:          data.Link,
		App:           data.App,
		Scores:        data.Scores,
		ScoresState:   data.ScoresState,
		ScoresTotal:   data.ScoresTotal,
		ScoresUpdated: data.ScoresUpdated,
		Votes:         data.Votes,
	}
}

func getProposalSubject(action string) (string, error) {
	switch action {
	case ActionCreated:
		return inbox.SubjectProposalCreated, nil
	case ActionUpdated:
		return inbox.SubjectProposalUpdated, nil
	case ActionUpdatedState:
		return inbox.SubjectProposalUpdatedState, nil
	case ActionVotingStartsSoon:
		return inbox.SubjectProposalVotingStartsSoon, nil
	case ActionVotingStarted:
		return inbox.SubjectProposalVotingStarted, nil
	case ActionVotingQuorumReached:
		return inbox.SubjectProposalVotingQuorumReached, nil
	case ActionVotingEnded:
		return inbox.SubjectProposalVotingEnded, nil
	default:
		return "", errors.New("unsupported action")
	}
}
