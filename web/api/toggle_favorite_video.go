package api

import (
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"youtube-manager/middlewares"
	"youtube-manager/models"
)

type ToggleFavoriteVideoResponse struct {
	VideoId    string `json:"video_id"`
	IsFavorite bool   `json:"is_favorite"`
}

func ToggleFavoriteVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		videoId := c.Param("id")
		token := c.Get("auth").(*auth.Token)
		user := models.User{}
		if dbs.DB.Table("users").
			Where(models.User{UID: token.UID}).First(&user).RecordNotFound() { // usersテーブルにトークンから取得したUIDを保持するレコードが存在しているか確認レコードが見つからなかった場合は、新規ユーザとしてusersテーブルにレコードを作成する。サインアップ時点ではfirebase側でユーザができているだけなので、userは存在してない可能性がある。
			user = models.User{UID: token.UID}
			dbs.DB.Create(&user)
		}

		favorite := models.Favorite{}
		isFavorite := false
		if dbs.DB.Table("favorites").
			Where(models.Favorite{UserId: user.ID, VideoId: videoId}).
			First(&favorite).RecordNotFound() {
			favorite = models.Favorite{UserId: user.ID, VideoId: videoId}
			dbs.DB.Create(&favorite) // お気に入り追加
			isFavorite = true
		} else {
			dbs.DB.Delete(&favorite) // お気に入り削除
		}

		res := ToggleFavoriteVideoResponse{
			VideoId:    videoId,
			IsFavorite: isFavorite,
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
