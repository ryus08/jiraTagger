package controller

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ryus08/jiraTagger/config"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type TaggerController struct {
	Config *config.Config
}

func (taggerController *TaggerController) Authorize(header http.Header, body *string) error {
	sv, err := slack.NewSecretsVerifier(header, taggerController.Config.SigningSecret)
	if err != nil {
		return err
	}
	sv.Write([]byte(*body))
	err = sv.Ensure()
	return err
}

func (taggerController *TaggerController) Handle(request *http.Request) (int, interface{}) {
	var body string
	if request.Method == "GET" {
		body = "{Content: \"Hello!\"}"
	} else {
		buf := new(bytes.Buffer)
		buf.ReadFrom(request.Body)
		body = buf.String()
	}

	// TODO: Probably need to break this up so API Gateway can run them as independent lambdas
	e := taggerController.Authorize(request.Header, &body)
	if e != nil {
		return http.StatusUnauthorized, e
	}

	response, e := taggerController.HandleRequest(&body)
	if e != nil {
		return http.StatusInternalServerError, e
	}

	return http.StatusOK, response
}

func (taggerController *TaggerController) HandleRequest(body *string) (slackevents.EventsAPIEvent, error) {
	eventsAPIEvent, e := slackevents.ParseEvent(
		json.RawMessage([]byte(*body)),
		slackevents.OptionNoVerifyToken())

	// TODO:
	// Config pass in
	// Handle Challenge again

	return eventsAPIEvent, e
}
