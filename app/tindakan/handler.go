package tindakan

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

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
	userName := loggedName(c)
	userRole := ""
	if val, exists := c.Get("role"); exists {
		if role, ok := val.(string); ok {
			userRole = role
		}
	}

	department := c.Query("department")
	if unescaped, err := url.QueryUnescape(department); err == nil {
		department = unescaped
	}
	department = strings.ReplaceAll(department, "+", " ")
	department = strings.TrimSpace(department)

	summary, err := h.service.GetSummary(c.Request.Context(), userName, userRole, department)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil rekap logbook tindakan", summary)
}

func (h *Handler) GetDPJP(c *gin.Context) {
	programStudi := c.Query("program_studi")
	dpjp, err := h.service.GetDPJP(c.Request.Context(), programStudi)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil daftar DPJP", dpjp)
}

func (h *Handler) GetByDepartment(c *gin.Context) {
	department := c.Query("department")
	if department == "" {
		response.Error(c, http.StatusBadRequest, "Parameter department diperlukan")
		return
	}
	data, err := h.service.GetByDepartment(c.Request.Context(), department)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Berhasil mengambil data tindakan per departemen", data)
}

func (h *Handler) GetBySupervisor(c *gin.Context) {
	supervisorName := c.Param("name")
	if supervisorName == "" {
		response.Error(c, http.StatusBadRequest, "Parameter supervisor name diperlukan")
		return
	}
	if unescaped, err := url.QueryUnescape(supervisorName); err == nil {
		supervisorName = unescaped
	}
	supervisorName = strings.ReplaceAll(supervisorName, "+", " ")
	supervisorName = strings.TrimSpace(supervisorName)

	division := c.Query("division")
	if unescaped, err := url.QueryUnescape(division); err == nil {
		division = unescaped
	}
	division = strings.TrimSpace(division)

	data, err := h.service.GetBySupervisor(c.Request.Context(), supervisorName, division)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	if data == nil {
		data = []TindakanResponse{}
	}
	response.Success(c, http.StatusOK, "Berhasil mengambil data tindakan per supervisor", data)
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

	userName := loggedName(c)

	res, err := h.service.Create(c.Request.Context(), req, userName)
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

func loggedName(c *gin.Context) string {
	if val, exists := c.Get("name"); exists {
		if name, ok := val.(string); ok && name != "" {
			return name
		}
	}
	if val, exists := c.Get("username"); exists {
		if username, ok := val.(string); ok {
			return username
		}
	}
	return ""
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
