// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarHTTPAseTLSVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.TLSVersion.Get(
		context.TODO(),
		cloudflare.RadarHTTPAseTLSVersionGetParamsTLSVersionTlSv1_0,
		cloudflare.RadarHTTPAseTLSVersionGetParams{
			Asn:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsBotClass{cloudflare.RadarHTTPAseTLSVersionGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPAseTLSVersionGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsDateRange{cloudflare.RadarHTTPAseTLSVersionGetParamsDateRange1d, cloudflare.RadarHTTPAseTLSVersionGetParamsDateRange2d, cloudflare.RadarHTTPAseTLSVersionGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsDeviceType{cloudflare.RadarHTTPAseTLSVersionGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPAseTLSVersionGetParamsDeviceTypeMobile, cloudflare.RadarHTTPAseTLSVersionGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPAseTLSVersionGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPProtocol{cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPVersion{cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPAseTLSVersionGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsIPVersion{cloudflare.RadarHTTPAseTLSVersionGetParamsIPVersionIPv4, cloudflare.RadarHTTPAseTLSVersionGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			Os:           cloudflare.F([]cloudflare.RadarHTTPAseTLSVersionGetParamsO{cloudflare.RadarHTTPAseTLSVersionGetParamsOWindows, cloudflare.RadarHTTPAseTLSVersionGetParamsOMacosx, cloudflare.RadarHTTPAseTLSVersionGetParamsOIos}),
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
