// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gocc/internal/app/demo/dao/internal"
)

// internalCcBumenRenwuDao is internal type for wrapping internal DAO implements.
type internalCcBumenRenwuDao = *internal.CcBumenRenwuDao

// ccBumenRenwuDao is the data access object for table cc_bumen_renwu.
// You can define custom methods on it to extend its functionality as you wish.
type ccBumenRenwuDao struct {
	internalCcBumenRenwuDao
}

var (
	// CcBumenRenwu is globally public accessible object for table cc_bumen_renwu operations.
	CcBumenRenwu = ccBumenRenwuDao{
		internal.NewCcBumenRenwuDao(),
	}
)

// Fill with you ideas below.
