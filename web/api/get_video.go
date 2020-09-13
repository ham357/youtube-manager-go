package api

import (
	"youtube-manager-go/middlewares"
	"youtube-manager-go/models"

	"github.com/ham357/youtube-manager-gomiddlewares"
	"github.com/ham357/youtube-manager-gomodels"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

type VideoResponse struct {
	VideoList *youtube.VideoListResponse `json:"video_list"`
	IsFavorite bool                       `json:"is_favorite"`
}

func GetVideos() echo.HandlerFunc {
	return func (c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		token := c.Get("auth").(*auth.Token)
		videoId := c.Param("id")
		isFavorite := false
		if token != nil {
			favorite := models.Favorite{}
			isFavoriteNotFound := dbs.DB.Table("favorites").
				Joins("INNER JOIN users ON users.id = favorites.user_id").
				Where(models.User{UID: token.UID}).
				Where(models.Favorite{VideoId: videoId}).
				First(&favorite).
				RecordNotFound()

			logrus.Debug("isFavoriteNotFound: ", isFavoriteNotFound)
			if !isFavoriteNotFound {
				isFavorite = true
			}
		}
		lp := []string{"id","snippet"}
		call := yts.Videos.List(lp).Id(videoId)
		res, err := call.Do()

		if err != nil {
			logrus.Fatalf("ErrorcallingYoutubeAPI:%v",err)
		}
		v := VideoResponse {
			VideoList: res,
			IsFavorite: isFavorite,
		}

		return c.JSON(fasthttp.StatusOK, v)
	}
}
