package handler

import (
	"gin-socmed/dto"
	"gin-socmed/errorhandler"
	"gin-socmed/helper"
	"gin-socmed/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

// Register godoc
// @Summary Register user
// @Description Melakukan registrasi pengguna
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Request body"
// @Success 200 {object} dto.ResponseParams
// @Router /auth/register [post]
func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register Successfully!, please login",
	})

	c.JSON(http.StatusCreated, res)
}

// Login godoc
// @Summary Login user
// @Description Melakukan autentikasi pengguna
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Request body"
// @Success 200 {object} dto.LoginResponse
// @Router /auth/login [post]
func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully Login",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)
}
