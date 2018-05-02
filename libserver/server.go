package libserver

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/kkirsche/cronmon/libserver/libapi"
	"github.com/kkirsche/cronmon/libserver/libapi/apiv1"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

// Run is used to run the web server and listen for a graceful shutdown
// requirement
func Run(listenPort int) error {
	listenPort, err := retrieveListeningPortFromEnvironment(listenPort)
	if err != nil {
		return err
	}
	e := echo.New()
	e.Logger.Info("registering middleware...")
	registerMiddleware(e)
	e.Logger.Info("registering routes...")
	registerRoutes(e)

	// Start the server in the background so we can listen for the shutdown
	// command
	e.Logger.Info("starting server...")
	go func() {
		err := e.Start(fmt.Sprintf(":%d", listenPort))
		if err != nil {
			e.Logger.Info("server has shut down")
		}
	}()

	// Listen for the shutdown / interrupt call
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	e.Logger.Info("shutting down the server...")
	// Start the shutdown with a 10 second maximum wait time
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}

func registerMiddleware(e *echo.Echo) {
	m := newMetricsMiddleware()
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(m.GeneratePrometheusMetrics)
}

func registerRoutes(e *echo.Echo) {
	// Setup base handlers
	e.GET("/", libapi.RootPath)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	api := e.Group("/api")

	// Setup api version 1
	v1 := api.Group("/v1")
	// Tasks APIs
	v1.GET("/tasks", apiv1.GetTasks)
	v1.POST("/tasks", apiv1.CreateTask)
	v1.GET("/tasks/:id", apiv1.GetTask)
	v1.PUT("/tasks/:id", apiv1.UpdateTask)
	v1.DELETE("/tasks/:id", apiv1.DeleteTask)

	v1.GET("/tasks/:access_id/run", apiv1.StartTask)
	v1.GET("/tasks/:access_id/complete", apiv1.CompleteTask)
}

// retrieveListeningPortFromEnvironment is used to take a default listening port
// number, check if one was set in the environment variable DS_LISTEN_PORT and
// then set it if detected
func retrieveListeningPortFromEnvironment(listenPort int) (int, error) {
	var err error
	p := os.Getenv("LISTEN_PORT")
	if p != "" {
		logrus.Infof("detected LISTEN_PORT of %s", p)
		listenPort, err = strconv.Atoi(p)
		if err != nil {
			logrus.WithError(err).Errorln("failed to parse LISTEN_PORT. exiting...")
			return listenPort, err
		}
	}

	if listenPort < 1 || listenPort > 65535 {
		logrus.Errorf("port %d is not a valid port number. exiting...", listenPort)
		return listenPort, fmt.Errorf("port %d is not a valid port number. exiting", listenPort)
	}

	return listenPort, nil
}
