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

func GetVideos() echo.HandlerFunc {
	return func (c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)
		videoId := c.Param("id")
		lp := []string{"id","snippet"}
		call := yts.Videos.List(lp).Id(videoId)
		res, err := call.Do()

		if err != nil {
			logrus.Fatalf("ErrorcallingYoutubeAPI:%v",err)
		}
		v := VideoResponse {
			VideoList: res,
		}

		return c.JSON(fasthttp.StatusOK, v)
	}
}
