// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// GatewayService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewGatewayService] method instead.
type GatewayService struct {
	Options          []option.RequestOption
	AuditSSHSettings *GatewayAuditSSHSettingService
	Categories       *GatewayCategoryService
	AppTypes         *GatewayAppTypeService
	Configurations   *GatewayConfigurationService
	Lists            *GatewayListService
	Locations        *GatewayLocationService
	Loggings         *GatewayLoggingService
	ProxyEndpoints   *GatewayProxyEndpointService
	Rules            *GatewayRuleService
}

// NewGatewayService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGatewayService(opts ...option.RequestOption) (r *GatewayService) {
	r = &GatewayService{}
	r.Options = opts
	r.AuditSSHSettings = NewGatewayAuditSSHSettingService(opts...)
	r.Categories = NewGatewayCategoryService(opts...)
	r.AppTypes = NewGatewayAppTypeService(opts...)
	r.Configurations = NewGatewayConfigurationService(opts...)
	r.Lists = NewGatewayListService(opts...)
	r.Locations = NewGatewayLocationService(opts...)
	r.Loggings = NewGatewayLoggingService(opts...)
	r.ProxyEndpoints = NewGatewayProxyEndpointService(opts...)
	r.Rules = NewGatewayRuleService(opts...)
	return
}

// Creates a Zero Trust account with an existing Cloudflare account.
func (r *GatewayService) New(ctx context.Context, body GatewayNewParams, opts ...option.RequestOption) (res *GatewayNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets information about the current Zero Trust account.
func (r *GatewayService) List(ctx context.Context, query GatewayListParams, opts ...option.RequestOption) (res *GatewayListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayNewResponse struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                 `json:"provider_name"`
	JSON         gatewayNewResponseJSON `json:"-"`
}

// gatewayNewResponseJSON contains the JSON metadata for the struct
// [GatewayNewResponse]
type gatewayNewResponseJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListResponse struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                  `json:"provider_name"`
	JSON         gatewayListResponseJSON `json:"-"`
}

// gatewayListResponseJSON contains the JSON metadata for the struct
// [GatewayListResponse]
type gatewayListResponseJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayNewResponseEnvelope struct {
	Errors   []GatewayNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayNewResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayNewResponseEnvelope]
type gatewayNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    gatewayNewResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [GatewayNewResponseEnvelopeErrors]
type gatewayNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    gatewayNewResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [GatewayNewResponseEnvelopeMessages]
type gatewayNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayNewResponseEnvelopeSuccess bool

const (
	GatewayNewResponseEnvelopeSuccessTrue GatewayNewResponseEnvelopeSuccess = true
)

type GatewayListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayListResponseEnvelope struct {
	Errors   []GatewayListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayListResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayListResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayListResponseEnvelopeJSON    `json:"-"`
}

// gatewayListResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListResponseEnvelope]
type gatewayListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    gatewayListResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [GatewayListResponseEnvelopeErrors]
type gatewayListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    gatewayListResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayListResponseEnvelopeMessages]
type gatewayListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayListResponseEnvelopeSuccess bool

const (
	GatewayListResponseEnvelopeSuccessTrue GatewayListResponseEnvelopeSuccess = true
)
