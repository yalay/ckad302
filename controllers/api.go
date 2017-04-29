package controllers

import (
	"conf"

	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type AjaxRsp struct {
	Success int32  `json:"success"`
	Msg     string `json:"msg"`
	DMsg    string `json:"d_msg"`
	DTxt    string `json:"d_txt"`
	DLink   string `json:"d_link"`
}

func GetDlLinks(c echo.Context) error {
	articleId := Atoi32(c.Param("id"))
	if articleId == 0 {
		return fmt.Errorf("articleId empty")
	}

	ajaxRsp := &AjaxRsp{
		Success: 1,
	}
	adLinks := conf.GetArticleAdLinks(articleId)
	if len(adLinks) == 0 {
		return c.JSON(http.StatusOK, AjaxRsp{
			Msg: "资源还未上传，请稍等片刻",
		})
	}

	ip := c.RealIP()
	ckUrl := GetCkLeastUrl(ip, adLinks)
	ajaxRsp.DLink = ckUrl
	return c.JSON(http.StatusOK, AjaxRsp{
		Success: 1,
		Msg:     "vip资源免费下载",
		DMsg:    "zip压缩包",
		DTxt:    "非会员有广告页面，会提示等待几秒后出现跳过广告、GET LINK或者Continue等按钮后点击就可以进入网盘(mail.ru)下载",
		DLink:   ckUrl,
	})
}
