package feed

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/goverland-labs/inbox-protocol-consumer/internal/response"
	"github.com/goverland-labs/inbox-protocol-consumer/internal/response/errs"
	"github.com/goverland-labs/inbox-protocol-consumer/internal/rest/form"
)

type CallBackRequest struct {
	DaoID        uuid.UUID `json:"dao_id"`
	ProposalID   string    `json:"proposal_id"`
	DiscussionID string    `json:"discussion_id"`
	Type         string    `json:"type"` // dao || proposal
	Action       string    `json:"action"`

	Snapshot json.RawMessage `json:"snapshot"`
}

type CallbackForm struct {
	DaoID        uuid.UUID
	ProposalID   string
	DiscussionID string
	Type         string
	Action       string

	Snapshot json.RawMessage `json:"snapshot"`
}

func NewCallbackForm() *CallbackForm {
	return &CallbackForm{}
}

func (f *CallbackForm) ParseAndValidate(r *http.Request) (form.Former, response.Error) {
	var req *CallBackRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ve := response.NewValidationError()
		ve.SetError(response.GeneralErrorKey, errs.InvalidRequestStructure, "invalid request structure")

		return nil, ve
	}

	f.DaoID = req.DaoID
	f.ProposalID = req.ProposalID
	f.DiscussionID = req.DiscussionID
	f.Type = req.Type
	f.Action = req.Action
	f.Snapshot = req.Snapshot

	return f, nil
}

func (f *CallbackForm) ConvertToMap() map[string]interface{} {
	return map[string]interface{}{
		"dao_id":        f.DaoID,
		"proposal_id":   f.ProposalID,
		"discussion_id": f.DiscussionID,
		"type":          f.Type,
		"action":        f.Action,
	}
}
