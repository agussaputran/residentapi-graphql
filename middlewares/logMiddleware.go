package middlewares

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// LogRequest for logging request to senrty.io
func LogRequest(c *gin.Context) string {
	reqMethod := c.Request.Method
	reqPath := c.Request.URL.Path
	buf, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	return string(reqMethod) + " -> " + reqPath
}

// LogRequestBody for logging request to senrty.io
func LogRequestBody(c *gin.Context) string {
	reqMethod := c.Request.Method
	reqPath := c.Request.URL.Path
	buf, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	return string(reqMethod) + " -> " + reqPath + "\n" + string(buf)
}

// Sentry func to log with sentry.io
func Sentry(data string) {
	dsn := os.Getenv("DSN")

	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dsn,
	})
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage(data)
}
