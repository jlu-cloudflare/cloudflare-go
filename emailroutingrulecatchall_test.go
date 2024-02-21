// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestEmailRoutingRuleCatchAllUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Emails.Routing.Rules.CatchAlls.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.EmailRoutingRuleCatchAllUpdateParams{
			Actions: cloudflare.F([]cloudflare.EmailRoutingRuleCatchAllUpdateParamsAction{{
				Type:  cloudflare.F(cloudflare.EmailRoutingRuleCatchAllUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.EmailRoutingRuleCatchAllUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}, {
				Type:  cloudflare.F(cloudflare.EmailRoutingRuleCatchAllUpdateParamsActionsTypeForward),
				Value: cloudflare.F([]string{"destinationaddress@example.net", "destinationaddress@example.net", "destinationaddress@example.net"}),
			}}),
			Matchers: cloudflare.F([]cloudflare.EmailRoutingRuleCatchAllUpdateParamsMatcher{{
				Type: cloudflare.F(cloudflare.EmailRoutingRuleCatchAllUpdateParamsMatchersTypeAll),
			}, {
				Type: cloudflare.F(cloudflare.EmailRoutingRuleCatchAllUpdateParamsMatchersTypeAll),
			}, {
				Type: cloudflare.F(cloudflare.EmailRoutingRuleCatchAllUpdateParamsMatchersTypeAll),
			}}),
			Enabled: cloudflare.F(cloudflare.EmailRoutingRuleCatchAllUpdateParamsEnabledTrue),
			Name:    cloudflare.F("Send to user@example.net rule."),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRoutingRuleCatchAllGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Emails.Routing.Rules.CatchAlls.Get(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
