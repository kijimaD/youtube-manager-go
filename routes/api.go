package routes

import (
	"github.com/labstack/echo"
	"youtube-manager/middlewares"
	"youtube-manager/web/api"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
		g.GET("/video/:id", api.GetVideo())
		g.GET("/related/:id", api.FetchRelatedVideos())
		g.GET("/search", api.SearchVideos())
	}

	fg := g.Group("/favorite", middlewares.FirebaseGuard()) // /api/favorite配下のルーティングにアクセスする場合は、Authミドルウェアを適用する
	{
		fg.POST("/:id/toggle", api.ToggleFavoriteVideo())
	}
}
