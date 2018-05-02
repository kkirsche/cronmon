package libserver

import (
	"strconv"
	"strings"

	"github.com/kkirsche/cronmon/libmetrics"
	"github.com/labstack/echo"
)

// MetricsMiddleware holds the mutex so we can do middleware functions
type MetricsMiddleware struct {
}

// NewMetricsMiddleware is used to generate a new metrics middleware object for us to use.
func newMetricsMiddleware() *MetricsMiddleware {
	return &MetricsMiddleware{}
}

// GeneratePrometheusMetrics is the actual middleware function
func (m *MetricsMiddleware) GeneratePrometheusMetrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		libmetrics.HTTPRequestsTotal.WithLabelValues(
			strings.ToLower(c.Request().Method),
			c.Request().RequestURI,
			strconv.Itoa(c.Response().Status)).Inc()

		libmetrics.HTTPResponseSize.Observe(float64(c.Response().Size))

		return nil
	}
}
