package auth

import (
	"net/http"
)

type options struct {
	disabled   bool
	skipFilter func(r *http.Request) bool
	check      func(http.ResponseWriter, *http.Request) bool
}

type Option func(opts *options)

func WithDisabled(disable bool) Option {
	return func(opts *options) {
		opts.disabled = disable
	}
}

func WithSkipFilter(filter func(r *http.Request) bool) Option {
	return func(opts *options) {
		opts.skipFilter = filter
	}
}

func WithCheck(check func(http.ResponseWriter, *http.Request) bool) Option {
	return func(opts *options) {
		opts.check = check
	}
}

func AuthMiddleware(opts ...Option) func(next http.Handler) http.Handler {
	opt := options{
		disabled: false,
		check: func(w http.ResponseWriter, _ *http.Request) bool {
			return true
		},
		skipFilter: func(r *http.Request) bool {
			return false
		},
	}
	for _, o := range opts {
		o(&opt)
	}
	return func(handler http.Handler) http.Handler {
		if opt.disabled {
			return handler
		}
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if opt.skipFilter(request) {
				handler.ServeHTTP(writer, request)
				return
			}
			if checkRes := opt.check(writer, request); !checkRes {
				return
			}
			handler.ServeHTTP(writer, request)
		})
	}

}
