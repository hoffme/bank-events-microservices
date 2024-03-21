// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

// handleTransactionsGetRequest handles GET /transactions operation.
//
// Search transactions.
//
// GET /transactions
func (s *Server) handleTransactionsGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/transactions"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TransactionsGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "TransactionsGet",
			ID:   "",
		}
	)
	params, err := decodeTransactionsGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *TransactionsGetOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "TransactionsGet",
			OperationSummary: "",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "type",
					In:   "query",
				}: params.Type,
				{
					Name: "account_id",
					In:   "query",
				}: params.AccountID,
				{
					Name: "state",
					In:   "query",
				}: params.State,
				{
					Name: "date_from",
					In:   "query",
				}: params.DateFrom,
				{
					Name: "date_to",
					In:   "query",
				}: params.DateTo,
				{
					Name: "currency",
					In:   "query",
				}: params.Currency,
				{
					Name: "limit",
					In:   "query",
				}: params.Limit,
				{
					Name: "skip",
					In:   "query",
				}: params.Skip,
				{
					Name: "order_by",
					In:   "query",
				}: params.OrderBy,
				{
					Name: "order_dir",
					In:   "query",
				}: params.OrderDir,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = TransactionsGetParams
			Response = *TransactionsGetOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackTransactionsGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.TransactionsGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.TransactionsGet(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			recordError("Internal", err)
		}
		return
	}

	if err := encodeTransactionsGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTransactionsTransactionIDGetRequest handles GET /transactions/{transaction_id} operation.
//
// Get transaction with id.
//
// GET /transactions/{transaction_id}
func (s *Server) handleTransactionsTransactionIDGetRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/transactions/{transaction_id}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TransactionsTransactionIDGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "TransactionsTransactionIDGet",
			ID:   "",
		}
	)
	params, err := decodeTransactionsTransactionIDGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Transaction
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "TransactionsTransactionIDGet",
			OperationSummary: "",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "transaction_id",
					In:   "path",
				}: params.TransactionID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = TransactionsTransactionIDGetParams
			Response = *Transaction
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackTransactionsTransactionIDGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.TransactionsTransactionIDGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.TransactionsTransactionIDGet(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			recordError("Internal", err)
		}
		return
	}

	if err := encodeTransactionsTransactionIDGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTransactionsTransactionIDPutRequest handles PUT /transactions/{transaction_id} operation.
//
// Create new Transaction.
//
// PUT /transactions/{transaction_id}
func (s *Server) handleTransactionsTransactionIDPutRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/transactions/{transaction_id}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TransactionsTransactionIDPut",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "TransactionsTransactionIDPut",
			ID:   "",
		}
	)
	params, err := decodeTransactionsTransactionIDPutParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeTransactionsTransactionIDPutRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *TransactionsTransactionIDPutCreated
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "TransactionsTransactionIDPut",
			OperationSummary: "",
			OperationID:      "",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "transaction_id",
					In:   "path",
				}: params.TransactionID,
			},
			Raw: r,
		}

		type (
			Request  = OptTransactionsTransactionIDPutReq
			Params   = TransactionsTransactionIDPutParams
			Response = *TransactionsTransactionIDPutCreated
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackTransactionsTransactionIDPutParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.TransactionsTransactionIDPut(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.TransactionsTransactionIDPut(ctx, request, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			recordError("Internal", err)
		}
		return
	}

	if err := encodeTransactionsTransactionIDPutResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
