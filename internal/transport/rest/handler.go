package rest

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/dmytrodemianchuk/crud-music/internal/domain"

	"github.com/gorilla/mux"
)

type Musics interface {
	Create(ctx context.Context, book domain.Music) error
	GetByID(ctx context.Context, id int64) (domain.Music, error)
	GetAll(ctx context.Context) ([]domain.Music, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp domain.UpdateMusicInput) error
}

type User interface {
	SignUp(ctx context.Context, inp domain.SignUpInput) error
	SignIn(ctx context.Context, inp domain.SignInInput) (string, string, error)
	ParseToken(ctx context.Context, accessToken string) (int64, error)
	RefreshTokens(ctx context.Context, refreshToken string) (string, string, error)
}

type Handler struct {
	musicsService Musics
	usersService  User
}

func NewHandler(books Musics, users User) *Handler {
	return &Handler{
		musicsService: books,
		usersService:  users,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
		auth.HandleFunc("/sign-in", h.signIn).Methods(http.MethodGet)
		auth.HandleFunc("/refresh", h.refresh).Methods(http.MethodGet)
	}

	musics := r.PathPrefix("/books").Subrouter()
	{
		musics.Use(h.authMiddleware)

		musics.HandleFunc("", h.createMusic).Methods(http.MethodPost)
		musics.HandleFunc("", h.getAllMusics).Methods(http.MethodGet)
		musics.HandleFunc("/{id:[0-9]+}", h.getMusicByID).Methods(http.MethodGet)
		musics.HandleFunc("/{id:[0-9]+}", h.deleteMusic).Methods(http.MethodDelete)
		musics.HandleFunc("/{id:[0-9]+}", h.updateMusic).Methods(http.MethodPut)
	}

	return r
}

func getIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return id, nil
}
