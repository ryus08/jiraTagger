package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	Convey("Start Testing!", t, func() {

		Convey("Validating Lambda Authorizes!", func() {
			raw, err := ioutil.ReadFile("./event.json")

			So(raw, ShouldNotBeNil)
			So(err, ShouldBeNil)

			req := &events.APIGatewayProxyRequest{}
			err = json.Unmarshal(raw, req)

			So(err, ShouldBeNil)

			resp, err := Handler(context.Background(), req)

			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 401)
		})

	})
}
