package book

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllBooks() ([]Book, error) {
	return s.repo.GetAllBooks()
}

func (s *Service) GetBookByID(id int) (*Book, error) {
	return s.repo.GetBookByID(id)
}

func (s *Service) CreateBook(book *Book) error {
	return s.repo.CreateBook(book)
}

func (s *Service) UpdateBook(id int, book *Book) error {
	return s.repo.UpdateBook(id, book)
}

func (s *Service) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}
