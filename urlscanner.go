// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// URLScannerService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewURLScannerService] method instead.
type URLScannerService struct {
	Options []option.RequestOption
	Scans   *URLScannerScanService
}

// NewURLScannerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewURLScannerService(opts ...option.RequestOption) (r *URLScannerService) {
	r = &URLScannerService{}
	r.Options = opts
	r.Scans = NewURLScannerScanService(opts...)
	return
}

// Search scans by date and webpages' requests, including full URL (after
// redirects), hostname, and path. <br/> A successful scan will appear in search
// results a few minutes after finishing but may take much longer if the system in
// under load. By default, only successfully completed scans will appear in search
// results, unless searching by `scanId`. Please take into account that older scans
// may be removed from the search index at an unspecified time.
func (r *URLScannerService) Scan(ctx context.Context, accountID string, query URLScannerScanParams, opts ...option.RequestOption) (res *URLScannerScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env URLScannerScanResponseEnvelope
	path := fmt.Sprintf("accounts/%s/urlscanner/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type URLScannerScanResponse struct {
	Tasks []URLScannerScanResponseTask `json:"tasks,required"`
	JSON  urlScannerScanResponseJSON   `json:"-"`
}

// urlScannerScanResponseJSON contains the JSON metadata for the struct
// [URLScannerScanResponse]
type urlScannerScanResponseJSON struct {
	Tasks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanResponseTask struct {
	// Whether scan was successful or not
	Success bool `json:"success,required"`
	// When scan was submitted (UTC)
	Time time.Time `json:"time,required" format:"date-time"`
	// Scan url (after redirects)
	URL string `json:"url,required"`
	// Scan id
	UUID string                         `json:"uuid,required" format:"uuid"`
	JSON urlScannerScanResponseTaskJSON `json:"-"`
}

// urlScannerScanResponseTaskJSON contains the JSON metadata for the struct
// [URLScannerScanResponseTask]
type urlScannerScanResponseTaskJSON struct {
	Success     apijson.Field
	Time        apijson.Field
	URL         apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanParams struct {
	// Return only scans created by account.
	AccountScans param.Field[bool] `query:"account_scans"`
	// Filter scans requested before date (inclusive).
	DateEnd param.Field[time.Time] `query:"date_end" format:"date-time"`
	// Filter scans requested after date (inclusive).
	DateStart param.Field[time.Time] `query:"date_start" format:"date-time"`
	// Filter scans by hostname of _any_ request made by the webpage.
	Hostname param.Field[string] `query:"hostname"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Pagination cursor to get the next set of results.
	NextCursor param.Field[string] `query:"next_cursor"`
	// Filter scans by main page hostname .
	PageHostname param.Field[string] `query:"page_hostname"`
	// Filter scans by exact match URL path (also supports suffix search).
	PagePath param.Field[string] `query:"page_path"`
	// Filter scans by exact match to scanned URL (_after redirects_)
	PageURL param.Field[string] `query:"page_url"`
	// Filter scans by url path of _any_ request made by the webpage.
	Path param.Field[string] `query:"path"`
	// Scan uuid
	ScanID param.Field[string] `query:"scanId" format:"uuid"`
	// Filter scans by exact match URL of _any_ request made by the webpage
	URL param.Field[string] `query:"url"`
}

// URLQuery serializes [URLScannerScanParams]'s query parameters as `url.Values`.
func (r URLScannerScanParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type URLScannerScanResponseEnvelope struct {
	Errors   []URLScannerScanResponseEnvelopeErrors   `json:"errors,required"`
	Messages []URLScannerScanResponseEnvelopeMessages `json:"messages,required"`
	Result   URLScannerScanResponse                   `json:"result,required"`
	// Whether search request was successful or not
	Success bool                               `json:"success,required"`
	JSON    urlScannerScanResponseEnvelopeJSON `json:"-"`
}

// urlScannerScanResponseEnvelopeJSON contains the JSON metadata for the struct
// [URLScannerScanResponseEnvelope]
type urlScannerScanResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanResponseEnvelopeErrors struct {
	Message string                                   `json:"message,required"`
	JSON    urlScannerScanResponseEnvelopeErrorsJSON `json:"-"`
}

// urlScannerScanResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [URLScannerScanResponseEnvelopeErrors]
type urlScannerScanResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanResponseEnvelopeMessages struct {
	Message string                                     `json:"message,required"`
	JSON    urlScannerScanResponseEnvelopeMessagesJSON `json:"-"`
}

// urlScannerScanResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [URLScannerScanResponseEnvelopeMessages]
type urlScannerScanResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
