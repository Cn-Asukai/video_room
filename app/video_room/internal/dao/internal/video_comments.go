// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VideoCommentsDao is the data access object for table video_comments.
type VideoCommentsDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns VideoCommentsColumns // columns contains all the column names of Table for convenient usage.
}

// VideoCommentsColumns defines and stores column names for table video_comments.
type VideoCommentsColumns struct {
	Id          string //
	VideoId     string //
	Content     string //
	Commentator string // 评论者id
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string //
}

// videoCommentsColumns holds the columns for table video_comments.
var videoCommentsColumns = VideoCommentsColumns{
	Id:          "id",
	VideoId:     "video_id",
	Content:     "content",
	Commentator: "commentator",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewVideoCommentsDao creates and returns a new DAO object for table data access.
func NewVideoCommentsDao() *VideoCommentsDao {
	return &VideoCommentsDao{
		group:   "default",
		table:   "video_comments",
		columns: videoCommentsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VideoCommentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VideoCommentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VideoCommentsDao) Columns() VideoCommentsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VideoCommentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VideoCommentsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VideoCommentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
