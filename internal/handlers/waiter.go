package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"main/pkg/models"
	"net/http"
	"os"
	"strconv"
)

func (h *Handler) GetFreeTables(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId")
	if roleId != 1 {
		h.Logger.Error("Not Waiter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	query := r.URL.Query()

	zone := query.Get("zone")
	strPage := query.Get("page")
	strLimit := query.Get("limit")

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

	tables, err := h.Service.FindTables(zone, page, limit)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.MarshalIndent(tables, "", " ")
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}

func (h *Handler) ReservedTable(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId")
	if roleId != 1 {
		h.Logger.Error("Not Waiter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	strId := vars["id"]

	tableId, err := strconv.Atoi(strId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.BookingTheTable(tableId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UnreservedTable(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	roleId := ctx.Value("roleId")
	if roleId != 1 {
		h.Logger.Error("Not Waiter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	strTableId := vars["id"]
	tableId, err := strconv.Atoi(strTableId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.CancelReservation(tableId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	roleId := ctx.Value("roleId")
	if roleId != 1 {
		h.Logger.Error("Not Waiter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := r.URL.Query()

	categories := query.Get("categories")
	strPage := query.Get("page")
	strLimit := query.Get("limit")

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

	menuCategories, err := h.Service.ReturnCategories(categories, page, limit)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.MarshalIndent(menuCategories, "", " ")
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (h *Handler) GetMenu(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	roleId := ctx.Value("roleId")
	if roleId != 1 {
		h.Logger.Error("Not Waiter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := r.URL.Query()

	dishesName := query.Get("dishes")
	strPage := query.Get("page")
	strLimit := query.Get("limit")

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

	menu, err := h.Service.ReturnMenu(dishesName, page, limit)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.MarshalIndent(menu, "", " ")
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (h *Handler) TakeOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId")
	if roleId != 1 {
		h.Logger.Error("Not Waiter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	orderItems := make([]models.OrderItem, 0)

	err := json.NewDecoder(r.Body).Decode(&orderItems)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	strOrderType := r.URL.Query().Get("order_type")
	orderType, err := strconv.ParseBool(strOrderType)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	strTableId := vars["id"]
	tableId, err := strconv.Atoi(strTableId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId := r.Context().Value("id").(int)
	err = h.Service.CheckTable(tableId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var order models.Order
	order.OrderType = orderType
	order.WaiterId = userId
	order.TableId = tableId

	orderId, err := h.Service.AddOrder(&order, orderItems)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	order.Id = orderId
	err = h.Service.AddOrderItems(orderId, orderItems)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.MarshalIndent(order, "", " ")
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(bytes)
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId")
	if roleId != 1 {
		h.Logger.Error("Not Waiter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)

	strId := vars["id"]
	query := r.URL.Query()
	strPage := query.Get("page")
	strLimit := query.Get("limit")

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

	id, err := strconv.Atoi(strId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	order, err := h.Service.ReturnOrder(id, page, limit)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.MarshalIndent(order, "", " ")
	if err != nil {
		return
	}
	w.Write(bytes)
}

func (h *Handler) GetCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roleId := ctx.Value("roleId")
	if roleId != 1 {
		h.Logger.Error("Not Waiter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.CreateCheck(id)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	file, err := os.Open("output.pdf")
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	_, err = io.Copy(w, file)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Pay(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("id").(int)
	roleId := ctx.Value("roleId").(int)
	if roleId != 1 {
		h.Logger.Error("Not Waiter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, checkUserId := h.Service.CheckOrder(id)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if checkUserId != userId {
		h.Logger.Error(errors.New("incorrect user"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.CloseCheck(userId, id)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
