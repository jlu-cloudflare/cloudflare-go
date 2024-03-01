// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZeroTrustGatewayLocationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustGatewayLocationService] method instead.
type ZeroTrustGatewayLocationService struct {
	Options []option.RequestOption
}

// NewZeroTrustGatewayLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayLocationService(opts ...option.RequestOption) (r *ZeroTrustGatewayLocationService) {
	r = &ZeroTrustGatewayLocationService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway location.
func (r *ZeroTrustGatewayLocationService) New(ctx context.Context, params ZeroTrustGatewayLocationNewParams, opts ...option.RequestOption) (res *ZeroTrustGatewayLocationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayLocationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway location.
func (r *ZeroTrustGatewayLocationService) Update(ctx context.Context, locationID interface{}, params ZeroTrustGatewayLocationUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewayLocationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayLocationUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", params.AccountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Zero Trust Gateway locations for an account.
func (r *ZeroTrustGatewayLocationService) List(ctx context.Context, query ZeroTrustGatewayLocationListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayLocationListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a configured Zero Trust Gateway location.
func (r *ZeroTrustGatewayLocationService) Delete(ctx context.Context, locationID interface{}, body ZeroTrustGatewayLocationDeleteParams, opts ...option.RequestOption) (res *ZeroTrustGatewayLocationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayLocationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", body.AccountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway location.
func (r *ZeroTrustGatewayLocationService) Get(ctx context.Context, locationID interface{}, query ZeroTrustGatewayLocationGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayLocationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", query.AccountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayLocationNewResponse struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []ZeroTrustGatewayLocationNewResponseNetwork `json:"networks"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayLocationNewResponseJSON      `json:"-"`
}

// zeroTrustGatewayLocationNewResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayLocationNewResponse]
type zeroTrustGatewayLocationNewResponseJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationNewResponseNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                         `json:"network,required"`
	JSON    zeroTrustGatewayLocationNewResponseNetworkJSON `json:"-"`
}

// zeroTrustGatewayLocationNewResponseNetworkJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayLocationNewResponseNetwork]
type zeroTrustGatewayLocationNewResponseNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationNewResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationUpdateResponse struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []ZeroTrustGatewayLocationUpdateResponseNetwork `json:"networks"`
	UpdatedAt time.Time                                       `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayLocationUpdateResponseJSON      `json:"-"`
}

// zeroTrustGatewayLocationUpdateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayLocationUpdateResponse]
type zeroTrustGatewayLocationUpdateResponseJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationUpdateResponseNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                            `json:"network,required"`
	JSON    zeroTrustGatewayLocationUpdateResponseNetworkJSON `json:"-"`
}

// zeroTrustGatewayLocationUpdateResponseNetworkJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayLocationUpdateResponseNetwork]
type zeroTrustGatewayLocationUpdateResponseNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationUpdateResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationListResponse struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []ZeroTrustGatewayLocationListResponseNetwork `json:"networks"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayLocationListResponseJSON      `json:"-"`
}

// zeroTrustGatewayLocationListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayLocationListResponse]
type zeroTrustGatewayLocationListResponseJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationListResponseNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                          `json:"network,required"`
	JSON    zeroTrustGatewayLocationListResponseNetworkJSON `json:"-"`
}

// zeroTrustGatewayLocationListResponseNetworkJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayLocationListResponseNetwork]
type zeroTrustGatewayLocationListResponseNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationListResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZeroTrustGatewayLocationDeleteResponseUnknown] or
// [shared.UnionString].
type ZeroTrustGatewayLocationDeleteResponse interface {
	ImplementsZeroTrustGatewayLocationDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGatewayLocationDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustGatewayLocationGetResponse struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []ZeroTrustGatewayLocationGetResponseNetwork `json:"networks"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayLocationGetResponseJSON      `json:"-"`
}

// zeroTrustGatewayLocationGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayLocationGetResponse]
type zeroTrustGatewayLocationGetResponseJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationGetResponseNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                         `json:"network,required"`
	JSON    zeroTrustGatewayLocationGetResponseNetworkJSON `json:"-"`
}

// zeroTrustGatewayLocationGetResponseNetworkJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayLocationGetResponseNetwork]
type zeroTrustGatewayLocationGetResponseNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationGetResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]ZeroTrustGatewayLocationNewParamsNetwork] `json:"networks"`
}

func (r ZeroTrustGatewayLocationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayLocationNewParamsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r ZeroTrustGatewayLocationNewParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayLocationNewResponseEnvelope struct {
	Errors   []ZeroTrustGatewayLocationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayLocationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayLocationNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayLocationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayLocationNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayLocationNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayLocationNewResponseEnvelope]
type zeroTrustGatewayLocationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationNewResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustGatewayLocationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayLocationNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayLocationNewResponseEnvelopeErrors]
type zeroTrustGatewayLocationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationNewResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustGatewayLocationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayLocationNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLocationNewResponseEnvelopeMessages]
type zeroTrustGatewayLocationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayLocationNewResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayLocationNewResponseEnvelopeSuccessTrue ZeroTrustGatewayLocationNewResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayLocationUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]ZeroTrustGatewayLocationUpdateParamsNetwork] `json:"networks"`
}

func (r ZeroTrustGatewayLocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayLocationUpdateParamsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r ZeroTrustGatewayLocationUpdateParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayLocationUpdateResponseEnvelope struct {
	Errors   []ZeroTrustGatewayLocationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayLocationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayLocationUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayLocationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayLocationUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayLocationUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayLocationUpdateResponseEnvelope]
type zeroTrustGatewayLocationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustGatewayLocationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayLocationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLocationUpdateResponseEnvelopeErrors]
type zeroTrustGatewayLocationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustGatewayLocationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayLocationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLocationUpdateResponseEnvelopeMessages]
type zeroTrustGatewayLocationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayLocationUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayLocationUpdateResponseEnvelopeSuccessTrue ZeroTrustGatewayLocationUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayLocationListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayLocationListResponseEnvelope struct {
	Errors   []ZeroTrustGatewayLocationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayLocationListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayLocationListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustGatewayLocationListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustGatewayLocationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustGatewayLocationListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustGatewayLocationListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayLocationListResponseEnvelope]
type zeroTrustGatewayLocationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationListResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustGatewayLocationListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayLocationListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLocationListResponseEnvelopeErrors]
type zeroTrustGatewayLocationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationListResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustGatewayLocationListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayLocationListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLocationListResponseEnvelopeMessages]
type zeroTrustGatewayLocationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayLocationListResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayLocationListResponseEnvelopeSuccessTrue ZeroTrustGatewayLocationListResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayLocationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       zeroTrustGatewayLocationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustGatewayLocationListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLocationListResponseEnvelopeResultInfo]
type zeroTrustGatewayLocationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayLocationDeleteResponseEnvelope struct {
	Errors   []ZeroTrustGatewayLocationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayLocationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayLocationDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayLocationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayLocationDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayLocationDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayLocationDeleteResponseEnvelope]
type zeroTrustGatewayLocationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationDeleteResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustGatewayLocationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayLocationDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLocationDeleteResponseEnvelopeErrors]
type zeroTrustGatewayLocationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationDeleteResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustGatewayLocationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayLocationDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLocationDeleteResponseEnvelopeMessages]
type zeroTrustGatewayLocationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayLocationDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayLocationDeleteResponseEnvelopeSuccessTrue ZeroTrustGatewayLocationDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayLocationGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayLocationGetResponseEnvelope struct {
	Errors   []ZeroTrustGatewayLocationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayLocationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayLocationGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayLocationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayLocationGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayLocationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayLocationGetResponseEnvelope]
type zeroTrustGatewayLocationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustGatewayLocationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayLocationGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayLocationGetResponseEnvelopeErrors]
type zeroTrustGatewayLocationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLocationGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustGatewayLocationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayLocationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLocationGetResponseEnvelopeMessages]
type zeroTrustGatewayLocationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayLocationGetResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayLocationGetResponseEnvelopeSuccessTrue ZeroTrustGatewayLocationGetResponseEnvelopeSuccess = true
)
