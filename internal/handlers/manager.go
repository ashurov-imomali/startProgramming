package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"main/pkg/models"
	"net/http"
	"strconv"
)

func (h *Handler) AddTable(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 3 {
		h.Logger.Error(errors.New("not manager"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newTable models.Table
	err := json.NewDecoder(r.Body).Decode(&newTable)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.CheckNewTable(&newTable)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.CreateTable(&newTable)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AddZone(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 3 {
		h.Logger.Error(errors.New("not manager"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newZone models.Zone
	err := json.NewDecoder(r.Body).Decode(&newZone)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.CheckNewZone(&newZone)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.AddNewZone(&newZone)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) AddCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 3 {
		h.Logger.Error(errors.New("not manager"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newCategory models.Categories
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.CheckNewCategory(&newCategory)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.CreateNewCategory(&newCategory)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) AddMenuItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 3 {
		h.Logger.Error(errors.New("not manager"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var menuItem models.Menu
	err := json.NewDecoder(r.Body).Decode(&menuItem)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.CheckNewMenuItem(&menuItem)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.CreateNewMenuItem(&menuItem)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) ChangeMenuItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 3 {
		h.Logger.Error(errors.New("not manager"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var menuItem models.Menu
	err := json.NewDecoder(r.Body).Decode(&menuItem)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateMenuItem(&menuItem)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetSalary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 3 {
		h.Logger.Error(errors.New("not manager"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	strUserId := vars["id"]
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	employer, err := h.Service.GetEmployer(userId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.MarshalIndent(employer, "", " ")
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (h *Handler) AddAllowance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 3 {
		h.Logger.Error(errors.New("not manager"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	strUserId := vars["id"]
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	strAllowance := r.URL.Query().Get("allowance")
	allowance, err := strconv.ParseFloat(strAllowance, 64)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.GiveAllowance(userId, allowance)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AddFines(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 3 {
		h.Logger.Error(errors.New("not manager"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	strUserId := vars["id"]
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	strFines := r.URL.Query().Get("fines")
	fines, err := strconv.ParseFloat(strFines, 64)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.GiveFines(userId, fines)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
