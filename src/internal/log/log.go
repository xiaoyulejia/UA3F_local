package log

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sunbk201/ua3f/internal/config"
	"github.com/sunbk201/ua3f/internal/paths"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetLogConf(level string) {
	primaryWriter := os.Stdout

	var writers []io.Writer
	writers = append(writers, primaryWriter)

	if logDir, err := paths.EnsureLogDir(); err == nil {
		fileWriter := &lumberjack.Logger{
			Filename:   filepath.Join(logDir, "ua3f.log"),
			MaxSize:    5, // megabytes
			MaxBackups: 5,
			MaxAge:     7, // days
			LocalTime:  true,
			Compress:   true,
		}
		writers = append(writers, fileWriter)
	} else {
		fmt.Fprintf(os.Stderr, "UA3F: failed to prepare log directory %q: %v\n", paths.LogDir(), err)
	}

	multiWriter := io.MultiWriter(writers...)

	var logLevel slog.Level
	switch strings.ToLower(level) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	loc := LoadLocalLocation()
	opts := &slog.HandlerOptions{
		Level: logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := a.Value.Time().In(loc)
				return slog.String(slog.TimeKey, t.Format("2006-01-02 15:04:05"))
			}
			return a
		},
	}
	logger := slog.New(slog.NewTextHandler(multiWriter, opts))
	slog.SetDefault(logger)
}

func LogHeader(version string, cfg *config.Config) {
	slog.Info("UA3F started", "version", version, "", cfg)
	slog.Info("OS Info", GetOSInfo()...)
}

func LogDebugWithAddr(src string, dest string, msg string) {
	slog.Debug(msg, slog.String("src", src), slog.String("dest", dest))
}

func LogInfoWithAddr(src string, dest string, msg string) {
	slog.Info(msg, slog.String("src", src), slog.String("dest", dest))
}

func LogWarnWithAddr(src string, dest string, msg string) {
	slog.Warn(msg, slog.String("src", src), slog.String("dest", dest))
}

func LogErrorWithAddr(src string, dest string, msg string) {
	slog.Error(msg, slog.String("src", src), slog.String("dest", dest))
}

// LoadLocalLocation tries to detect and load the system local timezone from
// `/etc/localtime` or `/etc/TZ`. Compatible with OpenWrt and normal Linux.
func LoadLocalLocation() *time.Location {
	if _, err := os.Stat("/etc/localtime"); err == nil {
		if loc, _ := time.LoadLocation("Local"); loc != nil {
			return loc
		}
	}
	if data, err := os.ReadFile("/etc/TZ"); err == nil {
		tz := trim(string(data))
		if len(tz) > 0 {
			if strings.HasPrefix(tz, "CST-8") {
				return time.FixedZone("CST", 8*3600)
			}
			if strings.HasPrefix(tz, "UTC") {
				return time.UTC
			}
		}
	}
	return time.UTC
}

func trim(s string) string {
	for len(s) > 0 && (s[len(s)-1] == '\n' || s[len(s)-1] == '\r' || s[len(s)-1] == ' ') {
		s = s[:len(s)-1]
	}
	for len(s) > 0 && (s[0] == '\n' || s[0] == '\r' || s[0] == ' ') {
		s = s[1:]
	}
	return s
}
