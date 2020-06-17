package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mreider/koto/backend/central/repo"
	"github.com/mreider/koto/backend/central/service"
	"github.com/mreider/koto/backend/common"
)

func Friend(friendRepo repo.FriendRepo) http.Handler {
	h := &friendHandlers{
		friendRepo: friendRepo,
	}
	r := chi.NewRouter()
	r.Post("/", h.Friends)
	return r
}

type friendHandlers struct {
	friendRepo repo.FriendRepo
}

func (h *friendHandlers) Friends(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(service.ContextUserKey).(repo.User)
	friends, err := h.friendRepo.Friends(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if friends == nil {
		friends = []repo.User{}
	}
	common.WriteJSONToResponse(w, friends)
}