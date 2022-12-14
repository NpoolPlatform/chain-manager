package coinbase

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/chain-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/chain-manager/pkg/tracer/coin/base"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"

	"github.com/google/uuid"
)

func CreateSet(c *ent.CoinBaseCreate, in *npool.CoinBaseReq) *ent.CoinBaseCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.Name != nil {
		c.SetName(in.GetName())
	}
	if in.Logo != nil {
		c.SetLogo(in.GetLogo())
	}
	if in.Presale != nil {
		c.SetPresale(in.GetPresale())
	}
	if in.Unit != nil {
		c.SetUnit(in.GetUnit())
	}
	if in.ENV != nil {
		c.SetEnv(in.GetENV())
	}
	if in.ReservedAmount != nil {
		c.SetReservedAmount(decimal.RequireFromString(in.GetReservedAmount()))
	}
	if in.ForPay != nil {
		c.SetForPay(in.GetForPay())
	}
	if in.Disabled != nil {
		c.SetDisabled(in.GetDisabled())
	}
	return c
}

func Create(ctx context.Context, in *npool.CoinBaseReq) (*ent.CoinBase, error) {
	var info *ent.CoinBase
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.CoinBase.Create()
		info, err = CreateSet(c, in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.CoinBaseReq) ([]*ent.CoinBase, error) {
	var err error
	rows := []*ent.CoinBase{}

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.CoinBaseCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.CoinBase.Create(), info)
		}
		rows, err = tx.CoinBase.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(info *ent.CoinBase, in *npool.CoinBaseReq) *ent.CoinBaseUpdateOne {
	stm := info.Update()

	if in.Logo != nil {
		stm = stm.SetLogo(in.GetLogo())
	}
	if in.Presale != nil {
		stm = stm.SetPresale(in.GetPresale())
	}
	if in.ReservedAmount != nil {
		stm = stm.SetReservedAmount(decimal.RequireFromString(in.GetReservedAmount()))
	}
	if in.ForPay != nil {
		stm = stm.SetForPay(in.GetForPay())
	}
	if in.Disabled != nil {
		stm = stm.SetDisabled(in.GetDisabled())
	}

	return stm
}

func Update(ctx context.Context, in *npool.CoinBaseReq) (*ent.CoinBase, error) {
	var info *ent.CoinBase
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.CoinBase.Query().Where(coinbase.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query coinbase: %v", err)
		}

		stm := UpdateSet(info, in)

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update coinbase: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update coinbase: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.CoinBase, error) {
	var info *ent.CoinBase
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.CoinBase.Query().Where(coinbase.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.CoinBaseQuery, error) {
	stm := cli.CoinBase.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(coinbase.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.Name != nil {
		switch conds.GetName().GetOp() {
		case cruder.EQ:
			stm.Where(
				coinbase.Name(conds.GetName().GetValue()),
			)
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.ENV != nil {
		switch conds.GetENV().GetOp() {
		case cruder.EQ:
			stm.Where(coinbase.Env(conds.GetENV().GetValue()))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.Presale != nil {
		switch conds.GetPresale().GetOp() {
		case cruder.EQ:
			stm.Where(coinbase.Presale(conds.GetPresale().GetValue()))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.ForPay != nil {
		switch conds.GetForPay().GetOp() {
		case cruder.EQ:
			stm.Where(coinbase.ForPay(conds.GetForPay().GetValue()))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.Disabled != nil {
		switch conds.GetDisabled().GetOp() {
		case cruder.EQ:
			stm.Where(coinbase.Disabled(conds.GetDisabled().GetValue()))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.Names != nil {
		switch conds.GetNames().GetOp() {
		case cruder.IN:
			stm.Where(coinbase.NameIn(conds.GetNames().GetValue()...))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.CoinBase, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.CoinBase{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(coinbase.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.CoinBase, error) {
	var info *ent.CoinBase
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.CoinBase.Query().Where(coinbase.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.CoinBase, error) {
	var info *ent.CoinBase
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.CoinBase.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
