package container

import (
	"github.com/sarulabs/di"

	"github.com/wojciechpawlinow/find-indexes/internal/application/service"
	"github.com/wojciechpawlinow/find-indexes/internal/domain/index"
	"github.com/wojciechpawlinow/find-indexes/internal/infrastructure/database/memory"
	"github.com/wojciechpawlinow/find-indexes/internal/infrastructure/httpserver/handlers"
)

func New() di.Container {
	builder, _ := di.NewBuilder()

	builder.Add(di.Def{
		Name: "http-index",
		Build: func(ctn di.Container) (interface{}, error) {
			return &handlers.IndexFinderHTTPHandler{}, nil
		},
	})

	builder.Add(di.Def{
		Name: "repo-index",
		Build: func(ctn di.Container) (interface{}, error) {
			return &memory.SliceRepository{}, nil
		},
	})

	builder.Add(di.Def{
		Name: "service-index",
		Build: func(ctn di.Container) (interface{}, error) {
			return &service.IndexFinderService{
				Repo: ctn.Get("repo-index").(index.Repository),
			}, nil
		},
	})

	return builder.Build()
}
