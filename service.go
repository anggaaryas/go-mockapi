package mockapi

type service struct {
	repository BookRepository
}

type Service interface {
	GetBookByID(id string) (Book, error)
	GetBooks(page int, pageSize int) ([]Book, error)
}

func NewService(repository BookRepository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetBookByID(id string) (Book, error) {
	return s.repository.GetBookByID(id)
}

func (s *service) GetBooks(page int, pageSize int) ([]Book, error) {
	return s.repository.GetBooks(page, pageSize)
}