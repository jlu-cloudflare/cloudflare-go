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

// DevicePolicyExcludeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDevicePolicyExcludeService]
// method instead.
type DevicePolicyExcludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyExcludeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevicePolicyExcludeService(opts ...option.RequestOption) (r *DevicePolicyExcludeService) {
	r = &DevicePolicyExcludeService{}
	r.Options = opts
	return
}

// Sets the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyExcludeService) Update(ctx context.Context, params DevicePolicyExcludeUpdateParams, opts ...option.RequestOption) (res *[]DevicePolicyExcludeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyExcludeService) List(ctx context.Context, query DevicePolicyExcludeListParams, opts ...option.RequestOption) (res *[]DevicePolicyExcludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyExcludeService) Get(ctx context.Context, policyID string, query DevicePolicyExcludeGetParams, opts ...option.RequestOption) (res *[]DevicePolicyExcludeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/exclude", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyExcludeUpdateResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                `json:"host"`
	JSON devicePolicyExcludeUpdateResponseJSON `json:"-"`
}

// devicePolicyExcludeUpdateResponseJSON contains the JSON metadata for the struct
// [DevicePolicyExcludeUpdateResponse]
type devicePolicyExcludeUpdateResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeListResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                              `json:"host"`
	JSON devicePolicyExcludeListResponseJSON `json:"-"`
}

// devicePolicyExcludeListResponseJSON contains the JSON metadata for the struct
// [DevicePolicyExcludeListResponse]
type devicePolicyExcludeListResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeGetResponse struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                             `json:"host"`
	JSON devicePolicyExcludeGetResponseJSON `json:"-"`
}

// devicePolicyExcludeGetResponseJSON contains the JSON metadata for the struct
// [DevicePolicyExcludeGetResponse]
type devicePolicyExcludeGetResponseJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeUpdateParams struct {
	AccountID param.Field[interface{}]                           `path:"account_id,required"`
	Body      param.Field[[]DevicePolicyExcludeUpdateParamsBody] `json:"body,required"`
}

func (r DevicePolicyExcludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyExcludeUpdateParamsBody struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r DevicePolicyExcludeUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyExcludeUpdateResponseEnvelope struct {
	Errors   []DevicePolicyExcludeUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyExcludeUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyExcludeUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyExcludeUpdateResponseEnvelope]
type devicePolicyExcludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    devicePolicyExcludeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyExcludeUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DevicePolicyExcludeUpdateResponseEnvelopeErrors]
type devicePolicyExcludeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    devicePolicyExcludeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyExcludeUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DevicePolicyExcludeUpdateResponseEnvelopeMessages]
type devicePolicyExcludeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyExcludeUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeUpdateResponseEnvelopeSuccessTrue DevicePolicyExcludeUpdateResponseEnvelopeSuccess = true
)

type DevicePolicyExcludeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       devicePolicyExcludeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DevicePolicyExcludeUpdateResponseEnvelopeResultInfo]
type devicePolicyExcludeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DevicePolicyExcludeListResponseEnvelope struct {
	Errors   []DevicePolicyExcludeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyExcludeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyExcludeListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyExcludeListResponseEnvelope]
type devicePolicyExcludeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    devicePolicyExcludeListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyExcludeListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DevicePolicyExcludeListResponseEnvelopeErrors]
type devicePolicyExcludeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    devicePolicyExcludeListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyExcludeListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DevicePolicyExcludeListResponseEnvelopeMessages]
type devicePolicyExcludeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyExcludeListResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeListResponseEnvelopeSuccessTrue DevicePolicyExcludeListResponseEnvelopeSuccess = true
)

type DevicePolicyExcludeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       devicePolicyExcludeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DevicePolicyExcludeListResponseEnvelopeResultInfo]
type devicePolicyExcludeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DevicePolicyExcludeGetResponseEnvelope struct {
	Errors   []DevicePolicyExcludeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyExcludeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyExcludeGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyExcludeGetResponseEnvelope]
type devicePolicyExcludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    devicePolicyExcludeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyExcludeGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DevicePolicyExcludeGetResponseEnvelopeErrors]
type devicePolicyExcludeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyExcludeGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    devicePolicyExcludeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyExcludeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DevicePolicyExcludeGetResponseEnvelopeMessages]
type devicePolicyExcludeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyExcludeGetResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeGetResponseEnvelopeSuccessTrue DevicePolicyExcludeGetResponseEnvelopeSuccess = true
)

type DevicePolicyExcludeGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       devicePolicyExcludeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DevicePolicyExcludeGetResponseEnvelopeResultInfo]
type devicePolicyExcludeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
