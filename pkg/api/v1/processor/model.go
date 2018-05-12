package processor

import "github.com/go-pg/pg/orm"

type ProcessingNode struct {
	Type   string                 `json:"type"`
	Params map[string]interface{} `json:"params"`
}

type Processor struct {
	Id         string           `json:"id"`
	OwnerId    int              `json:"-"`
	Name       string           `json:"name"`
	Processors []ProcessingNode `json:"processors"`
}

// BeforeInsert hook executed before database insert operation.
func (i *Processor) BeforeInsert(db orm.DB) error {
	return i.Validate()
}

// BeforeUpdate hook executed before database update operation.
func (i *Processor) BeforeUpdate(db orm.DB) error {
	return i.Validate()
}

// Validate validates User struct and returns validation errors.
func (i *Processor) Validate() error {
	return nil
	//return validation.ValidateStruct(u,
	//	validation.Field(&u.Theme, validation.Required, validation.In("default", "dark")),
	//)
}
