package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)
var loggers *zap.Logger
var sugarLogger *zap.SugaredLogger
// InitLogger 初始化Logger
//func Init() (err error) {
//	writeSyncer := getLogWriter(
//		viper.GetString("log.filename"),
//		viper.GetInt("log.max_size"),
//		viper.GetInt("log.max_backups"),
//		viper.GetInt("log.max_age"),
//	) //获取配置文件
//	encoder := getEncoder()
//	var l = new(zapcore.Level)
//	err = l.UnmarshalText([]byte(viper.GetString("log.level")))
//	if err != nil {
//		return
//	}
//	core := zapcore.NewCore(encoder, writeSyncer, l)
//
//	lg := zap.New(core, zap.AddCaller())
//	zap.ReplaceGlobals(lg)
//	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
//	return
//}

func Init() (err error) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	loggers = zap.New(core, zap.AddCaller())
	sugarLogger = loggers.Sugar()
		var l = new(zapcore.Level)
		err = l.UnmarshalText([]byte(viper.GetString("log.level")))
		if err != nil {
			return
		}
		return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder, //当前时间
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log", //文件名
		MaxSize:    1,            //大小 单位是M
		MaxBackups: 5,            //备份数量
		MaxAge:     30,           //最大备份天数
		Compress:   false,        //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}


// GinLogger 接收gin框架默认的日志
func GinLogger(loggers *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		loggers.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()

	}
}
