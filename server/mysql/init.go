package database

import (
	"fmt"

	accountsDB "github.com/nmarcetic/gorchase/server/accounts/store/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config define Postgres params
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Init initialize DB connection
func Init(cfg Config) (*gorm.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		//TODO Log error
		return &gorm.DB{}, err
	}
	// Auto migration
	if err := db.AutoMigrate(
		&accountsDB.UserDB{},
	); err != nil {
		fmt.Println("faild to migrate db")
		return &gorm.DB{}, err
	}
	return db, nil

}
