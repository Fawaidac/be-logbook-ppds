package auth

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginData struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Jabatan  string `json:"jabatan"`
}

type LoginResponse struct {
	User  LoginData `json:"user"`
	Token string    `json:"token"`
}

