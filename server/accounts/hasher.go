package accounts

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const cost int = 10

var (
	errHashPassword    = errors.New("Generate hash from password failed")
	errComparePassword = errors.New("Compare hash and password failed")
)

// Hasher specifies an API for generating hashes of an arbitrary textual
// content.
type Hasher interface {
	// Hash generates the hashed string from plain-text.
	Hash(string) (string, error)

	// Compare compares plain-text version to the hashed one. An error should
	// indicate failed comparison.
	Compare(string, string) error
}

var _ Hasher = (*bcryptHasher)(nil)

type bcryptHasher struct{}

// NewHasher instantiates a bcrypt-based hasher implementation.
func NewHasher() Hasher {
	return &bcryptHasher{}
}

func (bh *bcryptHasher) Hash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		//TODO: Log err
		return "", errHashPassword
	}

	return string(hash), nil
}

func (bh *bcryptHasher) Compare(plain, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		//TODO: Log err
		return errComparePassword
	}
	return nil
}
