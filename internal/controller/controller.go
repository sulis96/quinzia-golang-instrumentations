package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sulis96/quinzia-golang-instrumentations/internal/model"
	"github.com/sulis96/quinzia-golang-instrumentations/internal/service"
)

type (
	controller struct {
		service service.IService
	}

	IController interface {
		Health(w http.ResponseWriter, r *http.Request)
		CreateMember(w http.ResponseWriter, r *http.Request)
		ReadMember(w http.ResponseWriter, r *http.Request)
	}
)

func NewController(s service.IService) IController {
	return &controller{
		service: s,
	}
}

func (c *controller) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{status: OK}"))
}

func (c *controller) CreateMember(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		c.ResponseError(http.StatusMethodNotAllowed, "method not allowed", w)
		return
	}
	var member model.Member
	var respBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		c.ResponseError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	err = c.service.CreateMember(r.Context(), member)
	if err != nil {
		c.ResponseError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	respBody = map[string]interface{}{
		"status":  "success",
		"message": "success create member",
		"data":    member,
	}
	c.ResponseSuccess(http.StatusOK, respBody, w)
}

func (c *controller) ReadMember(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		c.ResponseError(http.StatusMethodNotAllowed, "method not allowed", w)
		return
	}
	member, err := c.service.ReadMember(r.Context())
	if err != nil {
		c.ResponseError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	c.ResponseSuccess(http.StatusOK, member, w)
}
func (c *controller) ResponseError(code int, message string, w http.ResponseWriter) {
	body := model.ErrorResponse{
		Code:  code,
		Error: message,
	}
	bodyByte, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bodyByte)
}

func (c *controller) ResponseSuccess(code int, body interface{}, w http.ResponseWriter) {
	bodyByte, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bodyByte)
}
