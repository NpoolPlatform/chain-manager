// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppCoinsColumns holds the columns for the "app_coins" table.
	AppCoinsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "for_pay", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "withdraw_auto_review_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// AppCoinsTable holds the schema information for the "app_coins" table.
	AppCoinsTable = &schema.Table{
		Name:       "app_coins",
		Columns:    AppCoinsColumns,
		PrimaryKey: []*schema.Column{AppCoinsColumns[0]},
	}
	// CoinBasesColumns holds the columns for the "coin_bases" table.
	CoinBasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "presale", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "unit", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "env", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "reserved_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "for_pay", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// CoinBasesTable holds the schema information for the "coin_bases" table.
	CoinBasesTable = &schema.Table{
		Name:       "coin_bases",
		Columns:    CoinBasesColumns,
		PrimaryKey: []*schema.Column{CoinBasesColumns[0]},
	}
	// CoinDescriptionsColumns holds the columns for the "coin_descriptions" table.
	CoinDescriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// CoinDescriptionsTable holds the schema information for the "coin_descriptions" table.
	CoinDescriptionsTable = &schema.Table{
		Name:       "coin_descriptions",
		Columns:    CoinDescriptionsColumns,
		PrimaryKey: []*schema.Column{CoinDescriptionsColumns[0]},
	}
	// CoinExtrasColumns holds the columns for the "coin_extras" table.
	CoinExtrasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "home_page", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "specs", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// CoinExtrasTable holds the schema information for the "coin_extras" table.
	CoinExtrasTable = &schema.Table{
		Name:       "coin_extras",
		Columns:    CoinExtrasColumns,
		PrimaryKey: []*schema.Column{CoinExtrasColumns[0]},
	}
	// ExchangeRatesColumns holds the columns for the "exchange_rates" table.
	ExchangeRatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "market_value", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "settle_value", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "settle_percent", Type: field.TypeUint32, Nullable: true, Default: 100},
		{Name: "setter", Type: field.TypeUUID, Nullable: true},
	}
	// ExchangeRatesTable holds the schema information for the "exchange_rates" table.
	ExchangeRatesTable = &schema.Table{
		Name:       "exchange_rates",
		Columns:    ExchangeRatesColumns,
		PrimaryKey: []*schema.Column{ExchangeRatesColumns[0]},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "fee_coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "withdraw_fee_by_stable_usd", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "withdraw_fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "collect_fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "hot_wallet_fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "low_fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "warm_account_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "payment_account_collect_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
	}
	// TransColumns holds the columns for the "trans" table.
	TransColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "from_account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "to_account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "chain_tx_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultTxState"},
		{Name: "extra", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// TransTable holds the schema information for the "trans" table.
	TransTable = &schema.Table{
		Name:       "trans",
		Columns:    TransColumns,
		PrimaryKey: []*schema.Column{TransColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppCoinsTable,
		CoinBasesTable,
		CoinDescriptionsTable,
		CoinExtrasTable,
		ExchangeRatesTable,
		SettingsTable,
		TransTable,
	}
)

func init() {
}
