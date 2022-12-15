// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/fiatcurrency"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FiatCurrency is the model entity for the FiatCurrency schema.
type FiatCurrency struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// FiatCurrencyTypeID holds the value of the "fiat_currency_type_id" field.
	FiatCurrencyTypeID uuid.UUID `json:"fiat_currency_type_id,omitempty"`
	// FeedType holds the value of the "feed_type" field.
	FeedType string `json:"feed_type,omitempty"`
	// MarketValueLow holds the value of the "market_value_low" field.
	MarketValueLow decimal.Decimal `json:"market_value_low,omitempty"`
	// MarketValueHigh holds the value of the "market_value_high" field.
	MarketValueHigh decimal.Decimal `json:"market_value_high,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FiatCurrency) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case fiatcurrency.FieldMarketValueLow, fiatcurrency.FieldMarketValueHigh:
			values[i] = new(decimal.Decimal)
		case fiatcurrency.FieldCreatedAt, fiatcurrency.FieldUpdatedAt, fiatcurrency.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case fiatcurrency.FieldFeedType:
			values[i] = new(sql.NullString)
		case fiatcurrency.FieldID, fiatcurrency.FieldFiatCurrencyTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FiatCurrency", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FiatCurrency fields.
func (fc *FiatCurrency) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fiatcurrency.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fc.ID = *value
			}
		case fiatcurrency.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fc.CreatedAt = uint32(value.Int64)
			}
		case fiatcurrency.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fc.UpdatedAt = uint32(value.Int64)
			}
		case fiatcurrency.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				fc.DeletedAt = uint32(value.Int64)
			}
		case fiatcurrency.FieldFiatCurrencyTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fiat_currency_type_id", values[i])
			} else if value != nil {
				fc.FiatCurrencyTypeID = *value
			}
		case fiatcurrency.FieldFeedType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feed_type", values[i])
			} else if value.Valid {
				fc.FeedType = value.String
			}
		case fiatcurrency.FieldMarketValueLow:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field market_value_low", values[i])
			} else if value != nil {
				fc.MarketValueLow = *value
			}
		case fiatcurrency.FieldMarketValueHigh:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field market_value_high", values[i])
			} else if value != nil {
				fc.MarketValueHigh = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this FiatCurrency.
// Note that you need to call FiatCurrency.Unwrap() before calling this method if this FiatCurrency
// was returned from a transaction, and the transaction was committed or rolled back.
func (fc *FiatCurrency) Update() *FiatCurrencyUpdateOne {
	return (&FiatCurrencyClient{config: fc.config}).UpdateOne(fc)
}

// Unwrap unwraps the FiatCurrency entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fc *FiatCurrency) Unwrap() *FiatCurrency {
	_tx, ok := fc.config.driver.(*txDriver)
	if !ok {
		panic("ent: FiatCurrency is not a transactional entity")
	}
	fc.config.driver = _tx.drv
	return fc
}

// String implements the fmt.Stringer.
func (fc *FiatCurrency) String() string {
	var builder strings.Builder
	builder.WriteString("FiatCurrency(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", fc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", fc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", fc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("fiat_currency_type_id=")
	builder.WriteString(fmt.Sprintf("%v", fc.FiatCurrencyTypeID))
	builder.WriteString(", ")
	builder.WriteString("feed_type=")
	builder.WriteString(fc.FeedType)
	builder.WriteString(", ")
	builder.WriteString("market_value_low=")
	builder.WriteString(fmt.Sprintf("%v", fc.MarketValueLow))
	builder.WriteString(", ")
	builder.WriteString("market_value_high=")
	builder.WriteString(fmt.Sprintf("%v", fc.MarketValueHigh))
	builder.WriteByte(')')
	return builder.String()
}

// FiatCurrencies is a parsable slice of FiatCurrency.
type FiatCurrencies []*FiatCurrency

func (fc FiatCurrencies) config(cfg config) {
	for _i := range fc {
		fc[_i].config = cfg
	}
}
