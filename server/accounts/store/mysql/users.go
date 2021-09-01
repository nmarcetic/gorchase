package postgres

import (
	"encoding/json"
	"fmt"

	"github.com/nmarcetic/gorchase/server/accounts"
	"gorm.io/gorm"
)

// UserDB  defines user structure as Gorm model representation
type UserDB struct {
	gorm.Model
	ID       string `gorm:"primaryKey;uniqueIndex;not null;<-:create"`
	Email    string `gorm:"size:255;uniqueIndex;not null"`
	Password []byte `gorm:"size:255;unique;not null" json:"-"`
	Metadata []byte
}

// TableName tells gorm to set UserDB table name
func (UserDB) TableName() string {
	return "users"
}

type userRepository struct {
	db *gorm.DB
}

func toDBObject(u accounts.User) (UserDB, error) {
	data := []byte("{}")
	if len(u.Metadata) > 0 {
		b, err := json.Marshal(u.Metadata)
		if err != nil {
			return UserDB{}, fmt.Errorf("Err parsing json")
		}
		data = b
	}
	return UserDB{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
		Metadata: data,
	}, nil
}

func fromDBObject(udb UserDB, secrets bool) (accounts.User, error) {
	var metadata map[string]interface{}
	if udb.Metadata != nil {
		if err := json.Unmarshal([]byte(udb.Metadata), &metadata); err != nil {
			return accounts.User{}, fmt.Errorf("Error pasing json")
		}
	}
	// Exclude secrets in response if not explicitly request
	var pwd []byte
	if secrets {
		pwd = udb.Password
	}

	return accounts.User{
		ID:        udb.ID,
		Email:     udb.Email,
		Password:  pwd,
		Metadata:  metadata,
		CreatedAt: udb.CreatedAt,
		UpadtedAt: udb.UpdatedAt,
	}, nil
}

//NewUserRepository Returns new instance of userRepository
func NewUserRepository(db *gorm.DB) accounts.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo userRepository) Create(u accounts.User) (id string, err error) {
	dbu, err := toDBObject(u)
	if err != nil {
		return "", err
	}
	result := repo.db.Create(&dbu)
	if result.Error != nil {
		return "", result.Error
	}
	return dbu.ID, nil
}

func (repo userRepository) Get(id string) (u accounts.User, err error) {
	var user UserDB
	result := repo.db.Where("id", id).First(&user)
	if result.Error != nil {
		return accounts.User{}, result.Error
	}
	return fromDBObject(user, false)
}

func (repo userRepository) GetByEmail(email string) (accounts.User, error) {
	var user UserDB
	result := repo.db.Where("email", email).First(&user)
	if result.Error != nil {
		return accounts.User{}, result.Error
	}
	return fromDBObject(user, true)
}

func (repo userRepository) Update(u accounts.User) error {
	user, err := toDBObject(u)
	if err != nil {
		return err
	}
	result := repo.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) Delete(id string) error {
	return nil
}
