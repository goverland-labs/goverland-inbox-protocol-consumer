package feed

import (
	"encoding/json"
	"net/http"

	"github.com/goverland-labs/goverland-platform-events/events/inbox"

	"github.com/goverland-labs/goverland-inbox-protocol-consumer/internal/response"
	"github.com/goverland-labs/goverland-inbox-protocol-consumer/internal/response/errs"
	"github.com/goverland-labs/goverland-inbox-protocol-consumer/internal/rest/form"
)

type WebhookRequest = inbox.FeedPayload

type CallbackForm struct {
	Request WebhookRequest
}

func NewCallbackForm() *CallbackForm {
	return &CallbackForm{}
}

func (f *CallbackForm) ParseAndValidate(r *http.Request) (form.Former, response.Error) {
	var req WebhookRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ve := response.NewValidationError()
		ve.SetError(response.GeneralErrorKey, errs.InvalidRequestStructure, "invalid request structure")

		return nil, ve
	}

	f.Request = req

	return f, nil
}

func (f *CallbackForm) ConvertToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            f.Request.ID,
		"dao_id":        f.Request.DaoID,
		"proposal_id":   f.Request.ProposalID,
		"discussion_id": f.Request.DiscussionID,
		"type":          f.Request.Type,
		"action":        f.Request.Action,
	}
}
