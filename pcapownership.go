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

// PCAPOwnershipService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPCAPOwnershipService] method
// instead.
type PCAPOwnershipService struct {
	Options []option.RequestOption
}

// NewPCAPOwnershipService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPCAPOwnershipService(opts ...option.RequestOption) (r *PCAPOwnershipService) {
	r = &PCAPOwnershipService{}
	r.Options = opts
	return
}

// Adds an AWS or GCP bucket to use with full packet captures.
func (r *PCAPOwnershipService) New(ctx context.Context, accountIdentifier string, body PCAPOwnershipNewParams, opts ...option.RequestOption) (res *PCAPOwnershipNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPOwnershipNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes buckets added to the packet captures API.
func (r *PCAPOwnershipService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// List all buckets configured for use with PCAPs API.
func (r *PCAPOwnershipService) Get(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]PCAPOwnershipGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPOwnershipGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates buckets added to the packet captures API.
func (r *PCAPOwnershipService) Validate(ctx context.Context, accountIdentifier string, body PCAPOwnershipValidateParams, opts ...option.RequestOption) (res *PCAPOwnershipValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPOwnershipValidateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/validate", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PCAPOwnershipNewResponse struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PCAPOwnershipNewResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                       `json:"validated"`
	JSON      pcapOwnershipNewResponseJSON `json:"-"`
}

// pcapOwnershipNewResponseJSON contains the JSON metadata for the struct
// [PCAPOwnershipNewResponse]
type pcapOwnershipNewResponseJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PCAPOwnershipNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PCAPOwnershipNewResponseStatus string

const (
	PCAPOwnershipNewResponseStatusPending PCAPOwnershipNewResponseStatus = "pending"
	PCAPOwnershipNewResponseStatusSuccess PCAPOwnershipNewResponseStatus = "success"
	PCAPOwnershipNewResponseStatusFailed  PCAPOwnershipNewResponseStatus = "failed"
)

type PCAPOwnershipGetResponse struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PCAPOwnershipGetResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                       `json:"validated"`
	JSON      pcapOwnershipGetResponseJSON `json:"-"`
}

// pcapOwnershipGetResponseJSON contains the JSON metadata for the struct
// [PCAPOwnershipGetResponse]
type pcapOwnershipGetResponseJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PCAPOwnershipGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PCAPOwnershipGetResponseStatus string

const (
	PCAPOwnershipGetResponseStatusPending PCAPOwnershipGetResponseStatus = "pending"
	PCAPOwnershipGetResponseStatusSuccess PCAPOwnershipGetResponseStatus = "success"
	PCAPOwnershipGetResponseStatusFailed  PCAPOwnershipGetResponseStatus = "failed"
)

type PCAPOwnershipValidateResponse struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PCAPOwnershipValidateResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                            `json:"validated"`
	JSON      pcapOwnershipValidateResponseJSON `json:"-"`
}

// pcapOwnershipValidateResponseJSON contains the JSON metadata for the struct
// [PCAPOwnershipValidateResponse]
type pcapOwnershipValidateResponseJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PCAPOwnershipValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PCAPOwnershipValidateResponseStatus string

const (
	PCAPOwnershipValidateResponseStatusPending PCAPOwnershipValidateResponseStatus = "pending"
	PCAPOwnershipValidateResponseStatusSuccess PCAPOwnershipValidateResponseStatus = "success"
	PCAPOwnershipValidateResponseStatusFailed  PCAPOwnershipValidateResponseStatus = "failed"
)

type PCAPOwnershipNewParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
}

func (r PCAPOwnershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PCAPOwnershipNewResponseEnvelope struct {
	Errors   []PCAPOwnershipNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PCAPOwnershipNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PCAPOwnershipNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PCAPOwnershipNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapOwnershipNewResponseEnvelopeJSON    `json:"-"`
}

// pcapOwnershipNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PCAPOwnershipNewResponseEnvelope]
type pcapOwnershipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPOwnershipNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    pcapOwnershipNewResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapOwnershipNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PCAPOwnershipNewResponseEnvelopeErrors]
type pcapOwnershipNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPOwnershipNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    pcapOwnershipNewResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapOwnershipNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PCAPOwnershipNewResponseEnvelopeMessages]
type pcapOwnershipNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PCAPOwnershipNewResponseEnvelopeSuccess bool

const (
	PCAPOwnershipNewResponseEnvelopeSuccessTrue PCAPOwnershipNewResponseEnvelopeSuccess = true
)

type PCAPOwnershipGetResponseEnvelope struct {
	Errors   []PCAPOwnershipGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PCAPOwnershipGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []PCAPOwnershipGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PCAPOwnershipGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PCAPOwnershipGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pcapOwnershipGetResponseEnvelopeJSON       `json:"-"`
}

// pcapOwnershipGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PCAPOwnershipGetResponseEnvelope]
type pcapOwnershipGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPOwnershipGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    pcapOwnershipGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapOwnershipGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PCAPOwnershipGetResponseEnvelopeErrors]
type pcapOwnershipGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPOwnershipGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    pcapOwnershipGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapOwnershipGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PCAPOwnershipGetResponseEnvelopeMessages]
type pcapOwnershipGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PCAPOwnershipGetResponseEnvelopeSuccess bool

const (
	PCAPOwnershipGetResponseEnvelopeSuccessTrue PCAPOwnershipGetResponseEnvelopeSuccess = true
)

type PCAPOwnershipGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       pcapOwnershipGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// pcapOwnershipGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [PCAPOwnershipGetResponseEnvelopeResultInfo]
type pcapOwnershipGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPOwnershipValidateParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r PCAPOwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PCAPOwnershipValidateResponseEnvelope struct {
	Errors   []PCAPOwnershipValidateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PCAPOwnershipValidateResponseEnvelopeMessages `json:"messages,required"`
	Result   PCAPOwnershipValidateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PCAPOwnershipValidateResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapOwnershipValidateResponseEnvelopeJSON    `json:"-"`
}

// pcapOwnershipValidateResponseEnvelopeJSON contains the JSON metadata for the
// struct [PCAPOwnershipValidateResponseEnvelope]
type pcapOwnershipValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPOwnershipValidateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    pcapOwnershipValidateResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapOwnershipValidateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PCAPOwnershipValidateResponseEnvelopeErrors]
type pcapOwnershipValidateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipValidateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPOwnershipValidateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    pcapOwnershipValidateResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapOwnershipValidateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PCAPOwnershipValidateResponseEnvelopeMessages]
type pcapOwnershipValidateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipValidateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PCAPOwnershipValidateResponseEnvelopeSuccess bool

const (
	PCAPOwnershipValidateResponseEnvelopeSuccessTrue PCAPOwnershipValidateResponseEnvelopeSuccess = true
)
