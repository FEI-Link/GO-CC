// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gocc/internal/app/demo/dao/internal"
)

// internalSysStateDao is internal type for wrapping internal DAO implements.
type internalSysStateDao = *internal.SysStateDao

// sysStateDao is the data access object for table sys_state.
// You can define custom methods on it to extend its functionality as you wish.
type sysStateDao struct {
	internalSysStateDao
}

var (
	// SysState is globally public accessible object for table sys_state operations.
	SysState = sysStateDao{
		internal.NewSysStateDao(),
	}
)

// Fill with you ideas below.
