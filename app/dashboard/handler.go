package dashboard

import (
	"net/http"
	"strconv"
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

func (h *Handler) GetDashboardSummary(c *gin.Context) {
	yearStr := c.Query("year")
	year, _ := strconv.Atoi(yearStr)

	username := ""
	if val, exists := c.Get("username"); exists {
		if u, ok := val.(string); ok {
			username = u
		}
	}

	name := ""
	if val, exists := c.Get("name"); exists {
		if n, ok := val.(string); ok {
			name = n
		}
	}

	role := ""
	if val, exists := c.Get("role"); exists {
		if r, ok := val.(string); ok {
			role = r
		}
	}

	// Filter per user jika role residen / jika query param username diset
	filterUsername := username
	filterName := name
	if role == "superadmin" || role == "admin" || role == "supervisor" {
		if reqUser := c.Query("username"); reqUser != "" {
			filterUsername = reqUser
			filterName = ""
		} else if role == "superadmin" || role == "admin" {
			filterUsername = ""
			filterName = ""
		}
	}

	res, err := h.service.GetDashboardSummary(c.Request.Context(), year, filterUsername, filterName)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data dashboard", res)
}

func (h *Handler) GetLaporanSummary(c *gin.Context) {
	periode := c.Query("periode")
	stase := c.Query("stase")

	username := ""
	if val, exists := c.Get("username"); exists {
		if u, ok := val.(string); ok {
			username = u
		}
	}

	name := ""
	if val, exists := c.Get("name"); exists {
		if n, ok := val.(string); ok {
			name = n
		}
	}

	role := ""
	if val, exists := c.Get("role"); exists {
		if r, ok := val.(string); ok {
			role = r
		}
	}

	filterUsername := username
	filterName := name
	if role == "superadmin" || role == "admin" || role == "supervisor" {
		if reqUser := c.Query("username"); reqUser != "" {
			filterUsername = reqUser
			filterName = ""
		} else if role == "superadmin" || role == "admin" {
			filterUsername = ""
			filterName = ""
		}
	}

	res, err := h.service.GetLaporanSummary(c.Request.Context(), periode, stase, filterUsername, filterName)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data laporan", res)
}

// GetResidentDashboard mengembalikan ringkasan dashboard khusus residen.
// Berbeda dengan endpoint /dashboard umum, handler ini SELALU memfilter data
// berdasarkan username user yang login, sehingga aman untuk role residen.
func (h *Handler) GetResidentDashboard(c *gin.Context) {
	yearStr := c.Query("year")
	year, _ := strconv.Atoi(yearStr)
	if year == 0 {
		year = time.Now().Year()
	}

	username := c.GetString("username")
	name := c.GetString("name")

	res, err := h.service.GetDashboardSummary(c.Request.Context(), year, username, name)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data dashboard residen", res)
}

// GetAdminDashboard mengembalikan ringkasan dashboard khusus admin
// (agregasi seluruh pengguna dan aktivitas sistem).
func (h *Handler) GetAdminDashboard(c *gin.Context) {
	res, err := h.service.GetAdminDashboard(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data dashboard admin", res)
}

// GetSupervisorDashboard mengembalikan ringkasan dashboard khusus supervisor:
// statistik validasi logbook, perkembangan 6 bulan terakhir, daftar residen
// paling aktif, dan antrian validasi terbaru.
func (h *Handler) GetSupervisorDashboard(c *gin.Context) {
	// Supervisor hanya melihat data ppds yang dibimbingnya;
	// nama diambil dari klaim JWT (users.name).
	res, err := h.service.GetSupervisorDashboard(c.Request.Context(), c.GetString("name"))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data dashboard supervisor", res)
}
