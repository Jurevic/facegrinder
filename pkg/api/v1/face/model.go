package face

import (
	"github.com/go-pg/pg/orm"
)

type Face struct {
	Id       string       `json:"id"`
	OwnerId  string       `json:"-"`
	Image    string       `json:"image"`
	Name     string       `json:"name"`
	Encoding [128]float64 `json:"-"`
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
