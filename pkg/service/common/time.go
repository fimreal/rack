package common

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeStamp(c *gin.Context) {
	t := time.Now()
	tzone, _ := t.Zone()
	if _, ok := c.GetQuery("j"); ok {
		c.JSON(http.StatusOK, gin.H{"time": strconv.FormatInt(t.Unix(), 10)})
		return
	}
	if _, ok := c.GetQuery("h"); ok {
		c.String(http.StatusOK, "%s %s", t.Format("2006-01-02 15:04:05"), tzone)
		// c.String(http.StatusOK, "%s %s", t.Format("2006-01-02 15:04:05"), tzone)
		return
	}

	c.String(http.StatusOK, strconv.FormatInt(t.Unix(), 10))
}

func TimeStampTrans(c *gin.Context) {
	rt := c.Param("ts")

	if t, err := strconv.ParseInt(rt, 10, 64); err == nil {
		c.String(http.StatusOK, time.Unix(t, 0).String())
		return
	}
	if t, err := time.Parse("2006-01-02T15:04:05", rt); err == nil {
		c.String(http.StatusOK, "%d", t.Unix())
		return
	}
	if t, err := time.Parse(time.RFC3339, rt); err == nil {
		c.String(http.StatusOK, "%d", t.Unix())
		return
	}
	if t, err := time.Parse("Mon Jan 2 15:04::05 UTC 2006", rt); err == nil {
		c.String(http.StatusOK, "%d", t.Unix())
		return
	}
	// Feb 3, 2013 at 7:54pm (PST)

	c.String(601, `出错了! 无法识别输入时间格式: %s

时间戳 <=> RFC3339 互转，默认为 UTC 时间
支持格式参考:
	1. 时间戳, 1678696441
	2. RFC3339 格式(gunlinux date --iso-8601=seconds), 2006-01-02T15:04:05+07:00
	3. 默认 UTC 时区, 2006-01-02T15:04:05`, rt)
}
