// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VideoTagDao is the data access object for table video_tag.
type VideoTagDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns VideoTagColumns // columns contains all the column names of Table for convenient usage.
}

// VideoTagColumns defines and stores column names for table video_tag.
type VideoTagColumns struct {
	Id      string //
	VideoId string //
	TagId   string //
}

// videoTagColumns holds the columns for table video_tag.
var videoTagColumns = VideoTagColumns{
	Id:      "id",
	VideoId: "video_id",
	TagId:   "tag_id",
}

// NewVideoTagDao creates and returns a new DAO object for table data access.
func NewVideoTagDao() *VideoTagDao {
	return &VideoTagDao{
		group:   "default",
		table:   "video_tag",
		columns: videoTagColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VideoTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VideoTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VideoTagDao) Columns() VideoTagColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VideoTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VideoTagDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VideoTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
