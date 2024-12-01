// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomsDao is the data access object for table rooms.
type RoomsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns RoomsColumns // columns contains all the column names of Table for convenient usage.
}

// RoomsColumns defines and stores column names for table rooms.
type RoomsColumns struct {
	Id             string //
	CreatedAt      string //
	UpdatedAt      string //
	Creator        string // 创建者id
	Owner          string // 所有人id
	CurrentVideo   string // 当前播放的视频id
	Name           string // 房间名
	UserCount      string // 当前房间人数
	UserCountLimit string // 房间人数上限
	DeletedAt      string //
}

// roomsColumns holds the columns for table rooms.
var roomsColumns = RoomsColumns{
	Id:             "id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	Creator:        "creator",
	Owner:          "owner",
	CurrentVideo:   "current_video",
	Name:           "name",
	UserCount:      "user_count",
	UserCountLimit: "user_count_limit",
	DeletedAt:      "deleted_at",
}

// NewRoomsDao creates and returns a new DAO object for table data access.
func NewRoomsDao() *RoomsDao {
	return &RoomsDao{
		group:   "default",
		table:   "rooms",
		columns: roomsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomsDao) Columns() RoomsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
