package users

type Service interface {
	CreateUser(email string) (*User, error)
	GetUserByID(id uint) (*User, error)
	UpdateUser(id uint, email string) (*User, error)
	DeleteUser(id uint) error
	ListUsers() ([]User, error)
}

type service struct {
	repo UserRepository
}

func NewService(repo UserRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(email string) (*User, error) {
	user := &User{Email: email}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) GetUserByID(id uint) (*User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) UpdateUser(id uint, email string) (*User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	user.Email = email
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) DeleteUser(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *service) ListUsers() ([]User, error) {
	users, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return users, nil
}
