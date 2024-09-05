package user

type (
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)
