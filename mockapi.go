package mockapi

type DataSource interface {
	PopulateData() error
	GetBookByID(id string) (Book, error)
	GetBooks(page int, pageSize int) ([]Book, error)
}

type Router interface {
	SetupMockApiRoute(service Service) error
}

func Use(ds DataSource, r Router) {
	bookRepository := NewBookRepository(ds)
	service := NewService(bookRepository)

	if err := ds.PopulateData(); err != nil {
		panic(err)
	}
	if err := r.SetupMockApiRoute(service); err != nil {
		panic(err)
	}
}