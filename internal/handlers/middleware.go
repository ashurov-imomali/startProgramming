package handlers

import (
	"context"
	"main/pkg/models"
	"net/http"
)

func (h *Handler) CheckUser(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var newToken models.Token
		token := r.Header.Get("Authorization")
		newToken.StrToken = token
		err, user := h.Service.GetUser(&newToken)
		if err != nil {
			h.Logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "id", user.Id)
		r = r.WithContext(ctx)
		ctx = context.WithValue(r.Context(), "roleId", user.RoleId)
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}
