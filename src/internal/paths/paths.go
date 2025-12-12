package paths

import (
	"os"
	"path/filepath"
	"runtime"
)

const logDirEnv = "UA3F_LOG_DIR"

// LogDir returns the base directory for UA3F logs/stat files.
func LogDir() string {
	if dir := os.Getenv(logDirEnv); dir != "" {
		return dir
	}

	switch runtime.GOOS {
	case "linux":
		return "/var/log/ua3f"
	case "darwin":
		if home, err := os.UserHomeDir(); err == nil {
			return filepath.Join(home, "Library", "Logs", "ua3f")
		}
	default:
		if home, err := os.UserHomeDir(); err == nil {
			return filepath.Join(home, ".ua3f", "logs")
		}
	}

	return filepath.Join(os.TempDir(), "ua3f-logs")
}

// EnsureLogDir creates the log directory if it does not exist.
func EnsureLogDir() (string, error) {
	dir := LogDir()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return dir, nil
}
