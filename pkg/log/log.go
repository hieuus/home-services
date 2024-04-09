package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

func New() zerolog.Logger {
	out := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.DateTime,
	}

	out.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	out.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	out.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	out.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	log := zerolog.New(out).With().Timestamp().Logger()

	return log
}
