package approval

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

// supervisorScope menentukan filter antrian approval berdasarkan role:
// - supervisor  : hanya entri di mana dia dipilih sebagai DPJP/pembimbing
// - role lain   : kosong (melihat semua, mis. admin)
func supervisorScope(c *gin.Context) string {
	if c.GetString("role") == "supervisor" {
		return c.GetString("name")
	}
	return ""
}

func (h *Handler) GetMenunggu(c *gin.Context) {
	list, err := h.service.GetMenunggu(c.Request.Context(), supervisorScope(c))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Daftar persetujuan menunggu", list)
}

func (h *Handler) GetDisetujui(c *gin.Context) {
	list, err := h.service.GetDisetujui(c.Request.Context(), supervisorScope(c))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Daftar persetujuan disetujui", list)
}

func (h *Handler) GetDitolak(c *gin.Context) {
	list, err := h.service.GetDitolak(c.Request.Context(), supervisorScope(c))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Daftar persetujuan ditolak", list)
}

func (h *Handler) ApproveTindakan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.service.ApproveTindakan(c.Request.Context(), id, supervisorScope(c)); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Tindakan berhasil disetujui", nil)
}

func (h *Handler) RejectTindakan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req RejectRequest
	if err := c.ShouldBind(&req); err != nil {
		req = RejectRequest{Catatan: ""}
	}

	if err := h.service.RejectTindakan(c.Request.Context(), id, supervisorScope(c), req.Catatan); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Tindakan ditandai perlu revisi", nil)
}

func (h *Handler) ApproveKegiatanIlmiah(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.service.ApproveKegiatanIlmiah(c.Request.Context(), id, supervisorScope(c)); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Kegiatan ilmiah berhasil disetujui", nil)
}

func (h *Handler) RejectKegiatanIlmiah(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req RejectRequest
	if err := c.ShouldBind(&req); err != nil {
		req = RejectRequest{Catatan: ""}
	}

	if err := h.service.RejectKegiatanIlmiah(c.Request.Context(), id, supervisorScope(c), req.Catatan); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Kegiatan ilmiah ditandai perlu revisi", nil)
}


func (h *Handler) ApprovePendidikanEvaluasi(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.service.ApprovePendidikanEvaluasi(c.Request.Context(), id, supervisorScope(c)); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Pendidikan evaluasi berhasil disetujui", nil)
}

func (h *Handler) RejectPendidikanEvaluasi(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req RejectRequest
	if err := c.ShouldBind(&req); err != nil {
		req = RejectRequest{Catatan: ""}
	}

	if err := h.service.RejectPendidikanEvaluasi(c.Request.Context(), id, supervisorScope(c), req.Catatan); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Pendidikan evaluasi ditandai perlu revisi", nil)
}
