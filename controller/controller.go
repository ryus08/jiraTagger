package controller

import (
	"encoding/json"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type Receive struct {
}

func (receive *Receive) Authorize(header http.Header, body *string) error {
	sv, err := slack.NewSecretsVerifier(header, "SigningSecret")
	if err != nil {
		return err
	}
	sv.Write([]byte(*body))
	err = sv.Ensure()
	return err
}

func (receive *Receive) Handler(body *string) (slackevents.EventsAPIEvent, error) {
	eventsAPIEvent, e := slackevents.ParseEvent(
		json.RawMessage([]byte(*body)),
		slackevents.OptionNoVerifyToken())

	// TODO:
	// Config pass in
	// Handle Challenge again

	return eventsAPIEvent, e
}
