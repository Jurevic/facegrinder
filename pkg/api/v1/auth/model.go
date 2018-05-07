package auth

import (
	_ "github.com/go-ozzo/ozzo-validation"
	"github.com/go-pg/pg/orm"
	"time"
)

type User struct {
	Id          int       `json:"id"`
	Password    []byte    `json:"-"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`

	IsSuperuser bool      `json:"is_superuser"`
	IsActive    bool      `json:"is_active"`

	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	LastLogin   time.Time `json:"last_login,omitempty"`
}

// BeforeInsert hook executed before database insert operation.
func (u *User) BeforeInsert(db orm.DB) error {
	u.CreatedAt = time.Now()
	u.Password = hashAndSalt(u.Password)
	return nil
}

// BeforeUpdate hook executed before database update operation.
func (u *User) BeforeUpdate(db orm.DB) error {
	u.UpdatedAt = time.Now()
	return u.Validate()
}

// Validate validates User struct and returns validation errors.
func (u *User) Validate() error {
	return nil
	//return validation.ValidateStruct(u,
	//	validation.Field(&u.Theme, validation.Required, validation.In("default", "dark")),
	//)
}
