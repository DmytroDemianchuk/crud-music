package rest

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	restErrors "github.com/dmytrodemianchuk/crud-music/pkg/rest/errors"

	"github.com/dmytrodemianchuk/crud-music/internal/domain"
	"github.com/gin-gonic/gin"
)

type MusicService interface {
	List(ctx context.Context) (domain.ListMusic, error)
	Get(ctx context.Context, id int) (domain.Music, error)
	Create(ctx context.Context, music domain.Music) (domain.Music, error)
	Update(ctx context.Context, id int, music domain.Music) (domain.Music, error)
	Delete(ctx context.Context, id int) error
}

type Music struct {
	musicService MusicService
}

func NewMusic(musicService MusicService) *Music {
	return &Music{musicService: musicService}
}

func (m Music) List(ctx *gin.Context) {
	music, err := m.musicService.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusOK, music)
}

func (m Music) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	music, err := m.musicService.Get(ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			ctx.JSON(http.StatusNotFound, restErrors.NewNotFoundErr("music not found"))
		default:
			ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		}

		return
	}

	ctx.JSON(http.StatusOK, music)
}

func (m Music) Create(ctx *gin.Context) {
	var music domain.Music
	if err := ctx.BindJSON(&music); err != nil {
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("cannot parse body", nil))
		return
	}

	createdMusic, err := m.musicService.Create(ctx, music)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusCreated, createdMusic)
}

func (m Music) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	var music domain.Music
	if err := ctx.BindJSON(&music); err != nil {
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("cannot parse body", nil))
		return
	}

	updatedMusic, err := m.musicService.Update(ctx, id, music)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusOK, updatedMusic)
}

func (m Music) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	if err := m.musicService.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
