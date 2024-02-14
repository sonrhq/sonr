package middleware

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
)

// HeaderKey is the key for the htmx request header
type HeaderKey string

// HTMXHeaderKey is the key for the htmx request header
const HTMXHeaderKey HeaderKey = "htmx-request-header"

func HTMX(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		hxh := htmx.HxRequestHeaderFromRequest(c.Request())

		ctx = context.WithValue(ctx, HTMXHeaderKey, hxh)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

// ShowTempl renders a templ.Component
func ShowTempl(cmp templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Render the templ component to a `template.HTML` value.
		html, err := templ.ToGoHTML(context.Background(), cmp)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.HTML(http.StatusOK, string(html))
	}
}

// ShowPage renders a templ.Component
func ShowPage(cmp templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Render the templ component to a `template.HTML` value.
		html, err := templ.ToGoHTML(context.Background(), cmp)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.HTML(http.StatusOK, string(html))
	}
}
