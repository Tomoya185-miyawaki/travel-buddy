package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var layout = "2006-01-02"

func Log() echo.MiddlewareFunc {
	now := time.Now()
	strDate := now.Format(layout)
	file, err := os.OpenFile(fmt.Sprintf("log/%s.log", strDate), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		// ファイルが読み込めない場合は標準出力にする
		return middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		})
	}

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: file,
	})
}
