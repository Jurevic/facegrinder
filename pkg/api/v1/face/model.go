package face

import (
	"github.com/go-pg/pg/orm"
)

type Face struct {
	Id      string `json:"id"`
	OwnerId int    `json:"-"`
	Url     string `json:"url"`
	Name    string `json:"name"`
}

// BeforeInsert hook executed before database insert operation.
func (i *Face) BeforeInsert(db orm.DB) error {
	return i.Validate()
}

// BeforeUpdate hook executed before database update operation.
func (i *Face) BeforeUpdate(db orm.DB) error {
	return i.Validate()
}

// Validate validates User struct and returns validation errors.
func (i *Face) Validate() error {
	return nil
	//return validation.ValidateStruct(u,
	//	validation.Field(&u.Theme, validation.Required, validation.In("default", "dark")),
	//)
}
