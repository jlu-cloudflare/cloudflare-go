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

func TestAccessOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Access.Organizations.New(context.TODO(), cloudflare.AccessOrganizationNewParams{
		AccountID:                cloudflare.F("string"),
		ZoneID:                   cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		AuthDomain:               cloudflare.F("test.cloudflareaccess.com"),
		Name:                     cloudflare.F("Widget Corps Internal Applications"),
		AllowAuthenticateViaWarp: cloudflare.F(true),
		AutoRedirectToIdentity:   cloudflare.F(true),
		IsUiReadOnly:             cloudflare.F(true),
		LoginDesign: cloudflare.F(cloudflare.AccessOrganizationNewParamsLoginDesign{
			BackgroundColor: cloudflare.F("#c5ed1b"),
			FooterText:      cloudflare.F("This is an example description."),
			HeaderText:      cloudflare.F("This is an example description."),
			LogoPath:        cloudflare.F("https://example.com/logo.png"),
			TextColor:       cloudflare.F("#c5ed1b"),
		}),
		SessionDuration:                cloudflare.F("24h"),
		UiReadOnlyToggleReason:         cloudflare.F("Temporarily turn off the UI read only lock to make a change via the UI"),
		UserSeatExpirationInactiveTime: cloudflare.F("720h"),
		WarpAuthSessionDuration:        cloudflare.F("24h"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessOrganizationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Access.Organizations.Update(context.TODO(), cloudflare.AccessOrganizationUpdateParams{
		AccountID:                cloudflare.F("string"),
		ZoneID:                   cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		AllowAuthenticateViaWarp: cloudflare.F(true),
		AuthDomain:               cloudflare.F("test.cloudflareaccess.com"),
		AutoRedirectToIdentity:   cloudflare.F(true),
		CustomPages: cloudflare.F(cloudflare.AccessOrganizationUpdateParamsCustomPages{
			Forbidden:      cloudflare.F("699d98642c564d2e855e9661899b7252"),
			IdentityDenied: cloudflare.F("699d98642c564d2e855e9661899b7252"),
		}),
		IsUiReadOnly: cloudflare.F(true),
		LoginDesign: cloudflare.F(cloudflare.AccessOrganizationUpdateParamsLoginDesign{
			BackgroundColor: cloudflare.F("#c5ed1b"),
			FooterText:      cloudflare.F("This is an example description."),
			HeaderText:      cloudflare.F("This is an example description."),
			LogoPath:        cloudflare.F("https://example.com/logo.png"),
			TextColor:       cloudflare.F("#c5ed1b"),
		}),
		Name:                           cloudflare.F("Widget Corps Internal Applications"),
		SessionDuration:                cloudflare.F("24h"),
		UiReadOnlyToggleReason:         cloudflare.F("Temporarily turn off the UI read only lock to make a change via the UI"),
		UserSeatExpirationInactiveTime: cloudflare.F("720h"),
		WarpAuthSessionDuration:        cloudflare.F("24h"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessOrganizationList(t *testing.T) {
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
	_, err := client.Access.Organizations.List(context.TODO(), cloudflare.AccessOrganizationListParams{
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessOrganizationRevokeUsers(t *testing.T) {
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
	_, err := client.Access.Organizations.RevokeUsers(context.TODO(), cloudflare.AccessOrganizationRevokeUsersParams{
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Email:     cloudflare.F("test@example.com"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
