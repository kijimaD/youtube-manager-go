package middlewares

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"youtube-manager/databases"
)

type DatabaseClient struct {
	DB *gorm.DB
}

func DatabaseService() echo.MiddlewareFunc {
	return func(next ehco.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := databases.Connect()
			d := DatabaseClient{DB: session}

			defer d.DB.close()

			// デバッグ用に実行したSQLをログに出力する
			d.DB.LogMode(true)

			c.Set("dbs", &d)

			if err := next(c); err != nil {
				return nil
			}

			return nil
		}
	}
}
