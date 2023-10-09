package router

import (
	"github.com/gorilla/mux"
	"main/internal/handlers"
	"net/http"
)

func GetRouter(h *handlers.Handler) *mux.Router {
	newRouter := mux.NewRouter()
	newRouter.HandleFunc("/roles", h.GetRoles).Methods(http.MethodGet)
	newRouter.HandleFunc("/registration", h.Registration).Methods(http.MethodPost)
	newRouter.HandleFunc("/login", h.LogIn).Methods(http.MethodGet)

	employersRouter := newRouter.PathPrefix("").Subrouter()
	employersRouter.Use(h.CheckUser)

	tableRouter := employersRouter.PathPrefix("/table").Subrouter()
	tableRouter.HandleFunc("", h.GetFreeTables).Methods(http.MethodGet)
	tableRouter.HandleFunc("/reserved/{id}", h.ReservedTable).Methods(http.MethodPatch)
	tableRouter.HandleFunc("/unreserved/{id}", h.UnreservedTable).Methods(http.MethodPatch)
	tableRouter.HandleFunc("", h.AddTable).Methods(http.MethodPost)
	tableRouter.HandleFunc("/zone", h.AddZone).Methods(http.MethodPost)

	menuRouter := employersRouter.PathPrefix("/menu").Subrouter()
	menuRouter.HandleFunc("/category", h.GetCategories).Methods(http.MethodGet)
	menuRouter.HandleFunc("", h.GetMenu).Methods(http.MethodGet)
	menuRouter.HandleFunc("/category", h.AddCategories).Methods(http.MethodPost)
	menuRouter.HandleFunc("", h.AddMenuItem).Methods(http.MethodPost)
	menuRouter.HandleFunc("", h.ChangeMenuItem).Methods(http.MethodPatch)

	orderRouter := employersRouter.PathPrefix("/order").Subrouter()
	orderRouter.HandleFunc("/{id}", h.TakeOrder).Methods(http.MethodPost)
	orderRouter.HandleFunc("/chef/unfinished", h.GetUnfinishedOrders).Methods(http.MethodGet)
	orderRouter.HandleFunc("/chef/gettingReady/{id}", h.GettingReady).Methods(http.MethodPatch)
	orderRouter.HandleFunc("/chef/getReady/{id}", h.GetReady).Methods(http.MethodPatch)
	orderRouter.HandleFunc("/{id}", h.GetOrder).Methods(http.MethodGet)
	orderRouter.HandleFunc("/check/{id}", h.GetCheck).Methods(http.MethodGet)
	orderRouter.HandleFunc("/check/{id}", h.Pay).Methods(http.MethodPost)

	finesRouter := employersRouter.PathPrefix("/compensation").Subrouter()
	finesRouter.HandleFunc("/{id}", h.GetSalary).Methods(http.MethodGet)
	finesRouter.HandleFunc("/allowance/{id}", h.AddAllowance).Methods(http.MethodPost)
	finesRouter.HandleFunc("/fines/{id}", h.AddFines).Methods(http.MethodPost)

	return newRouter
}
