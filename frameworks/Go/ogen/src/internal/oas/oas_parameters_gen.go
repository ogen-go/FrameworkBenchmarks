// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"net/http"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// CachingParams is parameters of Caching operation.
type CachingParams struct {
	Count int64
}

func unpackCachingParams(packed middleware.Parameters) (params CachingParams) {
	{
		key := middleware.ParameterKey{
			Name: "count",
			In:   "query",
		}
		params.Count = packed[key].(int64)
	}
	return params
}

func decodeCachingParams(args [0]string, argsEscaped bool, r *http.Request) (params CachingParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: count.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "count",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.Count = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "count",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// QueriesParams is parameters of Queries operation.
type QueriesParams struct {
	Queries int64
}

func unpackQueriesParams(packed middleware.Parameters) (params QueriesParams) {
	{
		key := middleware.ParameterKey{
			Name: "queries",
			In:   "query",
		}
		params.Queries = packed[key].(int64)
	}
	return params
}

func decodeQueriesParams(args [0]string, argsEscaped bool, r *http.Request) (params QueriesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: queries.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "queries",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.Queries = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "queries",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// UpdatesParams is parameters of Updates operation.
type UpdatesParams struct {
	Queries int64
}

func unpackUpdatesParams(packed middleware.Parameters) (params UpdatesParams) {
	{
		key := middleware.ParameterKey{
			Name: "queries",
			In:   "query",
		}
		params.Queries = packed[key].(int64)
	}
	return params
}

func decodeUpdatesParams(args [0]string, argsEscaped bool, r *http.Request) (params UpdatesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: queries.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "queries",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.Queries = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "queries",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
