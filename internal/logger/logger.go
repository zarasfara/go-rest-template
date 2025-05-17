package logger

import (
	"log"
	"log/slog"
	"os"
)

func InitLogger(env string) {
	var handler slog.Handler

	if env == "development" {
		f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		// Ensure the file handle remains open for the application's lifetime.
		// Remember to close it explicitly when the application shuts down.

		handler = slog.NewTextHandler(f, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
	} else {
	}

	slog.SetDefault(slog.New(handler))
}