package http

import (
	"errors"
	nethttp "net/http"
	"strconv"

	"vananin/go-from-scratch/internal/domain"
	"vananin/go-from-scratch/internal/service"
)

type UserHandler struct {
	userSvc   *service.UserService
	responder *Responder
}

func NewUserHandler(userSvc *service.UserService) *UserHandler {
	return &UserHandler{
		userSvc:   userSvc,
		responder: NewResponder(),
	}
}

func (h *UserHandler) GetByID(w nethttp.ResponseWriter, r *nethttp.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		h.responder.Error(w, nethttp.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.userSvc.GetByID(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserNotFound):
			h.responder.Error(w, nethttp.StatusNotFound, "user not found")
		default:
			h.responder.Error(w, nethttp.StatusInternalServerError, "internal server error")
		}
		return
	}

	h.responder.JSON(w, nethttp.StatusOK, user)
}
