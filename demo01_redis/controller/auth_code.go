package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	. "redis/config"
	Email "redis/email"
	"redis/utils"
	"time"
)

// AuthCodeSave 发送验证码
func AuthCodeSave(c *gin.Context) {
	email := c.Query("email")
	if !utils.VerifyEmailFormat(email) {
		c.JSON(501, gin.H{
			"msg": "邮箱地址不合法",
		})
		return
	}
	var times = email + "_times"
	var expires = email + "_expires"
	//设置email每天发送次数
	if val := RedisCli.Exists(times).Val(); val == 0 {
		//设置过期时间
		RedisCli.Set(times, 0, time.Duration(24*60-time.Now().Hour()*60+time.Now().Minute())*time.Minute)
	} else if num, _ := RedisCli.Get(times).Int(); num > 3 {
		c.JSON(502, gin.H{
			"msg": "当天发送次数到达上限",
		})
		return
	}
	//生成验证码
	rand.Seed(time.Now().Unix())
	code := rand.Intn(int(1e6)-int(1e5)) + int(1e5)
	//发送验证码
	if err := Email.Send("来自磐石公司的验证码", []byte("验证码为："+fmt.Sprintf("%d", code)), email); err != nil {
		c.JSON(503, gin.H{
			"msg": "发送验证码失败",
		})
		return
	}
	//更新redis数据
	RedisCli.Set(expires, code, 2*time.Minute)
	RedisCli.Incr(times)
	c.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}

// AuthCodeJudge 判断验证码
func AuthCodeJudge(c *gin.Context) {
	email := c.Query("email")
	code := c.Query("code")
	var expires = email + "_expires"
	if RedisCli.Exists(expires).Val() != 0 && RedisCli.Get(expires).Val() == code {
		c.JSON(http.StatusOK, gin.H{
			"msg": "验证成功",
		})
		return
	}
	c.JSON(501, gin.H{
		"msg": "验证失败",
	})
}
