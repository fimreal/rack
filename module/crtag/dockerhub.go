package crtag

import (
	"encoding/json"
	"net/http"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
)

func ListDockerhubTags(c *gin.Context) {
	ns := c.Param("namespace")
	repo := c.Param("repository")
	result := c.Param("result")
	rawQuery := c.Request.URL.RawQuery

	httpcode, b, err := listRepoTags(ns, repo, rawQuery)
	if err != nil {
		ezap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get tags info", "detail": err.Error()})
		return
	}

	// 解析 dockerhub 返回错误
	if httpcode == 403 || httpcode == 404 {
		var e = &ErrorResp{}
		err := json.Unmarshal(b, &e)
		if err != nil {
			c.JSON(httpcode, e)
			return
		}
	}

	switch result {
	case "/image", "/images":
		res := &ImagesInfo{}
		json.Unmarshal(b, &res)
		c.JSON(http.StatusOK, res)
	case "/tag", "/tags":
		res := &TagsInfo{}
		json.Unmarshal(b, &res)
		c.JSON(http.StatusOK, res)
	default:
		res := &ListTagsInfo{}
		json.Unmarshal(b, &res)
		c.JSON(http.StatusOK, res)
	}
}
