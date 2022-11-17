// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/appcoin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppCoin is the model entity for the AppCoin schema.
type AppCoin struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Logo holds the value of the "logo" field.
	Logo string `json:"logo,omitempty"`
	// ForPay holds the value of the "for_pay" field.
	ForPay bool `json:"for_pay,omitempty"`
	// WithdrawAutoReviewAmount holds the value of the "withdraw_auto_review_amount" field.
	WithdrawAutoReviewAmount decimal.Decimal `json:"withdraw_auto_review_amount,omitempty"`
	// ProductPage holds the value of the "product_page" field.
	ProductPage string `json:"product_page,omitempty"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppCoin) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appcoin.FieldWithdrawAutoReviewAmount:
			values[i] = new(decimal.Decimal)
		case appcoin.FieldForPay, appcoin.FieldDisabled:
			values[i] = new(sql.NullBool)
		case appcoin.FieldCreatedAt, appcoin.FieldUpdatedAt, appcoin.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case appcoin.FieldName, appcoin.FieldLogo, appcoin.FieldProductPage:
			values[i] = new(sql.NullString)
		case appcoin.FieldID, appcoin.FieldAppID, appcoin.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppCoin", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppCoin fields.
func (ac *AppCoin) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appcoin.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ac.ID = *value
			}
		case appcoin.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = uint32(value.Int64)
			}
		case appcoin.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ac.UpdatedAt = uint32(value.Int64)
			}
		case appcoin.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ac.DeletedAt = uint32(value.Int64)
			}
		case appcoin.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ac.AppID = *value
			}
		case appcoin.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				ac.CoinTypeID = *value
			}
		case appcoin.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ac.Name = value.String
			}
		case appcoin.FieldLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo", values[i])
			} else if value.Valid {
				ac.Logo = value.String
			}
		case appcoin.FieldForPay:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field for_pay", values[i])
			} else if value.Valid {
				ac.ForPay = value.Bool
			}
		case appcoin.FieldWithdrawAutoReviewAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field withdraw_auto_review_amount", values[i])
			} else if value != nil {
				ac.WithdrawAutoReviewAmount = *value
			}
		case appcoin.FieldProductPage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_page", values[i])
			} else if value.Valid {
				ac.ProductPage = value.String
			}
		case appcoin.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				ac.Disabled = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppCoin.
// Note that you need to call AppCoin.Unwrap() before calling this method if this AppCoin
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AppCoin) Update() *AppCoinUpdateOne {
	return (&AppCoinClient{config: ac.config}).UpdateOne(ac)
}

// Unwrap unwraps the AppCoin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AppCoin) Unwrap() *AppCoin {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppCoin is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AppCoin) String() string {
	var builder strings.Builder
	builder.WriteString("AppCoin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ac.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ac.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ac.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.AppID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ac.Name)
	builder.WriteString(", ")
	builder.WriteString("logo=")
	builder.WriteString(ac.Logo)
	builder.WriteString(", ")
	builder.WriteString("for_pay=")
	builder.WriteString(fmt.Sprintf("%v", ac.ForPay))
	builder.WriteString(", ")
	builder.WriteString("withdraw_auto_review_amount=")
	builder.WriteString(fmt.Sprintf("%v", ac.WithdrawAutoReviewAmount))
	builder.WriteString(", ")
	builder.WriteString("product_page=")
	builder.WriteString(ac.ProductPage)
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", ac.Disabled))
	builder.WriteByte(')')
	return builder.String()
}

// AppCoins is a parsable slice of AppCoin.
type AppCoins []*AppCoin

func (ac AppCoins) config(cfg config) {
	for _i := range ac {
		ac[_i].config = cfg
	}
}
