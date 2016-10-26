package mysql

/*
#include "event.h"
*/
import "C"
import (
	"unsafe"

	"github.com/RideLink-carshare/db2mb"
)

type MySQLEvent struct {
	ptr unsafe.Pointer
}

func NewEvent(event unsafe.Pointer) db2mb.Event {
	return &MySQLEvent{ptr: event}
}

func (event *MySQLEvent) GetType() int {
	return C.event.ptr
}
