package zdi

import (
	"github.com/jigadhirasu/zzz/z"
	"gorm.io/gorm"
)

type DBDI func(db *gorm.DB) z.Bytes
