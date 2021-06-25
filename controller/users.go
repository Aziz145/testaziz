package controller

import (
	
	"net/http"

	//"os"

	"github.com/Aziz145/testaziz/tree/master/model"
	"github.com/Aziz145/testaziz/tree/master/util"
	"golang.org/x/crypto/bcrypt"

	"crypto/rand"


	"github.com/gin-gonic/gin"
)
func (idb *InDB) CreateUser(ctx *gin.Context){
	var (
		request model.Users
	)
	request.Password = util.PasswordGeneratoer()
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Bad Password")
		ctx.Abort()
		return
	}
	request.Password = string(password)
	// user.Status.Code = uuid.NewV4().String()
	b := make([]byte, 32) //equals 8 charachters
	rand.Read(b)


	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please Check Your Data")
		ctx.Abort()
		return
	}

	err = idb.DB.Create(&request).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Create Users")
		ctx.Abort()
		return
	}
	util.ResponseSuccessCustomMessage(ctx, http.StatusOK, "Users Has Been Created")

}
func (idb *InDB) GetUserLimit(ctx *gin.Context){
	var result = model.Users{}
	
	err := idb.DB.Limit(1).Offset(1).Find(&result).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Find User")
		ctx.Abort()
		return
	}
	util.ResponseSuccess(ctx, http.StatusOK, result)
}
func (idb *InDB) GetUser(ctx *gin.Context) {
	var result = model.Users{}
	id := ctx.Param("id")

	err := idb.DB.Table("users").Where("id = ?", id).First(&result).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Find Users")
		ctx.Abort()
		return
	}
	util.ResponseSuccess(ctx, http.StatusOK, result)
}

func (idb *InDB) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var (
		user     model.Users
		new_user model.Users
	)

	err := ctx.ShouldBindJSON(&new_user)
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please Check Your Data")
		ctx.Abort()
		return
	}

	err = idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please Check Your Data")
		ctx.Abort()
		return
	}

	err = idb.DB.Model(&user).Where("id = ?", id).Updates(new_user).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Update Failed")
		ctx.Abort()
		return
	}
	util.ResponseSuccessCustomMessage(ctx, http.StatusOK, "User Has Been Updated")
}

func (idb *InDB) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user model.Users

	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please Check Your Data")
		ctx.Abort()
		return
	}

	err = idb.DB.Model(&user).Where("id= ?", id).Delete(&user).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Update Failed")
		ctx.Abort()
		return
	}

	util.ResponseSuccessCustomMessage(ctx, http.StatusOK, "User Has Been Deleted")
}