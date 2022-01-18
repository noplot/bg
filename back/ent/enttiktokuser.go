// Code generated by entc, DO NOT EDIT.

package ent

import (
	"back/ent/enttiktokuser"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// EntTikTokUser is the model entity for the EntTikTokUser schema.
type EntTikTokUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntTikTokUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case enttiktokuser.FieldID:
			values[i] = new(sql.NullInt64)
		case enttiktokuser.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EntTikTokUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntTikTokUser fields.
func (ettu *EntTikTokUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case enttiktokuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ettu.ID = int(value.Int64)
		case enttiktokuser.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ettu.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this EntTikTokUser.
// Note that you need to call EntTikTokUser.Unwrap() before calling this method if this EntTikTokUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (ettu *EntTikTokUser) Update() *EntTikTokUserUpdateOne {
	return (&EntTikTokUserClient{config: ettu.config}).UpdateOne(ettu)
}

// Unwrap unwraps the EntTikTokUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ettu *EntTikTokUser) Unwrap() *EntTikTokUser {
	tx, ok := ettu.config.driver.(*txDriver)
	if !ok {
		panic("ent: EntTikTokUser is not a transactional entity")
	}
	ettu.config.driver = tx.drv
	return ettu
}

// String implements the fmt.Stringer.
func (ettu *EntTikTokUser) String() string {
	var builder strings.Builder
	builder.WriteString("EntTikTokUser(")
	builder.WriteString(fmt.Sprintf("id=%v", ettu.ID))
	builder.WriteString(", name=")
	builder.WriteString(ettu.Name)
	builder.WriteByte(')')
	return builder.String()
}

// EntTikTokUsers is a parsable slice of EntTikTokUser.
type EntTikTokUsers []*EntTikTokUser

func (ettu EntTikTokUsers) config(cfg config) {
	for _i := range ettu {
		ettu[_i].config = cfg
	}
}