package middleware

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"time"
)

func Middleware(log *zap.Logger, handle fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		log.Debug("request",
			zap.String("path", string(ctx.Path())),
			zap.String("headers", string(ctx.Request.Header.Header())),
			zap.String("body", string(ctx.Request.Body())),
		)

		now := time.Now()
		handle(ctx)
		duration := time.Since(now)

		log.Debug("response",
			zap.Int("status", ctx.Response.StatusCode()),
			zap.String("headers", string(ctx.Response.Header.Header())),
			zap.String("body", string(ctx.Response.Body())),
			zap.Float64("duration", duration.Seconds()),
		)
	}
}
