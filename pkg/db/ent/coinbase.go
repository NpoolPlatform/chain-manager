// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CoinBase is the model entity for the CoinBase schema.
type CoinBase struct {
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
	// Presale holds the value of the "presale" field.
	Presale bool `json:"presale,omitempty"`
	// Unit holds the value of the "unit" field.
	Unit string `json:"unit,omitempty"`
	// Env holds the value of the "env" field.
	Env string `json:"env,omitempty"`
	// ReservedAmount holds the value of the "reserved_amount" field.
	ReservedAmount decimal.Decimal `json:"reserved_amount,omitempty"`
	// ForPay holds the value of the "for_pay" field.
	ForPay bool `json:"for_pay,omitempty"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinBase) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coinbase.FieldReservedAmount:
			values[i] = new(decimal.Decimal)
		case coinbase.FieldPresale, coinbase.FieldForPay, coinbase.FieldDisabled:
			values[i] = new(sql.NullBool)
		case coinbase.FieldCreatedAt, coinbase.FieldUpdatedAt, coinbase.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case coinbase.FieldName, coinbase.FieldLogo, coinbase.FieldUnit, coinbase.FieldEnv:
			values[i] = new(sql.NullString)
		case coinbase.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinBase", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinBase fields.
func (cb *CoinBase) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coinbase.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cb.ID = *value
			}
		case coinbase.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cb.CreatedAt = uint32(value.Int64)
			}
		case coinbase.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cb.UpdatedAt = uint32(value.Int64)
			}
		case coinbase.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cb.DeletedAt = uint32(value.Int64)
			}
		case coinbase.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cb.Name = value.String
			}
		case coinbase.FieldLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo", values[i])
			} else if value.Valid {
				cb.Logo = value.String
			}
		case coinbase.FieldPresale:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field presale", values[i])
			} else if value.Valid {
				cb.Presale = value.Bool
			}
		case coinbase.FieldUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unit", values[i])
			} else if value.Valid {
				cb.Unit = value.String
			}
		case coinbase.FieldEnv:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field env", values[i])
			} else if value.Valid {
				cb.Env = value.String
			}
		case coinbase.FieldReservedAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field reserved_amount", values[i])
			} else if value != nil {
				cb.ReservedAmount = *value
			}
		case coinbase.FieldForPay:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field for_pay", values[i])
			} else if value.Valid {
				cb.ForPay = value.Bool
			}
		case coinbase.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				cb.Disabled = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CoinBase.
// Note that you need to call CoinBase.Unwrap() before calling this method if this CoinBase
// was returned from a transaction, and the transaction was committed or rolled back.
func (cb *CoinBase) Update() *CoinBaseUpdateOne {
	return (&CoinBaseClient{config: cb.config}).UpdateOne(cb)
}

// Unwrap unwraps the CoinBase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cb *CoinBase) Unwrap() *CoinBase {
	_tx, ok := cb.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinBase is not a transactional entity")
	}
	cb.config.driver = _tx.drv
	return cb
}

// String implements the fmt.Stringer.
func (cb *CoinBase) String() string {
	var builder strings.Builder
	builder.WriteString("CoinBase(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cb.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", cb.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", cb.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", cb.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(cb.Name)
	builder.WriteString(", ")
	builder.WriteString("logo=")
	builder.WriteString(cb.Logo)
	builder.WriteString(", ")
	builder.WriteString("presale=")
	builder.WriteString(fmt.Sprintf("%v", cb.Presale))
	builder.WriteString(", ")
	builder.WriteString("unit=")
	builder.WriteString(cb.Unit)
	builder.WriteString(", ")
	builder.WriteString("env=")
	builder.WriteString(cb.Env)
	builder.WriteString(", ")
	builder.WriteString("reserved_amount=")
	builder.WriteString(fmt.Sprintf("%v", cb.ReservedAmount))
	builder.WriteString(", ")
	builder.WriteString("for_pay=")
	builder.WriteString(fmt.Sprintf("%v", cb.ForPay))
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", cb.Disabled))
	builder.WriteByte(')')
	return builder.String()
}

// CoinBases is a parsable slice of CoinBase.
type CoinBases []*CoinBase

func (cb CoinBases) config(cfg config) {
	for _i := range cb {
		cb[_i].config = cfg
	}
}
