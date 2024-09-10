package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ForwardRequest(c echo.Context, serviceURL string) error {
	req := c.Request()
	res := c.Response()

	target, err := url.Parse(serviceURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Invalid service URL",
		})
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(res, req)
	return nil
}
