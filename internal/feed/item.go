package feed

import "encoding/json"

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
)

type Item struct {
	DaoID        string `json:"dao_id"`
	ProposalID   string `json:"proposal_id"`
	DiscussionID string `json:"discussion_id"`
	Type         string `json:"type"`
	Action       string `json:"action"`

	Snapshot json.RawMessage `json:"snapshot"`
}
