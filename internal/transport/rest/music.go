package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/dmytrodemianchuk/crud-music/internal/domain"
)

func (h *Handler) getMusicByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("getMusicByID", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	music, err := h.musicsService.GetByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, domain.ErrMusicNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logError("getMusicByID", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(music)
	if err != nil {
		logError("getMusicByID", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) createMusic(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logError("createMusic", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var music domain.Music
	if err = json.Unmarshal(reqBytes, &music); err != nil {
		logError("createBook", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.musicsService.Create(r.Context(), music)
	if err != nil {
		logError("createBook", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) deleteMusic(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("deleteMusic", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.musicsService.Delete(r.Context(), id)
	if err != nil {
		logError("deleteMusic", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getAllMusics(w http.ResponseWriter, r *http.Request) {
	musics, err := h.musicsService.GetAll(r.Context())
	if err != nil {
		logError("getAllBooks", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(musics)
	if err != nil {
		logError("getAllMusics", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) updateMusic(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("updateMusic", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logError("updateBook", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp domain.UpdateMusicInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		logError("updateBook", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.musicsService.Update(r.Context(), id, inp)
	if err != nil {
		logError("updateMusic", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
