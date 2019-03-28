package conf

import (
	"log"
	"time"
)

var (
	LogFlag = log.LstdFlags

	PendingWriteNum = 2000
	MaxMsgLen uint32 = 4096
	HTTPTImeout = 10 * time.Second
	LenMsgLen = 2
	LittleEndian = false

	GoLen
)
