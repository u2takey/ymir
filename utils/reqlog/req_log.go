package reqlog

import (
	"encoding/base64"
	"encoding/binary"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

var (
	HeaderReqID = "X-Reqid"
	KeyReqID    = "reqid"
)
var pid = uint32(time.Now().UnixNano() % 4294967291)

func genReqId() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}

func ReqLoggerMiddleware(logger *logrus.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		reqid := c.Request.Header.Get("X-Reqid")
		if reqid == "" {
			reqid = genReqId()
			c.Request.Header.Set("X-Reqid", reqid)
		}

		path := c.Request.URL.Path
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		entry := logger.WithFields(logrus.Fields{
			"reqid":      reqid,
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
			"time":       end.Format(timeFormat),
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info()
		}
	}
}

// usage: Entry(c).Debug(".....")
func Entry(c context.Context) *logrus.Entry {
	if ctx, ok := c.(*gin.Context); ok {
		return logrus.WithField(KeyReqID, ctx.Request.Header.Get(HeaderReqID))
	} else {
		return logrus.WithField(KeyReqID, c.Value(KeyReqID))
	}
}

func WithReqID(ctx context.Context, reqid string) context.Context {
	return context.WithValue(ctx, "reqid", reqid)
}

func Context(c *gin.Context) (ctx context.Context) {
	ctx = context.Background()
	reqID := c.Request.Header.Get(HeaderReqID)
	ctx = WithReqID(ctx, reqID)
	return
}

//---------------------------------------------------------------------------------
