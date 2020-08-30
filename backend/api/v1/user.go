package v1

import (
	error_msg "blog/error"
	"blog/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": error_msg.ERROR_REQ_PARAM_ERROR,
			"msg": error_msg.GetErrorMsg(error_msg.ERROR_REQ_PARAM_ERROR),
		})
		return
	}
	var user model.UpdateUserData
	err = c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": error_msg.ERROR_REQ_PARAM_ERROR,
			"msg": error_msg.GetErrorMsg(error_msg.ERROR_REQ_PARAM_ERROR)+err.Error(),
		})
		return
	}
	if model.UsernameIsExist(uint(id) ,user.Username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": error_msg.ERROR_USERNAME_IS_EXIST,
			"msg": error_msg.GetErrorMsg(error_msg.ERROR_USERNAME_IS_EXIST),
		})
	}else{
		code := model.UpdateUser(id, &user)
		if code == error_msg.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"status": error_msg.SUCCESS,
				"msg": error_msg.GetErrorMsg(error_msg.SUCCESS),
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"status": error_msg.ERROR,
				"msg": error_msg.GetErrorMsg(error_msg.ERROR),
			})
		}
	}
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": error_msg.ERROR_REQ_PARAM_ERROR,
			"msg": error_msg.GetErrorMsg(error_msg.ERROR_REQ_PARAM_ERROR),
		})
		return
	}
	if model.DeleteUser(id) == error_msg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": error_msg.SUCCESS,
			"msg": error_msg.GetErrorMsg(error_msg.SUCCESS),
		})
	}else{
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": error_msg.ERROR,
			"msg": error_msg.GetErrorMsg(error_msg.ERROR),
		})
	}
}