package mockapi

import (
	"embed"
)

//go:embed static/*
var staticFiles embed.FS

var mockapiPath = "/mockapi"

func GetStaticFiles() embed.FS {
	return staticFiles
}

func GetMockapiPath() string {
	return mockapiPath
}

type DataSource interface {
	PopulateData() error
	GetBookByID(id string) (Book, error)
	GetBooks(page int, pageSize int, search string) ([]Book, error)
	GetBooksCount(search string) (int64, error)
}

type Router interface {
	SetupMockApiRoute(service Service) error
}

func Use(ds DataSource, r Router) {
	service := NewService(ds)

	if err := ds.PopulateData(); err != nil {
		panic(err)
	}
	if err := r.SetupMockApiRoute(service); err != nil {
		panic(err)
	}
}