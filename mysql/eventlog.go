// +build !no_ldflags

package mysql

/*
#cgo CPPFLAGS: -I./mysql-binary-log-events/include -I/usr/include/mysql
#cgo LDFLAGS: -L./mysql-binary-log-events/lib -lmysqlstream -lbinlogevents

#include "driver.h"
#include "binarylog.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type EventLog struct {
	driver unsafe.Pointer
	binlog unsafe.Pointer
}

func handleError(resp C.struct_resp_error_t) (bool, error) {
	ok := bool(resp.ok)
	message := C.GoString(resp.message)

	return ok, errors.New(message)
}

func NewRemoteEventLog(host, user, password string, port uint) *EventLog {
	eventlog := &EventLog{}

	url := fmt.Sprintf("mysql://%s:%s@%s:%d", user, password, host, port)
	eventlog.driver = C.NewDriver(C.CString(url))
	eventlog.binlog = C.NewBinaryLog(eventlog.driver)

	return eventlog
}

func (eventlog *EventLog) Connect() error {
	resp := C.BinaryLog_Connect(eventlog.binlog)

	ok, err := handleError(resp)
	if !ok {
		return err
	}

	return nil
}

func (eventlog *EventLog) Disconnect() error {
	resp := C.BinaryLog_Disconnect(eventlog.binlog)

	ok, err := handleError(resp)
	if !ok {
		return err
	}

	return nil
}
