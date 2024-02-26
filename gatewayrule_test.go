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

func TestGatewayRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Gateways.Rules.New(context.TODO(), cloudflare.GatewayRuleNewParams{
		AccountID:     cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
		Action:        cloudflare.F(cloudflare.GatewayRuleNewParamsActionAllow),
		Name:          cloudflare.F("block bad websites"),
		Description:   cloudflare.F("Block bad websites based on their host name."),
		DevicePosture: cloudflare.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
		Enabled:       cloudflare.F(true),
		Filters:       cloudflare.F([]cloudflare.GatewayRuleNewParamsFilter{cloudflare.GatewayRuleNewParamsFilterHTTP}),
		Identity:      cloudflare.F("any(identity.groups.name[*] in {\"finance\"})"),
		Precedence:    cloudflare.F(int64(0)),
		RuleSettings: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettings{
			AddHeaders: cloudflare.F[any](map[string]interface{}{
				"My-Next-Header": map[string]interface{}{
					"0": "foo",
					"1": "bar",
				},
				"X-Custom-Header-Name": map[string]interface{}{
					"0": "somecustomvalue",
				},
			}),
			AllowChildBypass: cloudflare.F(false),
			AuditSSH: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettingsAuditSSH{
				CommandLogging: cloudflare.F(false),
			}),
			BisoAdminControls: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettingsBisoAdminControls{
				Dcp: cloudflare.F(false),
				Dd:  cloudflare.F(false),
				Dk:  cloudflare.F(false),
				Dp:  cloudflare.F(false),
				Du:  cloudflare.F(false),
			}),
			BlockPageEnabled: cloudflare.F(true),
			BlockReason:      cloudflare.F("This website is a security risk"),
			BypassParentRule: cloudflare.F(false),
			CheckSession: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettingsCheckSession{
				Duration: cloudflare.F("300s"),
				Enforce:  cloudflare.F(true),
			}),
			DNSResolvers: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettingsDNSResolvers{
				IPV4: cloudflare.F([]cloudflare.GatewayRuleNewParamsRuleSettingsDNSResolversIPV4{{
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}}),
				IPV6: cloudflare.F([]cloudflare.GatewayRuleNewParamsRuleSettingsDNSResolversIPV6{{
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}, {
					IP:                         cloudflare.F("2001:DB8::/64"),
					Port:                       cloudflare.F(int64(5053)),
					RouteThroughPrivateNetwork: cloudflare.F(true),
					VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				}}),
			}),
			Egress: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettingsEgress{
				IPV4:         cloudflare.F("192.0.2.2"),
				IPV4Fallback: cloudflare.F("192.0.2.3"),
				IPV6:         cloudflare.F("2001:DB8::/64"),
			}),
			InsecureDisableDNSSECValidation: cloudflare.F(false),
			IPCategories:                    cloudflare.F(true),
			IPIndicatorFeeds:                cloudflare.F(true),
			L4override: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettingsL4override{
				IP:   cloudflare.F("1.1.1.1"),
				Port: cloudflare.F(int64(0)),
			}),
			NotificationSettings: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettingsNotificationSettings{
				Enabled:    cloudflare.F(true),
				Msg:        cloudflare.F("string"),
				SupportURL: cloudflare.F("string"),
			}),
			OverrideHost: cloudflare.F("example.com"),
			OverrideIPs:  cloudflare.F([]string{"1.1.1.1", "2.2.2.2"}),
			PayloadLog: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettingsPayloadLog{
				Enabled: cloudflare.F(true),
			}),
			ResolveDNSThroughCloudflare: cloudflare.F(true),
			UntrustedCert: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettingsUntrustedCert{
				Action: cloudflare.F(cloudflare.GatewayRuleNewParamsRuleSettingsUntrustedCertActionError),
			}),
		}),
		Schedule: cloudflare.F(cloudflare.GatewayRuleNewParamsSchedule{
			Fri:      cloudflare.F("08:00-12:30,13:30-17:00"),
			Mon:      cloudflare.F("08:00-12:30,13:30-17:00"),
			Sat:      cloudflare.F("08:00-12:30,13:30-17:00"),
			Sun:      cloudflare.F("08:00-12:30,13:30-17:00"),
			Thu:      cloudflare.F("08:00-12:30,13:30-17:00"),
			TimeZone: cloudflare.F("America/New York"),
			Tue:      cloudflare.F("08:00-12:30,13:30-17:00"),
			Wed:      cloudflare.F("08:00-12:30,13:30-17:00"),
		}),
		Traffic: cloudflare.F("http.request.uri matches \".*a/partial/uri.*\" and http.request.host in $01302951-49f9-47c9-a400-0297e60b6a10"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Gateways.Rules.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.GatewayRuleUpdateParams{
			AccountID:     cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
			Action:        cloudflare.F(cloudflare.GatewayRuleUpdateParamsActionAllow),
			Name:          cloudflare.F("block bad websites"),
			Description:   cloudflare.F("Block bad websites based on their host name."),
			DevicePosture: cloudflare.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
			Enabled:       cloudflare.F(true),
			Filters:       cloudflare.F([]cloudflare.GatewayRuleUpdateParamsFilter{cloudflare.GatewayRuleUpdateParamsFilterHTTP}),
			Identity:      cloudflare.F("any(identity.groups.name[*] in {\"finance\"})"),
			Precedence:    cloudflare.F(int64(0)),
			RuleSettings: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettings{
				AddHeaders: cloudflare.F[any](map[string]interface{}{
					"My-Next-Header": map[string]interface{}{
						"0": "foo",
						"1": "bar",
					},
					"X-Custom-Header-Name": map[string]interface{}{
						"0": "somecustomvalue",
					},
				}),
				AllowChildBypass: cloudflare.F(false),
				AuditSSH: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettingsAuditSSH{
					CommandLogging: cloudflare.F(false),
				}),
				BisoAdminControls: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettingsBisoAdminControls{
					Dcp: cloudflare.F(false),
					Dd:  cloudflare.F(false),
					Dk:  cloudflare.F(false),
					Dp:  cloudflare.F(false),
					Du:  cloudflare.F(false),
				}),
				BlockPageEnabled: cloudflare.F(true),
				BlockReason:      cloudflare.F("This website is a security risk"),
				BypassParentRule: cloudflare.F(false),
				CheckSession: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettingsCheckSession{
					Duration: cloudflare.F("300s"),
					Enforce:  cloudflare.F(true),
				}),
				DNSResolvers: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettingsDNSResolvers{
					IPV4: cloudflare.F([]cloudflare.GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4{{
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
					IPV6: cloudflare.F([]cloudflare.GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6{{
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::/64"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
				}),
				Egress: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettingsEgress{
					IPV4:         cloudflare.F("192.0.2.2"),
					IPV4Fallback: cloudflare.F("192.0.2.3"),
					IPV6:         cloudflare.F("2001:DB8::/64"),
				}),
				InsecureDisableDNSSECValidation: cloudflare.F(false),
				IPCategories:                    cloudflare.F(true),
				IPIndicatorFeeds:                cloudflare.F(true),
				L4override: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettingsL4override{
					IP:   cloudflare.F("1.1.1.1"),
					Port: cloudflare.F(int64(0)),
				}),
				NotificationSettings: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettingsNotificationSettings{
					Enabled:    cloudflare.F(true),
					Msg:        cloudflare.F("string"),
					SupportURL: cloudflare.F("string"),
				}),
				OverrideHost: cloudflare.F("example.com"),
				OverrideIPs:  cloudflare.F([]string{"1.1.1.1", "2.2.2.2"}),
				PayloadLog: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettingsPayloadLog{
					Enabled: cloudflare.F(true),
				}),
				ResolveDNSThroughCloudflare: cloudflare.F(true),
				UntrustedCert: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettingsUntrustedCert{
					Action: cloudflare.F(cloudflare.GatewayRuleUpdateParamsRuleSettingsUntrustedCertActionError),
				}),
			}),
			Schedule: cloudflare.F(cloudflare.GatewayRuleUpdateParamsSchedule{
				Fri:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Mon:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Sat:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Sun:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Thu:      cloudflare.F("08:00-12:30,13:30-17:00"),
				TimeZone: cloudflare.F("America/New York"),
				Tue:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Wed:      cloudflare.F("08:00-12:30,13:30-17:00"),
			}),
			Traffic: cloudflare.F("http.request.uri matches \".*a/partial/uri.*\" and http.request.host in $01302951-49f9-47c9-a400-0297e60b6a10"),
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

func TestGatewayRuleList(t *testing.T) {
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
	_, err := client.Gateways.Rules.List(context.TODO(), cloudflare.GatewayRuleListParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayRuleDelete(t *testing.T) {
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
	_, err := client.Gateways.Rules.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.GatewayRuleDeleteParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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

func TestGatewayRuleGet(t *testing.T) {
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
	_, err := client.Gateways.Rules.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.GatewayRuleGetParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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
