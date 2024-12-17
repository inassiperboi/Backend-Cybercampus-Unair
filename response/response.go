package response


type Response struct {
	Status  int    `json:"STATUS"`
	Message string `json:"MESSAGE"`
	Data    interface{} `json:"DATA"`
}

type ValidateErrorResponse struct {
	Error bool `json:"ERROR"`
	FailedField string `json:"FAILED_FIELD"`
	Tag string `json:"TAG"`
	Value interface{} `json:"VALUE"`
}