package api

import (
	"context"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"os"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		key := os.Getenv("API_KEY")

		ctx := context.Background()
		yts, err := youtube.NewService(ctx, option.WithAPIKey(key))

		if err != nil {
			logrus.Fatalf("Error creating new YouTube service: %v", err)
		}

		call := yts.Videos.
			List([]string{"id", "snippet"}).
			Chart("mostPopular").
			MaxResults(3)

		// クエリパラメータにpageTokenが設定されている場合 、API呼び出し時にPageToken()を使う
		pageToken := c.QueryParam("pageToken")
		if len(pageToken) > 0 {
			call = call.PageToken(pageToken)
		}

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("error calling YouTube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
