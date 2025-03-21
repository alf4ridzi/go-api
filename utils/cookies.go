package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Cookies struct{}

func (c *Cookies) SetCookie(ctx *gin.Context, name string, value string, expiry time.Duration) {
	ctx.SetCookie(name, value, int(expiry.Seconds()), "/", "", false, true)
}

func (c *Cookies) GetCookie(ctx *gin.Context, name string) (string, error) {
	return ctx.Cookie(name)
}
