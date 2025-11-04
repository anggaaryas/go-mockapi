package mockapi

type service struct {
	dataSource DataSource
}

type Service interface {
	GetBookByID(id string) (Book, error)
	GetBooks(page int, pageSize int, search string) (PaginatedBooks, error)
}

func NewService(dataSource DataSource) Service {
	return &service{
		dataSource: dataSource,
	}
}

func (s *service) GetBookByID(id string) (Book, error) {
	return s.dataSource.GetBookByID(id)
}

func (s *service) GetBooks(page int, pageSize int, search string) (PaginatedBooks, error) {
	books, err := s.dataSource.GetBooks(page, pageSize, search)
	totalCount, err := s.dataSource.GetBooksCount(search)
	if err != nil {
		return PaginatedBooks{}, err
	}
	return PaginatedBooks{
		Data:       books,
		TotalItems: totalCount,
		TotalPages: int((totalCount + int64(pageSize) - 1) / int64(pageSize)),
		Page:       page,
		PageSize:   pageSize,
	}, nil
}