package authorized

import (
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"go.uber.org/zap"
)

type handler struct {
	db     mysql.Repo
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func (h *handler) Add() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("authorized_add", nil)
	}
}

func (h *handler) Demo() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("authorized_demo", nil)
	}
}

func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("authorized_list", nil)
	}
}

func (h *handler) Api() core.HandlerFunc {
	type apiRequest struct {
		Id string `uri:"id"` // 主键ID
	}

	type apiResponse struct {
		HashID string `json:"hash_id"` // hashID
	}

	return func(ctx core.Context) {
		req := new(apiRequest)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		obj := new(apiResponse)
		obj.HashID = req.Id

		ctx.HTML("authorized_api", obj)
	}
}
