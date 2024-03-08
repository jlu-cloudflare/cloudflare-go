// File generated from our OpenAPI spec by Stainless.

package kv

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// NamespaceBulkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNamespaceBulkService] method
// instead.
type NamespaceBulkService struct {
	Options []option.RequestOption
}

// NewNamespaceBulkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceBulkService(opts ...option.RequestOption) (r *NamespaceBulkService) {
	r = &NamespaceBulkService{}
	r.Options = opts
	return
}

// Write multiple keys and values at once. Body should be an array of up to 10,000
// key-value pairs to be stored, along with optional expiration information.
// Existing values and expirations will be overwritten. If neither `expiration` nor
// `expiration_ttl` is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored. The entire request
// size must be 100 megabytes or less.
func (r *NamespaceBulkService) Update(ctx context.Context, namespaceID string, params NamespaceBulkUpdateParams, opts ...option.RequestOption) (res *NamespaceBulkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceBulkUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove multiple KV pairs from the namespace. Body should be an array of up to
// 10,000 keys to be removed.
func (r *NamespaceBulkService) Delete(ctx context.Context, namespaceID string, params NamespaceBulkDeleteParams, opts ...option.RequestOption) (res *NamespaceBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceBulkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [kv.NamespaceBulkUpdateResponseUnknown] or
// [shared.UnionString].
type NamespaceBulkUpdateResponse interface {
	ImplementsKVNamespaceBulkUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceBulkUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [kv.NamespaceBulkDeleteResponseUnknown] or
// [shared.UnionString].
type NamespaceBulkDeleteResponse interface {
	ImplementsKVNamespaceBulkDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceBulkDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NamespaceBulkUpdateParams struct {
	// Identifier
	AccountID param.Field[string]                          `path:"account_id,required"`
	Body      param.Field[[]NamespaceBulkUpdateParamsBody] `json:"body,required"`
}

func (r NamespaceBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type NamespaceBulkUpdateParamsBody struct {
	// Whether or not the server should base64 decode the value before storing it.
	// Useful for writing values that wouldn't otherwise be valid JSON strings, such as
	// images.
	Base64 param.Field[bool] `json:"base64"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// should expire.
	Expiration param.Field[float64] `json:"expiration"`
	// The number of seconds for which the key should be visible before it expires. At
	// least 60.
	ExpirationTTL param.Field[float64] `json:"expiration_ttl"`
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid.
	Key param.Field[string] `json:"key"`
	// Arbitrary JSON that is associated with a key.
	Metadata param.Field[interface{}] `json:"metadata"`
	// A UTF-8 encoded string to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value"`
}

func (r NamespaceBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceBulkUpdateResponseEnvelope struct {
	Errors   []NamespaceBulkUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceBulkUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   NamespaceBulkUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NamespaceBulkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    namespaceBulkUpdateResponseEnvelopeJSON    `json:"-"`
}

// namespaceBulkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceBulkUpdateResponseEnvelope]
type namespaceBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceBulkUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    namespaceBulkUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceBulkUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [NamespaceBulkUpdateResponseEnvelopeErrors]
type namespaceBulkUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceBulkUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    namespaceBulkUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceBulkUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [NamespaceBulkUpdateResponseEnvelopeMessages]
type namespaceBulkUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceBulkUpdateResponseEnvelopeSuccess bool

const (
	NamespaceBulkUpdateResponseEnvelopeSuccessTrue NamespaceBulkUpdateResponseEnvelopeSuccess = true
)

type NamespaceBulkDeleteParams struct {
	// Identifier
	AccountID param.Field[string]   `path:"account_id,required"`
	Body      param.Field[[]string] `json:"body,required"`
}

func (r NamespaceBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type NamespaceBulkDeleteResponseEnvelope struct {
	Errors   []NamespaceBulkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceBulkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   NamespaceBulkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NamespaceBulkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    namespaceBulkDeleteResponseEnvelopeJSON    `json:"-"`
}

// namespaceBulkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceBulkDeleteResponseEnvelope]
type namespaceBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceBulkDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    namespaceBulkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceBulkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [NamespaceBulkDeleteResponseEnvelopeErrors]
type namespaceBulkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceBulkDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    namespaceBulkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceBulkDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [NamespaceBulkDeleteResponseEnvelopeMessages]
type namespaceBulkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceBulkDeleteResponseEnvelopeSuccess bool

const (
	NamespaceBulkDeleteResponseEnvelopeSuccessTrue NamespaceBulkDeleteResponseEnvelopeSuccess = true
)
