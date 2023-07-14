package feed

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// todo: think about location
const (
	ActionCreated             = "created"
	ActionUpdated             = "updated"
	ActionUpdatedState        = "updated.state"
	ActionVotingStarted       = "voting.started"
	ActionVotingEnded         = "voting.ended"
	ActionVotingQuorumReached = "voting.quorum_reached"
	ActionVotingStartsSoon    = "voting.starts_soon"
)

const (
	TypeDao      = "dao"
	TypeProposal = "proposal"

	DaoCreated                  = "dao.created"
	DaoUpdated                  = "dao.updated"
	ProposalCreated             = "proposal.created"
	ProposalUpdated             = "proposal.updated"
	ProposalVotingStartsSoon    = "proposal.voting.starts_soon"
	ProposalVotingStarted       = "proposal.voting.started"
	ProposalVotingQuorumReached = "proposal.voting.quorum_reached"
	ProposalVotingEnded         = "proposal.voting.ended"
)

type Item struct {
	ID           uuid.UUID `json:"id"`
	DaoID        uuid.UUID `json:"dao_id"`
	ProposalID   string    `json:"proposal_id"`
	DiscussionID string    `json:"discussion_id"`
	Type         string    `json:"type"`
	Action       string    `json:"action"`

	Snapshot json.RawMessage `json:"snapshot"`
	Timeline []TimelineItem  `json:"timeline"`
}

type TimelineItem struct {
	CreatedAt time.Time `json:"created_at"`
	Action    string    `json:"action"`
}
