package main

import (
	"time"

	"github.com/go-kit/log"
)

// func logginMiddleware(logger log.Logger) endpoint.Middleware {
// 	return func(next endpoint.Endpoint) endpoint.Endpoint {
// 		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 			logger.Log("mes", "calling endpoint")
// 			defer logger.Log("msg", "calling endpoint")
// 			return next(ctx, request)
// 		}
// 	}
// }

type logginMiddleware struct {
	logger log.Logger
	next   StringService
}

func (mw logginMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "upppercase", "input", s, "output", output, "Err", err, "took", time.Since(begin))
	}(time.Now())
	output, err = mw.next.Uppercase(s)
	return
}

func (mw logginMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())
	n = mw.next.Count(s)
	return
}
