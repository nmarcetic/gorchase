package accounts

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/idna"
)

const (
	minPassLen   = 8
	maxLocalLen  = 64
	maxDomainLen = 255
	maxTLDLen    = 24 // longest TLD currently in existence

	atSeparator       = "@"
	dotSeparator      = "."
	errInvalidEmail   = "Invalid Email format"
	errPasswordLenght = "Password lenght invalid"
)

var (
	userRegexp    = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
	hostRegexp    = regexp.MustCompile("^[^\\s]+\\.[^\\s]+$")
	userDotRegexp = regexp.MustCompile("(^[.]{1})|([.]{1}$)|([.]{2,})")
)

// Metadata to be used for storing user  arrbitary data
type Metadata map[string]interface{}

// User defines user account structures
type User struct {
	ID        string
	Email     string
	Password  []byte
	Metadata  Metadata
	CreatedAt time.Time
	UpadtedAt time.Time
}

// UserRepository specifies users persistent storage API
type UserRepository interface {
	Create(u User) (string, error)
	Get(id string) (User, error)
	GetByEmail(email string) (User, error)
	Update(u User) error
	Delete(id string) error
}

// Validate Do User object validation
func (u User) Validate() error {

	if !isEmail(u.Email) {
		return fmt.Errorf(errInvalidEmail)
	}

	if len(u.Password) < minPassLen {
		return fmt.Errorf(errPasswordLenght)
	}

	return nil

}

//TODO: Refactor this
func isEmail(email string) bool {
	if email == "" {
		return false
	}

	es := strings.Split(email, atSeparator)
	if len(es) != 2 {
		return false
	}
	local, host := es[0], es[1]

	if local == "" || len(local) > maxLocalLen {
		return false
	}

	hs := strings.Split(host, dotSeparator)
	if len(hs) < 2 {
		return false
	}
	domain, ext := hs[0], hs[1]

	// Check subdomain and validate
	if len(hs) > 2 {
		if domain == "" {
			return false
		}

		for i := 1; i < len(hs)-1; i++ {
			sub := hs[i]
			if sub == "" {
				return false
			}
			domain = fmt.Sprintf("%s.%s", domain, sub)
		}

		ext = hs[len(hs)-1]
	}

	if domain == "" || len(domain) > maxDomainLen {
		return false
	}
	if ext == "" || len(ext) > maxTLDLen {
		return false
	}

	punyLocal, err := idna.ToASCII(local)
	if err != nil {
		return false
	}
	punyHost, err := idna.ToASCII(host)
	if err != nil {
		return false
	}

	if userDotRegexp.MatchString(punyLocal) || !userRegexp.MatchString(punyLocal) || !hostRegexp.MatchString(punyHost) {
		return false
	}

	return true

}
