// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/fiatcurrencytype"
	"github.com/google/uuid"
)

// FiatCurrencyType is the model entity for the FiatCurrencyType schema.
type FiatCurrencyType struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Logo holds the value of the "logo" field.
	Logo string `json:"logo,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FiatCurrencyType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case fiatcurrencytype.FieldCreatedAt, fiatcurrencytype.FieldUpdatedAt, fiatcurrencytype.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case fiatcurrencytype.FieldName, fiatcurrencytype.FieldLogo:
			values[i] = new(sql.NullString)
		case fiatcurrencytype.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FiatCurrencyType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FiatCurrencyType fields.
func (fct *FiatCurrencyType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fiatcurrencytype.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fct.ID = *value
			}
		case fiatcurrencytype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fct.CreatedAt = uint32(value.Int64)
			}
		case fiatcurrencytype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fct.UpdatedAt = uint32(value.Int64)
			}
		case fiatcurrencytype.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				fct.DeletedAt = uint32(value.Int64)
			}
		case fiatcurrencytype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				fct.Name = value.String
			}
		case fiatcurrencytype.FieldLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo", values[i])
			} else if value.Valid {
				fct.Logo = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this FiatCurrencyType.
// Note that you need to call FiatCurrencyType.Unwrap() before calling this method if this FiatCurrencyType
// was returned from a transaction, and the transaction was committed or rolled back.
func (fct *FiatCurrencyType) Update() *FiatCurrencyTypeUpdateOne {
	return (&FiatCurrencyTypeClient{config: fct.config}).UpdateOne(fct)
}

// Unwrap unwraps the FiatCurrencyType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fct *FiatCurrencyType) Unwrap() *FiatCurrencyType {
	_tx, ok := fct.config.driver.(*txDriver)
	if !ok {
		panic("ent: FiatCurrencyType is not a transactional entity")
	}
	fct.config.driver = _tx.drv
	return fct
}

// String implements the fmt.Stringer.
func (fct *FiatCurrencyType) String() string {
	var builder strings.Builder
	builder.WriteString("FiatCurrencyType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fct.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", fct.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", fct.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", fct.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(fct.Name)
	builder.WriteString(", ")
	builder.WriteString("logo=")
	builder.WriteString(fct.Logo)
	builder.WriteByte(')')
	return builder.String()
}

// FiatCurrencyTypes is a parsable slice of FiatCurrencyType.
type FiatCurrencyTypes []*FiatCurrencyType

func (fct FiatCurrencyTypes) config(cfg config) {
	for _i := range fct {
		fct[_i].config = cfg
	}
}
