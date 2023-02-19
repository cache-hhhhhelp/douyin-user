package model

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindOneByUserName(ctx context.Context, username string) (*User, error)
		FindMany(ctx context.Context, id []int64) ([]User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func toString(v []int64) string {
	str := "("
	for i := 0; i < len(v)-1; i++ {
		str += strconv.Itoa(int(v[i])) + ","
	}
	if len(v) != 0 {
		str += strconv.Itoa(int(v[len(v)-1]))
	}
	str += ")"
	return str
}

func (c customUserModel) FindMany(ctx context.Context, id []int64) ([]User, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` in "+toString(id), userRows, c.table)
	var resp []User
	err := c.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c customUserModel) FindOneByUserName(ctx context.Context, username string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, c.table)
	var resp User
	err := c.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}
