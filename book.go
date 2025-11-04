package mockapi

type BookRepository interface {
	GetBookByID(id string) (Book, error)
	GetBooks(page int, pageSize int) ([]Book, error)
}

type bookRepository struct {
	dataSource DataSource
}

func NewBookRepository(dataSource DataSource) BookRepository {
	return &bookRepository{
		dataSource: dataSource,
	}
}

func (r *bookRepository) GetBookByID(id string) (Book, error) {
	return r.dataSource.GetBookByID(id)
}

func (r *bookRepository) GetBooks(page int, pageSize int) ([]Book, error) {
	return r.dataSource.GetBooks(page, pageSize)
}
