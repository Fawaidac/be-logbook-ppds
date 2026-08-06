package auth

import (
	"net/http"

	"be-logbook-ppds/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid (username & password wajib diisi)")
		return
	}

	res, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.SuccessWithToken(c, http.StatusOK, "Login Success", res.User, res.Token)
}

func (h *Handler) Logout(c *gin.Context) {
	response.Success(c, http.StatusOK, "Logout Success", nil)
}

func (h *Handler) Me(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userID, ok := userIDVal.(int)
	if !ok || userID == 0 {
		// Fallback untuk user demo jika login tanpa DB ID
		username := c.GetString("username")
		role := c.GetString("role")
		response.Success(c, http.StatusOK, "Get profile success", gin.H{
			"username": username,
			"role":     role,
		})
		return
	}

	profile, err := h.service.GetProfile(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Get profile success", profile)
}

