package jadwal

type CreateJadwalRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Start       string `json:"start" binding:"required"`
	End         string `json:"end" binding:"required"`
	AllDay      bool   `json:"all_day"`
	Type        string `json:"type" binding:"required"`
}

type UpdateJadwalRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Start       string `json:"start" binding:"required"`
	End         string `json:"end" binding:"required"`
	AllDay      bool   `json:"all_day"`
	Type        string `json:"type" binding:"required"`
}

type UpdateDatesRequest struct {
	Start  string `json:"start" binding:"required"`
	End    string `json:"end" binding:"required"`
	AllDay bool   `json:"all_day"`
}

type EventResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Start       string `json:"start"`
	End         string `json:"end"`
	AllDay      bool   `json:"allDay"`
	ClassName   string `json:"className"`
	Type        string `json:"type"`
	UserID      int    `json:"userId,omitempty"`
}
