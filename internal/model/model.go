package model

type (
	ErrorResponse struct {
		Code  int         `json:"code"`
		Error interface{} `json:"error"`
	}

	Member struct {
		Name     string `json:"name"`
		Country  string `json:"country"`
		Language string `json:"language"`
	}
)
