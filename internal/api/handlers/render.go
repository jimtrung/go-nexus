package handlers

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, component templ.Component) error {
    return component.Render(c.Request.Context(), c.Writer)
}
