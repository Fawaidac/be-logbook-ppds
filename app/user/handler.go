package user

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"be-logbook-ppds/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	userRes, err := h.service.CreateUser(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Pengguna berhasil dibuat", userRes)
}

func (h *Handler) FindAll(c *gin.Context) {
	users, err := h.service.GetAllUsers(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil daftar pengguna", users)
}

func (h *Handler) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	userRes, err := h.service.GetUserByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil detail pengguna", userRes)
}

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	userRes, err := h.service.UpdateUser(c.Request.Context(), id, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Data pengguna berhasil diperbarui", userRes)
}

func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.service.DeleteUser(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Pengguna berhasil dihapus", nil)
}

func (h *Handler) Register(c *gin.Context) {
	var req CreateRegistrationRequest
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input registrasi tidak valid: "+err.Error())
		return
	}

	selfiePath := ""
	strPath := ""
	sipPath := ""

	_ = os.MkdirAll("./uploads", 0755)

	nowNano := time.Now().UnixNano()

	if file, err := c.FormFile("selfie"); err == nil && file != nil {
		ext := filepath.Ext(file.Filename)
		cleanName := strings.TrimSuffix(file.Filename, ext)
		cleanName = strings.ReplaceAll(cleanName, " ", "_")
		uniqueName := fmt.Sprintf("selfie_%d_%s%s", nowNano, cleanName, ext)

		selfiePath = "/uploads/" + uniqueName
		_ = c.SaveUploadedFile(file, "./uploads/"+uniqueName)
	}
	if file, err := c.FormFile("str_file"); err == nil && file != nil {
		ext := filepath.Ext(file.Filename)
		cleanName := strings.TrimSuffix(file.Filename, ext)
		cleanName = strings.ReplaceAll(cleanName, " ", "_")
		uniqueName := fmt.Sprintf("str_%d_%s%s", nowNano, cleanName, ext)

		strPath = "/uploads/" + uniqueName
		_ = c.SaveUploadedFile(file, "./uploads/"+uniqueName)
	}
	if file, err := c.FormFile("sip_file"); err == nil && file != nil {
		ext := filepath.Ext(file.Filename)
		cleanName := strings.TrimSuffix(file.Filename, ext)
		cleanName = strings.ReplaceAll(cleanName, " ", "_")
		uniqueName := fmt.Sprintf("sip_%d_%s%s", nowNano, cleanName, ext)

		sipPath = "/uploads/" + uniqueName
		_ = c.SaveUploadedFile(file, "./uploads/"+uniqueName)
	}

	regRes, err := h.service.RegisterPPDS(c.Request.Context(), req, selfiePath, strPath, sipPath)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Pendaftaran PPDS berhasil dikirim. Menunggu verifikasi admin.", regRes)
}

func (h *Handler) GetRegistrations(c *gin.Context) {
	status := c.Query("status")
	regs, err := h.service.GetRegistrations(c.Request.Context(), status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil daftar permintaan registrasi", regs)
}

func (h *Handler) ApproveRegistration(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID registrasi tidak valid")
		return
	}

	userRes, err := h.service.ApproveRegistration(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Permintaan registrasi berhasil disetujui. Akun residen telah dibuat.", userRes)
}

func (h *Handler) RejectRegistration(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID registrasi tidak valid")
		return
	}

	var req RejectRegistrationRequest
	_ = c.ShouldBind(&req)

	if err := h.service.RejectRegistration(c.Request.Context(), id, req.Reason); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Permintaan registrasi telah ditolak dan notifikasi email dikirim.", nil)
}
