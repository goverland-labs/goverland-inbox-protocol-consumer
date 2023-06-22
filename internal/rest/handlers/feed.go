package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/goverland-labs/inbox-protocol-consumer/internal/feed"
	"github.com/goverland-labs/inbox-protocol-consumer/internal/response"
	forms "github.com/goverland-labs/inbox-protocol-consumer/internal/rest/form/feed"
)

type FeedProcessor interface {
	Handle(_ context.Context, item feed.Item) error
}

type Feed struct {
	processor FeedProcessor
}

func NewFeedHandler(p FeedProcessor) APIHandler {
	return &Feed{
		processor: p,
	}
}

func (h *Feed) EnrichRoutes(baseRouter *mux.Router) {
	baseRouter.HandleFunc("/feed", h.feedCallback).Methods(http.MethodPost).Name("feed_callback")
}

func (h *Feed) feedCallback(w http.ResponseWriter, r *http.Request) {
	form, verr := forms.NewCallbackForm().ParseAndValidate(r)
	if verr != nil {
		response.HandleError(verr, w)

		return
	}

	params := form.(*forms.CallbackForm)
	err := h.processor.Handle(context.TODO(), feed.Item{
		DaoID:        params.DaoID,
		ProposalID:   params.ProposalID,
		DiscussionID: params.DiscussionID,
		Type:         params.Type,
		Action:       params.Action,
		Snapshot:     params.Snapshot,
	})

	if err != nil {
		log.Error().Err(err).Fields(params.ConvertToMap()).Msg("process feed item")

		response.HandleError(response.ResolveError(err), w)

		return
	}

	log.Debug().Fields(params.ConvertToMap()).Msg("processed feed item")

	w.WriteHeader(http.StatusOK)
}
