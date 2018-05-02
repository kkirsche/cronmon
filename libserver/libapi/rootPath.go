package libapi

import "github.com/labstack/echo"

// RootPath returns for the / method showing a list of all routes
func RootPath(c echo.Context) error {
	resp := `<html>
<head><title>Cronmon API</title></head>
<body>
<h1>Cron Monitoring</h1>
<h2>Metrics</h2>
<p><a href='/metrics'>Metrics</a></p>
<h2>Version 1 API</h2>
<p>TBD</p>
</html>
`
	return c.HTML(200, resp)
}
