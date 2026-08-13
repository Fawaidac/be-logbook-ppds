package tindakan

import (
	"net/http"
	"strconv"

	"be-logbook-ppds/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetSummary(c *gin.Context) {
	userID := 0
	if val, exists := c.Get("user_id"); exists {
		if id, ok := val.(int); ok {
			userID = id
		}
	}

	summary, err := h.service.GetSummary(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil rekap logbook tindakan", summary)
}

func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tindakan, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil detail tindakan", tindakan)
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateTindakanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	userID := 0
	if val, exists := c.Get("user_id"); exists {
		if id, ok := val.(int); ok {
			userID = id
		}
	}

	res, err := h.service.Create(c.Request.Context(), req, userID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"message":  "Tindakan berhasil disimpan. Silakan klik 'Kirim' untuk mengirim ke approval DPJP.",
		"tindakan": res,
	})
}

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req UpdateTindakanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	res, err := h.service.Update(c.Request.Context(), id, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "Tindakan berhasil diperbarui.",
		"tindakan": res,
	})
}

func (h *Handler) Send(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.service.Send(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Anda berhasil dikirim ke DPJP untuk verifikasi.",
	})
}

func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Tindakan berhasil dihapus.",
	})
}