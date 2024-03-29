// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DemoGenClassDao is the data access object for table demo_gen_class.
type DemoGenClassDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns DemoGenClassColumns // columns contains all the column names of Table for convenient usage.
}

// DemoGenClassColumns defines and stores column names for table demo_gen_class.
type DemoGenClassColumns struct {
	Id        string // 分类id
	ClassName string // 分类名
}

// demoGenClassColumns holds the columns for table demo_gen_class.
var demoGenClassColumns = DemoGenClassColumns{
	Id:        "id",
	ClassName: "class_name",
}

// NewDemoGenClassDao creates and returns a new DAO object for table data access.
func NewDemoGenClassDao() *DemoGenClassDao {
	return &DemoGenClassDao{
		group:   "default",
		table:   "demo_gen_class",
		columns: demoGenClassColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DemoGenClassDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DemoGenClassDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DemoGenClassDao) Columns() DemoGenClassColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DemoGenClassDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DemoGenClassDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DemoGenClassDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
