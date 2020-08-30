package v1

import (
	error_msg "blog/error"
	"blog/middleware"
	"blog/model"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterResult struct {
	Token string `json:"token"`
	User User `json:"user"`
}

type User struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Tel string `json:"tel"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
}

func Register(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": error_msg.ERROR_REQ_PARAM_ERROR,
			"msg": error_msg.GetErrorMsg(error_msg.ERROR_REQ_PARAM_ERROR)+err.Error(),
		})
		return
	}
	if !model.UserIsExist(user.Username) {
		avatar, err := utils.UploadAvatar(&user.Avatar)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": error_msg.ERROR_UPLOAD_AVATAR_ERROR,
				"msg": error_msg.GetErrorMsg(error_msg.ERROR_UPLOAD_AVATAR_ERROR)+err.Error(),
			})
			return
		}
		user.Avatar = avatar
		if model.CreateUser(&user) == error_msg.SUCCESS {
			token, code := middleware.CreateToken(user.Username)
			if code == error_msg.SUCCESS {
				u := User{
					user.ID,
					user.Username,
					user.Tel,
					user.Email,
					user.Avatar,
				}
				c.JSON(http.StatusOK, gin.H{
					"status": code,
					"msg": error_msg.GetErrorMsg(code),
					"data":	RegisterResult{
						Token: token,
						User: u,
					},
				})
			}
		}else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": error_msg.ERROR,
				"msg": error_msg.GetErrorMsg(error_msg.ERROR),
			})
		}
	}else {
		c.JSON(http.StatusForbidden, gin.H{
			"status": error_msg.ERROR_USER_IS_EXIST,
			"msg": error_msg.GetErrorMsg(error_msg.ERROR_USER_IS_EXIST),
		})
	}
}
