// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VideosDao is the data access object for table videos.
type VideosDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns VideosColumns // columns contains all the column names of Table for convenient usage.
}

// VideosColumns defines and stores column names for table videos.
type VideosColumns struct {
	Id        string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
	Name      string //
	Url       string //
	CoverUrl  string //
	State     string // 0:审核中 1:审核完成
}

// videosColumns holds the columns for table videos.
var videosColumns = VideosColumns{
	Id:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Name:      "name",
	Url:       "url",
	CoverUrl:  "cover_url",
	State:     "state",
}

// NewVideosDao creates and returns a new DAO object for table data access.
func NewVideosDao() *VideosDao {
	return &VideosDao{
		group:   "default",
		table:   "videos",
		columns: videosColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VideosDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VideosDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VideosDao) Columns() VideosColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VideosDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VideosDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VideosDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
