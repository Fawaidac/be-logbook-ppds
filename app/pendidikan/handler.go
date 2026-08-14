package pendidikan

import (
	"net/http"
	"strconv"

	"be-logbook-ppds/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	kompetensiSvc KompetensiService
	rotasiSvc     RotasiService
	miniCexSvc    MiniCexService
	dopsSvc       DopsService
	seminarSvc    SeminarService
	cbdSvc        CbdService
}

func NewHandler(
	kompetensiSvc KompetensiService,
	rotasiSvc RotasiService,
	miniCexSvc MiniCexService,
	dopsSvc DopsService,
	seminarSvc SeminarService,
	cbdSvc CbdService,
) *Handler {
	return &Handler{
		kompetensiSvc: kompetensiSvc,
		rotasiSvc:     rotasiSvc,
		miniCexSvc:    miniCexSvc,
		dopsSvc:       dopsSvc,
		seminarSvc:    seminarSvc,
		cbdSvc:        cbdSvc,
	}
}

// Kompetensi Handlers
func (h *Handler) GetKompetensi(c *gin.Context) {
	entries, err := h.kompetensiSvc.GetAllKompetensi(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Kompetensi entries retrieved", entries)
}

func (h *Handler) CreateKompetensi(c *gin.Context) {
	var req CreateKompetensiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	username := ""
	if val, exists := c.Get("username"); exists {
		if u, ok := val.(string); ok {
			username = u
		}
	}

	entry, err := h.kompetensiSvc.CreateKompetensi(c.Request.Context(), req, username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Kompetensi berhasil dibuat", entry)
}

func (h *Handler) DeleteKompetensi(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.kompetensiSvc.DeleteKompetensi(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Kompetensi berhasil dihapus", nil)
}

// Rotasi Handlers
func (h *Handler) GetRotasi(c *gin.Context) {
	entries, err := h.rotasiSvc.GetAllRotasi(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Rotasi entries retrieved", entries)
}

func (h *Handler) CreateRotasi(c *gin.Context) {
	var req CreateRotasiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	username := ""
	if val, exists := c.Get("username"); exists {
		if u, ok := val.(string); ok {
			username = u
		}
	}

	entry, err := h.rotasiSvc.CreateRotasi(c.Request.Context(), req, username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Rotasi berhasil dibuat", entry)
}

func (h *Handler) DeleteRotasi(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.rotasiSvc.DeleteRotasi(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Rotasi berhasil dihapus", nil)
}

// MiniCex Handlers
func (h *Handler) GetMiniCex(c *gin.Context) {
	entries, err := h.miniCexSvc.GetAllMiniCex(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Mini-CEX entries retrieved", entries)
}

func (h *Handler) CreateMiniCex(c *gin.Context) {
	var req CreateMiniCexRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	username := ""
	if val, exists := c.Get("username"); exists {
		if u, ok := val.(string); ok {
			username = u
		}
	}

	entry, err := h.miniCexSvc.CreateMiniCex(c.Request.Context(), req, username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Mini-CEX berhasil dibuat", entry)
}

func (h *Handler) DeleteMiniCex(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.miniCexSvc.DeleteMiniCex(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Mini-CEX berhasil dihapus", nil)
}

// DOPS Handlers
func (h *Handler) GetDops(c *gin.Context) {
	entries, err := h.dopsSvc.GetAllDops(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "DOPS entries retrieved", entries)
}

func (h *Handler) CreateDops(c *gin.Context) {
	var req CreateDopsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	username := ""
	if val, exists := c.Get("username"); exists {
		if u, ok := val.(string); ok {
			username = u
		}
	}

	entry, err := h.dopsSvc.CreateDops(c.Request.Context(), req, username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "DOPS berhasil dibuat", entry)
}

func (h *Handler) DeleteDops(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.dopsSvc.DeleteDops(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "DOPS berhasil dihapus", nil)
}

// Seminar Handlers
func (h *Handler) GetSeminar(c *gin.Context) {
	entries, err := h.seminarSvc.GetAllSeminar(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Seminar entries retrieved", entries)
}

func (h *Handler) CreateSeminar(c *gin.Context) {
	var req CreateSeminarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	username := ""
	if val, exists := c.Get("username"); exists {
		if u, ok := val.(string); ok {
			username = u
		}
	}

	entry, err := h.seminarSvc.CreateSeminar(c.Request.Context(), req, username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Seminar berhasil dibuat", entry)
}

func (h *Handler) DeleteSeminar(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.seminarSvc.DeleteSeminar(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Seminar berhasil dihapus", nil)
}

// CBD Handlers
func (h *Handler) GetCbd(c *gin.Context) {
	entries, err := h.cbdSvc.GetAllCbd(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "CBD entries retrieved", entries)
}

func (h *Handler) CreateCbd(c *gin.Context) {
	var req CreateCbdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	username := ""
	if val, exists := c.Get("username"); exists {
		if u, ok := val.(string); ok {
			username = u
		}
	}

	entry, err := h.cbdSvc.CreateCbd(c.Request.Context(), req, username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "CBD berhasil dibuat", entry)
}

func (h *Handler) DeleteCbd(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.cbdSvc.DeleteCbd(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "CBD berhasil dihapus", nil)
}
