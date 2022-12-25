package record

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RecordModel = (*customRecordModel)(nil)

type (
	// RecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecordModel.
	RecordModel interface {
		recordModel
		Find(ctx context.Context, book string) (*[]Record, error)
	}

	customRecordModel struct {
		*defaultRecordModel
	}
)

// NewRecordModel returns a model for the database table.
func NewRecordModel(conn sqlx.SqlConn, c cache.CacheConf) RecordModel {
	return &customRecordModel{
		defaultRecordModel: newRecordModel(conn, c),
	}
}

func (m *defaultRecordModel) Find(ctx context.Context, book string) (*[]Record, error) {
	var resp []Record
	query := fmt.Sprintf("select %s from %s where book like $1", recordRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, "%"+book+"%")
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
