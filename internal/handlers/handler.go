package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"main/internal/service"
	"main/pkg/models"
	"net/http"
)

type Handler struct {
	Service *service.Service
	Logger  *logrus.Logger
}

func GetHandler(srv *service.Service, logger *logrus.Logger) *Handler {
	return &Handler{Service: srv, Logger: logger}
}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	var NewUser models.Employer
	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.CheckNewUser(NewUser.Login, NewUser.Password)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.CreateNewUser(&NewUser) //todo
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var User models.Employer
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.Logger.Error(err)
		return
	}

	err, user := h.Service.CheckUser(User.Login, User.Password)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.Service.GetToken(user) //todo
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err, token = h.Service.AddNewToken(token)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.MarshalIndent(token, "", " ")
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}

func (h *Handler) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.Service.ReturnRoles()
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.MarshalIndent(roles, "", " ")
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)

}
