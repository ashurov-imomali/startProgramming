package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) GetUnfinishedOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 2 {
		h.Logger.Error(errors.New("incorrect user"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	strPage := r.URL.Query().Get("page")
	strLimit := r.URL.Query().Get("limit")
	page, err := strconv.Atoi(strPage)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(strLimit)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	orders, err := h.Service.ReturnUnfinishedOrders(page, limit)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (h *Handler) GettingReady(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 2 {
		h.Logger.Error(errors.New("incorrect user"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	strOrderItemId := vars["id"]
	orderItemId, err := strconv.Atoi(strOrderItemId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.PatchOrderItem(orderItemId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetReady(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId").(int)
	if roleId != 2 {
		h.Logger.Error(errors.New("incorrect user"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId := ctx.Value("id").(int)
	vars := mux.Vars(r)
	strOrderItemId := vars["id"]
	orderItemId, err := strconv.Atoi(strOrderItemId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.FinishOrder(orderItemId, userId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
