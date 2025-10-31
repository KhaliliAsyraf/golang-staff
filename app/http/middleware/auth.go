package middleware

import (
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Auth() http.Middleware {
	return func(ctx http.Context) {
		// 1. Get token from "Authorization" header
		authHeader := ctx.Request().Header("Authorization", "")
		if authHeader == "" {
			ctx.Response().String(http.StatusUnauthorized, "Unauthorized").Abort()
			return
		}

		// 2. Extract the token from "Bearer <token>"
		token := strings.TrimPrefix(authHeader, "Bearer ")
		token = strings.TrimSpace(token)

		// 3. Parse token
		_, err := facades.Auth(ctx).Guard("employee").Parse(token)
		if err != nil {
			ctx.Response().String(http.StatusUnauthorized, "Unauthorized").Abort()
			return
		}

		// 4. Continue to the next handler
		ctx.Request().Next()
	}
}
