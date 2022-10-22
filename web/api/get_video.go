package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

type VideoResponse struct {
	VideoList *youtube.VideoListResponse `json:"video_list"`
}

func GetVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service) // youtube.Serviceはコンテキストから取得

		videoId := c.Param("id") // 動画のidをパラメータから取得して、APIリクエスト時に使う

		call := yts.Videos.
			List([]string{"id", "snippet"}).
			Id(videoId)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling YouTube API: %v", err)
		}

		// YouTube APIからのレスポンスをそのままnuxtに返却するのではなく、VideoResponseという名前の構造体に詰めて返却する
		v := VideoResponse{
			VideoList: res,
		}

		return c.JSON(fasthttp.StatusOK, v)
	}
}
