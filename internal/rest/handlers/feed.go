package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/goverland-labs/goverland-platform-events/events/inbox"
	"github.com/rs/zerolog/log"

	"github.com/goverland-labs/inbox-protocol-consumer/internal/response"
	forms "github.com/goverland-labs/inbox-protocol-consumer/internal/rest/form/feed"
)

type FeedProcessor interface {
	Handle(_ context.Context, item inbox.FeedPayload) error
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
	baseRouter.HandleFunc("/webhook", h.feedWebhook).Methods(http.MethodPost).Name("feed_callback")
}

func (h *Feed) feedWebhook(w http.ResponseWriter, r *http.Request) {
	form, verr := forms.NewCallbackForm().ParseAndValidate(r)
	if verr != nil {
		response.HandleError(verr, w)

		return
	}

	params := form.(*forms.CallbackForm)
	err := h.processor.Handle(context.TODO(), params.Request)

	if err != nil {
		log.Error().Err(err).Fields(params.ConvertToMap()).Msg("process feed item")

		response.HandleError(response.ResolveError(err), w)

		return
	}

	log.Debug().Fields(params.ConvertToMap()).Msg("processed feed item")

	w.WriteHeader(http.StatusOK)
}
