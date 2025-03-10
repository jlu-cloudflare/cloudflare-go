// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// RequestAssetService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestAssetService] method instead.
type RequestAssetService struct {
	Options []option.RequestOption
}

// NewRequestAssetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestAssetService(opts ...option.RequestOption) (r *RequestAssetService) {
	r = &RequestAssetService{}
	r.Options = opts
	return
}

// List Request Assets
func (r *RequestAssetService) New(ctx context.Context, accountIdentifier string, requestIdentifier string, body RequestAssetNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[RequestAssetNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset", accountIdentifier, requestIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Request Assets
func (r *RequestAssetService) NewAutoPaging(ctx context.Context, accountIdentifier string, requestIdentifier string, body RequestAssetNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RequestAssetNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, accountIdentifier, requestIdentifier, body, opts...))
}

// Update a Request Asset
func (r *RequestAssetService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, assetIdentifer string, body RequestAssetUpdateParams, opts ...option.RequestOption) (res *RequestAssetUpdateResponse, err error) {
	var env RequestAssetUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	if assetIdentifer == "" {
		err = errors.New("missing required asset_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset/%s", accountIdentifier, requestIdentifier, assetIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Request Asset
func (r *RequestAssetService) Delete(ctx context.Context, accountIdentifier string, requestIdentifier string, assetIdentifer string, opts ...option.RequestOption) (res *RequestAssetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	if assetIdentifer == "" {
		err = errors.New("missing required asset_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset/%s", accountIdentifier, requestIdentifier, assetIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a Request Asset
func (r *RequestAssetService) Get(ctx context.Context, accountIdentifier string, requestIdentifier string, assetIdentifer string, opts ...option.RequestOption) (res *pagination.SinglePage[RequestAssetGetResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	if assetIdentifer == "" {
		err = errors.New("missing required asset_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset/%s", accountIdentifier, requestIdentifier, assetIdentifer)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a Request Asset
func (r *RequestAssetService) GetAutoPaging(ctx context.Context, accountIdentifier string, requestIdentifier string, assetIdentifer string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RequestAssetGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, accountIdentifier, requestIdentifier, assetIdentifer, opts...))
}

type RequestAssetNewResponse struct {
	// Asset ID
	ID int64 `json:"id,required"`
	// Asset name
	Name string `json:"name,required"`
	// Asset creation time
	Created time.Time `json:"created" format:"date-time"`
	// Asset description
	Description string `json:"description"`
	// Asset file type
	FileType string                      `json:"file_type"`
	JSON     requestAssetNewResponseJSON `json:"-"`
}

// requestAssetNewResponseJSON contains the JSON metadata for the struct
// [RequestAssetNewResponse]
type requestAssetNewResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Created     apijson.Field
	Description apijson.Field
	FileType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetNewResponseJSON) RawJSON() string {
	return r.raw
}

type RequestAssetUpdateResponse struct {
	// Asset ID
	ID int64 `json:"id,required"`
	// Asset name
	Name string `json:"name,required"`
	// Asset creation time
	Created time.Time `json:"created" format:"date-time"`
	// Asset description
	Description string `json:"description"`
	// Asset file type
	FileType string                         `json:"file_type"`
	JSON     requestAssetUpdateResponseJSON `json:"-"`
}

// requestAssetUpdateResponseJSON contains the JSON metadata for the struct
// [RequestAssetUpdateResponse]
type requestAssetUpdateResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Created     apijson.Field
	Description apijson.Field
	FileType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RequestAssetDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestAssetDeleteResponseSuccess `json:"success,required"`
	JSON    requestAssetDeleteResponseJSON    `json:"-"`
}

// requestAssetDeleteResponseJSON contains the JSON metadata for the struct
// [RequestAssetDeleteResponse]
type requestAssetDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestAssetDeleteResponseSuccess bool

const (
	RequestAssetDeleteResponseSuccessTrue RequestAssetDeleteResponseSuccess = true
)

func (r RequestAssetDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case RequestAssetDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type RequestAssetGetResponse struct {
	// Asset ID
	ID int64 `json:"id,required"`
	// Asset name
	Name string `json:"name,required"`
	// Asset creation time
	Created time.Time `json:"created" format:"date-time"`
	// Asset description
	Description string `json:"description"`
	// Asset file type
	FileType string                      `json:"file_type"`
	JSON     requestAssetGetResponseJSON `json:"-"`
}

// requestAssetGetResponseJSON contains the JSON metadata for the struct
// [RequestAssetGetResponse]
type requestAssetGetResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Created     apijson.Field
	Description apijson.Field
	FileType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetGetResponseJSON) RawJSON() string {
	return r.raw
}

type RequestAssetNewParams struct {
	// Page number of results
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page
	PerPage param.Field[int64] `json:"per_page,required"`
}

func (r RequestAssetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestAssetUpdateParams struct {
	// Asset file to upload
	Source param.Field[string] `json:"source"`
}

func (r RequestAssetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestAssetUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RequestAssetUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestAssetUpdateResponse                `json:"result"`
	JSON    requestAssetUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestAssetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestAssetUpdateResponseEnvelope]
type requestAssetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestAssetUpdateResponseEnvelopeSuccess bool

const (
	RequestAssetUpdateResponseEnvelopeSuccessTrue RequestAssetUpdateResponseEnvelopeSuccess = true
)

func (r RequestAssetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestAssetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
