package container

import (
	"github.com/sarulabs/di"

	"github.com/wojciechpawlinow/find-indexes/internal/application/service"
	"github.com/wojciechpawlinow/find-indexes/internal/domain/index"
	"github.com/wojciechpawlinow/find-indexes/internal/infrastructure/database/memory"
	"github.com/wojciechpawlinow/find-indexes/internal/infrastructure/httpserver/handlers"
	"github.com/wojciechpawlinow/find-indexes/pkg/logger"
)

func New() di.Container {
	builder, _ := di.NewBuilder()

	if err := builder.Add(di.Def{
		Name: "repo-index",
		Build: func(ctn di.Container) (interface{}, error) {
			return &memory.SliceRepository{}, nil
		},
	}); err != nil {
		logger.Error(err)
	}

	if err := builder.Add(di.Def{
		Name: "service-index",
		Build: func(ctn di.Container) (interface{}, error) {
			return &service.IndexService{
				Repo: ctn.Get("repo-index").(index.Repository),
			}, nil
		},
	}); err != nil {
		logger.Error(err)
	}

	if err := builder.Add(di.Def{
		Name: "http-index",
		Build: func(ctn di.Container) (interface{}, error) {
			return &handlers.IndexHTTPHandler{
				Srv: ctn.Get("service-index").(service.IndexPort),
			}, nil
		},
	}); err != nil {
		logger.Error(err)

	}

	return builder.Build()
}
