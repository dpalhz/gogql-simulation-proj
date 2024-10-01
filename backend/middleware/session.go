package utils

import (
	"context"
	"net/http"
	"notes/backend/internal/logger"
	"notes/backend/utils"
)

const (
	authKey				= "auth"
	sessionKey			= "session"
	responseWriterKey	= "responseWriter"
)


// SessionMiddleware handles authentication and adds session information to context
func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.LogInfo("Entering SessionMiddleware")

		ctx := r.Context()
		ctx = context.WithValue(ctx, responseWriterKey, w)
		ctx = context.WithValue(ctx, authKey, true)

		session, err := utils.GetSession(r)		

		if err != nil {
			ctx = context.WithValue(ctx, authKey, false)
			logger.LogInfo("Session is empty or invalid, setting auth to false")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		ctx = context.WithValue(ctx, sessionKey, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
