package auth

type(
	//register
	RegisterRequest struct{
		Name string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	RegisterResponse struct{
		Token string `json:"token"`
	}

	//login
	LoginRequest struct{
		Email string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	LoginResponse struct{
		Token string `json:"token"`
	}
)