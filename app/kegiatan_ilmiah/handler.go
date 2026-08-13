package kegiatan_ilmiah

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

// Kegiatan Ilmiah Handlers
func (h *Handler) GetIndex(c *gin.Context) {
	userID := 0
	if val, exists := c.Get("user_id"); exists {
		if id, ok := val.(int); ok {
			userID = id
		}
	}

	entries, err := h.service.GetAllKegiatan(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Group by kategori
	kategorized := map[string][]KegiatanIlmiahResponse{
		"simposium":     {},
		"workshop":      {},
		"multidisiplin": {},
		"ilmiah_lain":   {},
	}

	for _, entry := range entries {
		kategorized[entry.Kategori] = append(kategorized[entry.Kategori], entry)
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil kegiatan ilmiah", kategorized)
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateKegiatanIlmiahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	// Extract user session data
	username := ""
	programStudi := ""
	ppdsName := ""
	nimNip := ""
	
	if val, exists := c.Get("username"); exists {
		if u, ok := val.(string); ok {
			username = u
		}
	}
	if val, exists := c.Get("program_studi"); exists {
		if ps, ok := val.(string); ok {
			programStudi = ps
		}
	}
	if val, exists := c.Get("name"); exists {
		if pn, ok := val.(string); ok {
			ppdsName = pn
		}
	}
	if val, exists := c.Get("nim_nip"); exists {
		if nn, ok := val.(string); ok {
			nimNip = nn
		}
	}

	res, err := h.service.CreateKegiatan(c.Request.Context(), req, username, programStudi, ppdsName, nimNip)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"message":  "Kegiatan ilmiah berhasil disimpan.",
		"kegiatan": res,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.service.DeleteKegiatan(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Kegiatan ilmiah berhasil dihapus.",
	})
}

// Bimbingan Penelitian Handlers
func (h *Handler) GetBimbinganIndex(c *gin.Context) {
	entries, err := h.service.GetAllBimbingan(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil bimbingan penelitian", entries)
}

func (h *Handler) CreateBimbingan(c *gin.Context) {
	var req CreateBimbinganRequest
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

	res, err := h.service.CreateBimbingan(c.Request.Context(), req, username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":    true,
		"message":    "Catatan bimbingan penelitian berhasil disimpan.",
		"bimbingan": res,
	})
}
