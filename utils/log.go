package utils

import (
	"log/slog"
	"os"

)

var Loger *slog.Logger

func init() {
	Loger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
