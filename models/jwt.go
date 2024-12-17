package models

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}


type PayLoad struct{
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	JenisUser string `json:"jenis_user"`
	Role string `json:"role"`
	Exp int64 `json:"exp"`
}

type JWTClaims struct {
	ID 	 string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	JenisUser string `json:"jenis_user"`
	Role string `json:"role"`
	Exp      int64  `json:"exp"`
}



