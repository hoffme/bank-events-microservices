// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// TransactionsGet invokes GET /transactions operation.
	//
	// Search transactions.
	//
	// GET /transactions
	TransactionsGet(ctx context.Context, params TransactionsGetParams) (*TransactionsGetOK, error)
	// TransactionsTransactionIDGet invokes GET /transactions/{transaction_id} operation.
	//
	// Get transaction with id.
	//
	// GET /transactions/{transaction_id}
	TransactionsTransactionIDGet(ctx context.Context, params TransactionsTransactionIDGetParams) (*Transaction, error)
	// TransactionsTransactionIDPut invokes PUT /transactions/{transaction_id} operation.
	//
	// Create new Transaction.
	//
	// PUT /transactions/{transaction_id}
	TransactionsTransactionIDPut(ctx context.Context, request OptTransactionsTransactionIDPutReq, params TransactionsTransactionIDPutParams) (*TransactionsTransactionIDPutCreated, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}
type errorHandler interface {
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

var _ Handler = struct {
	errorHandler
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// TransactionsGet invokes GET /transactions operation.
//
// Search transactions.
//
// GET /transactions
func (c *Client) TransactionsGet(ctx context.Context, params TransactionsGetParams) (*TransactionsGetOK, error) {
	res, err := c.sendTransactionsGet(ctx, params)
	return res, err
}

func (c *Client) sendTransactionsGet(ctx context.Context, params TransactionsGetParams) (res *TransactionsGetOK, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/transactions"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TransactionsGet",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/transactions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "type" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Type.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "account_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "account_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.AccountID.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "state" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "state",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range params.State {
					if err := func() error {
						return e.EncodeValue(conv.StringToString(string(item)))
					}(); err != nil {
						return errors.Wrapf(err, "[%d]", i)
					}
				}
				return nil
			})
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "date_from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "date_from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.DateFrom.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "date_to" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "date_to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.DateTo.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "currency" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "currency",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range params.Currency {
					if err := func() error {
						return e.EncodeValue(conv.StringToString(string(item)))
					}(); err != nil {
						return errors.Wrapf(err, "[%d]", i)
					}
				}
				return nil
			})
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "limit" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Limit.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "skip" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "skip",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Skip.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "order_by" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "order_by",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OrderBy.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "order_dir" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "order_dir",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OrderDir.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTransactionsGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TransactionsTransactionIDGet invokes GET /transactions/{transaction_id} operation.
//
// Get transaction with id.
//
// GET /transactions/{transaction_id}
func (c *Client) TransactionsTransactionIDGet(ctx context.Context, params TransactionsTransactionIDGetParams) (*Transaction, error) {
	res, err := c.sendTransactionsTransactionIDGet(ctx, params)
	return res, err
}

func (c *Client) sendTransactionsTransactionIDGet(ctx context.Context, params TransactionsTransactionIDGetParams) (res *Transaction, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/transactions/{transaction_id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TransactionsTransactionIDGet",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/transactions/"
	{
		// Encode "transaction_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "transaction_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.TransactionID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTransactionsTransactionIDGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TransactionsTransactionIDPut invokes PUT /transactions/{transaction_id} operation.
//
// Create new Transaction.
//
// PUT /transactions/{transaction_id}
func (c *Client) TransactionsTransactionIDPut(ctx context.Context, request OptTransactionsTransactionIDPutReq, params TransactionsTransactionIDPutParams) (*TransactionsTransactionIDPutCreated, error) {
	res, err := c.sendTransactionsTransactionIDPut(ctx, request, params)
	return res, err
}

func (c *Client) sendTransactionsTransactionIDPut(ctx context.Context, request OptTransactionsTransactionIDPutReq, params TransactionsTransactionIDPutParams) (res *TransactionsTransactionIDPutCreated, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/transactions/{transaction_id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TransactionsTransactionIDPut",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/transactions/"
	{
		// Encode "transaction_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "transaction_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.TransactionID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeTransactionsTransactionIDPutRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTransactionsTransactionIDPutResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
