package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Mobilizes/materi-be-alpro/modules/user/service"
    "github.com/Mobilizes/materi-be-alpro/modules/user/validation"
    "github.com/Mobilizes/materi-be-alpro/pkg/utils"
)

type UserController struct {
    service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
    return &UserController{service: service}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
    req, err := validation.ValidateCreateUser(c)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    user, err := ctrl.service.CreateUser(req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat user")
        return
    }

    utils.SuccessResponse(c, http.StatusCreated, "User berhasil dibuat", user)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
        return
    }

    user, err := ctrl.service.GetUserByID(uint(id))
    if err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "User tidak ditemukan")
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "OK", user)
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
    users, err := ctrl.service.GetAllUsers()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data")
        return
    }
    utils.SuccessResponse(c, http.StatusOK, "OK", users)
}
