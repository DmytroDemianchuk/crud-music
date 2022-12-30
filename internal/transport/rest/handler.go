package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/dmytrodemianchuk/crud-music/internal/domain"

	"github.com/gorilla/mux"
)

type Musics interface {
	Create(ctx context.Context, musics domain.Music) error
	GetByID(ctx context.Context, id int64) (domain.Music, error)
	GetAll(ctx context.Context) ([]domain.Music, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp domain.UpdateMusicInput) error
}

type Handler struct {
	musicsService Musics
}

func NewHandler(musics Musics) *Handler {
	return &Handler{
		musicsService: musics,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	musics := r.PathPrefix("/musics").Subrouter()
	{
		musics.HandleFunc("", h.createMusic).Methods(http.MethodPost)
		musics.HandleFunc("", h.getAllMusics).Methods(http.MethodGet)
		musics.HandleFunc("/{id:[0-9]+}", h.getMusicByID).Methods(http.MethodGet)
		musics.HandleFunc("/{id:[0-9]+}", h.deleteMusic).Methods(http.MethodDelete)
		musics.HandleFunc("/{id:[0-9]+}", h.updateMusic).Methods(http.MethodPut)
	}

	return r
}

func (h *Handler) getMusicByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("getMusicByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	music, err := h.musicsService.GetByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, domain.ErrMusicNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("getMusicByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(music)
	if err != nil {
		log.Println("getMusicByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) createMusic(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "createMusic",
			"problem": "reading request body",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var music domain.Music
	if err = json.Unmarshal(reqBytes, &music); err != nil {
		log.WithFields(log.Fields{
			"handler": "createMusic",
			"problem": "unmarshaling request",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.musicsService.Create(context.TODO(), music)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "createMusic",
			"problem": "service error",
		}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) deleteMusic(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("deleteMusic() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.musicsService.Delete(context.TODO(), id)
	if err != nil {
		log.Println("deleteMusic() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getAllMusics(w http.ResponseWriter, r *http.Request) {
	musics, err := h.musicsService.GetAll(context.TODO())
	if err != nil {
		log.Println("getAllMusics() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(musics)
	if err != nil {
		log.Println("getAllMusics() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) updateMusic(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp domain.UpdateMusicInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.musicsService.Update(context.TODO(), id, inp)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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
