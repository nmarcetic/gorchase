package accounts

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nmarcetic/gorchase/pkg/logger"
)

type usersService struct {
	users  UserRepository
	hasher Hasher
	logger *logger.Logger
	//TODO: Add transactions service
}

var _ Service = (*usersService)(nil)

// Service specifies users API
type Service interface {
	// Register creates new user account.
	Register(u User) (id string, err error)
	// Login verifies provided Email/Password with record in DB
	Login(email, password string) (u User, err error)
	// Retrieves user info for a given user ID.
	Get(id string) (u User, err error)
	// GetByEmail retrives user for a given email.
	GetByEmail(email string) (u User, err error)
	// Update updates user account
	Update(u User) error
}

// NewUsersService init Service
func NewUsersService(repo UserRepository, logger logger.Logger) Service {
	return &usersService{
		users:  repo,
		hasher: NewHasher(),
		logger: &logger,
	}
}

func (svc usersService) Register(u User) (string, error) {
	if err := u.Validate(); err != nil {
		return "", err
	}
	hash, err := svc.hasher.Hash(string(u.Password))
	if err != nil {
		return "", err
	}
	u.Password = []byte(hash)
	u.ID = uuid.New().String()
	u.Metadata = make(Metadata)
	id, err := svc.users.Create(u)
	if err != nil {
		return "", err
	}
	return id, nil

}

func (svc usersService) Login(email, password string) (u User, err error) {
	if email == "" || password == "" {
		return User{}, fmt.Errorf("Invalid credential provided")
	}

	user, err := svc.users.GetByEmail(email)
	if err != nil {
		return User{}, err
	}
	hashed := string(user.Password)
	if err := svc.hasher.Compare(password, hashed); err != nil {
		return User{}, fmt.Errorf("Invalid credential provided")
	}
	return user, nil
}

func (svc usersService) Get(id string) (User, error) {
	u, err := svc.users.Get(id)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (svc usersService) GetByEmail(email string) (User, error) {
	u, err := svc.users.GetByEmail(email)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (svc usersService) Update(u User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := svc.users.Update(u); err != nil {
		return err
	}
	return nil
}
