// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dnssec

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// DNSSECService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDNSSECService] method instead.
type DNSSECService struct {
	Options []option.RequestOption
}

// NewDNSSECService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSSECService(opts ...option.RequestOption) (r *DNSSECService) {
	r = &DNSSECService{}
	r.Options = opts
	return
}

// Delete DNSSEC.
func (r *DNSSECService) Delete(ctx context.Context, body DNSSECDeleteParams, opts ...option.RequestOption) (res *DNSSECDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSSECDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/dnssec", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable or disable DNSSEC.
func (r *DNSSECService) Edit(ctx context.Context, params DNSSECEditParams, opts ...option.RequestOption) (res *DNSSEC, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSSECEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/dnssec", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Details about DNSSEC status and configuration.
func (r *DNSSECService) Get(ctx context.Context, query DNSSECGetParams, opts ...option.RequestOption) (res *DNSSEC, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSSECGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dnssec", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSSEC struct {
	// Algorithm key code.
	Algorithm string `json:"algorithm,nullable"`
	// Digest hash.
	Digest string `json:"digest,nullable"`
	// Type of digest algorithm.
	DigestAlgorithm string `json:"digest_algorithm,nullable"`
	// Coded type for digest algorithm.
	DigestType string `json:"digest_type,nullable"`
	// If true, multi-signer DNSSEC is enabled on the zone, allowing multiple providers
	// to serve a DNSSEC-signed zone at the same time. This is required for DNSKEY
	// records (except those automatically generated by Cloudflare) to be added to the
	// zone.
	//
	// See
	// [Multi-signer DNSSEC](https://developers.cloudflare.com/dns/dnssec/multi-signer-dnssec/)
	// for details.
	DNSSECMultiSigner bool `json:"dnssec_multi_signer"`
	// If true, allows Cloudflare to transfer in a DNSSEC-signed zone including
	// signatures from an external provider, without requiring Cloudflare to sign any
	// records on the fly.
	//
	// Note that this feature has some limitations. See
	// [Cloudflare as Secondary](https://developers.cloudflare.com/dns/zone-setups/zone-transfers/cloudflare-as-secondary/setup/#dnssec)
	// for details.
	DNSSECPresigned bool `json:"dnssec_presigned"`
	// Full DS record.
	DS string `json:"ds,nullable"`
	// Flag for DNSSEC record.
	Flags float64 `json:"flags,nullable"`
	// Code for key tag.
	KeyTag float64 `json:"key_tag,nullable"`
	// Algorithm key type.
	KeyType string `json:"key_type,nullable"`
	// When DNSSEC was last modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Public key for DS record.
	PublicKey string `json:"public_key,nullable"`
	// Status of DNSSEC, based on user-desired state and presence of necessary records.
	Status DNSSECStatus `json:"status"`
	JSON   dnssecJSON   `json:"-"`
}

// dnssecJSON contains the JSON metadata for the struct [DNSSEC]
type dnssecJSON struct {
	Algorithm         apijson.Field
	Digest            apijson.Field
	DigestAlgorithm   apijson.Field
	DigestType        apijson.Field
	DNSSECMultiSigner apijson.Field
	DNSSECPresigned   apijson.Field
	DS                apijson.Field
	Flags             apijson.Field
	KeyTag            apijson.Field
	KeyType           apijson.Field
	ModifiedOn        apijson.Field
	PublicKey         apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSSEC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecJSON) RawJSON() string {
	return r.raw
}

// Status of DNSSEC, based on user-desired state and presence of necessary records.
type DNSSECStatus string

const (
	DNSSECStatusActive          DNSSECStatus = "active"
	DNSSECStatusPending         DNSSECStatus = "pending"
	DNSSECStatusDisabled        DNSSECStatus = "disabled"
	DNSSECStatusPendingDisabled DNSSECStatus = "pending-disabled"
	DNSSECStatusError           DNSSECStatus = "error"
)

func (r DNSSECStatus) IsKnown() bool {
	switch r {
	case DNSSECStatusActive, DNSSECStatusPending, DNSSECStatusDisabled, DNSSECStatusPendingDisabled, DNSSECStatusError:
		return true
	}
	return false
}

// Union satisfied by [dnssec.DNSSECDeleteResponseUnknown] or [shared.UnionString].
type DNSSECDeleteResponseUnion interface {
	ImplementsDNSSECDNSSECDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSSECDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DNSSECDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSSECDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSSECDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSSECDeleteResponseUnion           `json:"result"`
	JSON    dnssecDeleteResponseEnvelopeJSON    `json:"-"`
}

// dnssecDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSECDeleteResponseEnvelope]
type dnssecDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSECDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSSECDeleteResponseEnvelopeSuccess bool

const (
	DNSSECDeleteResponseEnvelopeSuccessTrue DNSSECDeleteResponseEnvelopeSuccess = true
)

func (r DNSSECDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSSECDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSSECEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If true, multi-signer DNSSEC is enabled on the zone, allowing multiple providers
	// to serve a DNSSEC-signed zone at the same time. This is required for DNSKEY
	// records (except those automatically generated by Cloudflare) to be added to the
	// zone.
	//
	// See
	// [Multi-signer DNSSEC](https://developers.cloudflare.com/dns/dnssec/multi-signer-dnssec/)
	// for details.
	DNSSECMultiSigner param.Field[bool] `json:"dnssec_multi_signer"`
	// If true, allows Cloudflare to transfer in a DNSSEC-signed zone including
	// signatures from an external provider, without requiring Cloudflare to sign any
	// records on the fly.
	//
	// Note that this feature has some limitations. See
	// [Cloudflare as Secondary](https://developers.cloudflare.com/dns/zone-setups/zone-transfers/cloudflare-as-secondary/setup/#dnssec)
	// for details.
	DNSSECPresigned param.Field[bool] `json:"dnssec_presigned"`
	// Status of DNSSEC, based on user-desired state and presence of necessary records.
	Status param.Field[DNSSECEditParamsStatus] `json:"status"`
}

func (r DNSSECEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of DNSSEC, based on user-desired state and presence of necessary records.
type DNSSECEditParamsStatus string

const (
	DNSSECEditParamsStatusActive   DNSSECEditParamsStatus = "active"
	DNSSECEditParamsStatusDisabled DNSSECEditParamsStatus = "disabled"
)

func (r DNSSECEditParamsStatus) IsKnown() bool {
	switch r {
	case DNSSECEditParamsStatusActive, DNSSECEditParamsStatusDisabled:
		return true
	}
	return false
}

type DNSSECEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSSECEditResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSSEC                            `json:"result"`
	JSON    dnssecEditResponseEnvelopeJSON    `json:"-"`
}

// dnssecEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSECEditResponseEnvelope]
type dnssecEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSECEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSSECEditResponseEnvelopeSuccess bool

const (
	DNSSECEditResponseEnvelopeSuccessTrue DNSSECEditResponseEnvelopeSuccess = true
)

func (r DNSSECEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSSECEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSSECGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSSECGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSSECGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSSEC                           `json:"result"`
	JSON    dnssecGetResponseEnvelopeJSON    `json:"-"`
}

// dnssecGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSECGetResponseEnvelope]
type dnssecGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSECGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSSECGetResponseEnvelopeSuccess bool

const (
	DNSSECGetResponseEnvelopeSuccessTrue DNSSECGetResponseEnvelopeSuccess = true
)

func (r DNSSECGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSSECGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
