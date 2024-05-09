package log

import (
	"runtime"
	"strconv"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func WithContext(c *gin.Context) *logrus.Entry {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	requestID := requestid.Get(c)
	return logrus.WithFields(
		logrus.Fields{
			"location":   frame.File + ":" + strconv.Itoa(frame.Line) + " " + frame.Function,
			"request_id": requestID,
			"user_ip":    c.ClientIP(),
		},
	)
}
