package v1

import (
	error_msg "blog/error"
	"blog/middleware"
	"blog/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginData struct {
	Username string `from:"username" binding:"required,UsernameValidator"`
	Password string `from:"password" binding:"required"`
}

type LoginResult struct {
	Token string `json:"token"`
	User User `json:"user"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H {
			"status": error_msg.ERROR_REQ_PARAM_ERROR,
			"msg": error_msg.GetErrorMsg(error_msg.ERROR_REQ_PARAM_ERROR) + err.Error(),
		})
		return
	}
	if model.UserIsExist(data.Username) {
		user, code := model.CheckPwd(data.Username, data.Password)
		if code == error_msg.SUCCESS {
			token, code2 := middleware.CreateToken(user.Username)
			if code2 == error_msg.SUCCESS {
				u := User{
					user.ID,
					user.Username,
					user.Tel,
					user.Email,
					user.Avatar,
				}
				c.JSON(http.StatusOK, gin.H{
					"status": error_msg.SUCCESS,
					"msg": error_msg.GetErrorMsg(error_msg.SUCCESS),
					"data": LoginResult{
						Token: token,
						User: u,
					},
				})
			}else{
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": error_msg.ERROR,
					"msg": error_msg.GetErrorMsg(error_msg.ERROR),
				})
			}
		}else{
			c.JSON(http.StatusBadRequest, gin.H{
				"status": code,
				"msg": error_msg.GetErrorMsg(code),
			})
		}
	}else {
		c.JSON(http.StatusForbidden, gin.H{
			"status": error_msg.ERROR_USER_NOT_EXIST,
			"msg": error_msg.GetErrorMsg(error_msg.ERROR_USER_NOT_EXIST),
		})
	}
}