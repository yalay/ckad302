package controllers

import (
	"conf"

	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func GetDlLinks(c echo.Context) error {
	articleId := Atoi32(c.Param("id"))
	if articleId == 0 {
		return fmt.Errorf("articleId empty")
	}
	adLinks := conf.GetArticleAdLinks(articleId)
	if len(adLinks) == 0 {
		return fmt.Errorf("no dl links")
	}

	ip := c.RealIP()
	ckUrl := GetCkLeastUrl(ip, adLinks)
	c.String(http.StatusOK, ckUrl)
	return nil
}
