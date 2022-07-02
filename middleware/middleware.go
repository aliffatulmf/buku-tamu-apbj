package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var SECRET_KEY = "mwQ5VGVVT7MdtbbqecnDPmbeDnTtugjNgy2PzRkd"

func WebviewMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		webview := ctx.Request.Header["Masterkey"]

		if len(webview) == 0 {
			ctx.HTML(http.StatusForbidden, "launcher.html", gin.H{})
			ctx.Abort()
		} else {
			if webview[0] != SECRET_KEY {
				ctx.HTML(http.StatusForbidden, "launcher.html", gin.H{})
				ctx.Abort()
			}
		}

		ctx.Next()
	}
}
