// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatMessagesDao is the data access object for table chat_messages.
type ChatMessagesDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ChatMessagesColumns // columns contains all the column names of Table for convenient usage.
}

// ChatMessagesColumns defines and stores column names for table chat_messages.
type ChatMessagesColumns struct {
	Id        string //
	CreatedAt string //
	Sender    string // 发展者id
	Receiver  string // 接收者id
	Seen      string // 是否已读
}

// chatMessagesColumns holds the columns for table chat_messages.
var chatMessagesColumns = ChatMessagesColumns{
	Id:        "id",
	CreatedAt: "created_at",
	Sender:    "sender",
	Receiver:  "receiver",
	Seen:      "seen",
}

// NewChatMessagesDao creates and returns a new DAO object for table data access.
func NewChatMessagesDao() *ChatMessagesDao {
	return &ChatMessagesDao{
		group:   "default",
		table:   "chat_messages",
		columns: chatMessagesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChatMessagesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChatMessagesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChatMessagesDao) Columns() ChatMessagesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChatMessagesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChatMessagesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChatMessagesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
