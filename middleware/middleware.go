package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/runnerktb/libs/helper"
	"github.com/runnerktb/libs/token"
	"github.com/runnerktb/libs/universe"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"net/http"
)

func Guard(module, serviceID string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get id from dictionary
		acc := universe.ParseModule(module)
		if acc.ID == "" {
			serr := fmt.Sprintf("module %s not found", module)
			Response(http.StatusForbidden, serr, "", serviceID, module, acc.Action, "", ctx)
			ctx.Abort()
			return
		}
		tok, err := token.ClaimToken(ctx.Request.Header["Authorization"])
		if err != nil {
			serr := serror.NewFromError(err)
			Response(http.StatusForbidden, serr.Error(), "", serviceID, module, acc.Action, "", ctx)
			ctx.Abort()
			return
		}

		ctx.Set("x-token", tok)
		ctx.Set("x-module", module)
		ctx.Set("x-app-id", tok.App)
		ctx.Set("x-action", acc.Action)

		// handle tti super
		if tok.IsOrgAdmin != nil && *tok.IsOrgAdmin == 0 {
			ctx.Next()
			return
		}
		if !helper.CheckArrayString(acc.ID, tok.UserAccess) {
			serr := serror.New("Token module denied")
			Response(http.StatusForbidden, serr.Error(), "", serviceID, module, acc.Action, "", ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
