// +build !no_ldflags

package mysql

/*
#cgo CPPFLAGS: -I/home/guilherme/Downloads/mysql-binary-log-events-1.0.2-labs/libbinlogevents/include
#cgo CPPFLAGS: -I/home/guilherme/Downloads/mysql-binary-log-events-1.0.2-labs/libbinlogevents/export
#cgo CPPFLAGS: -I/home/guilherme/Downloads/mysql-binary-log-events-1.0.2-labs/bindings/include
#cgo CPPFLAGS: -I/home/guilherme/Downloads/mysql-5.7.16-linux-glibc2.5-x86_64/include/
#cgo LDFLAGS: -L/home/guilherme/Downloads/mysql-binary-log-events-1.0.2-labs/lib -lmysqlstream
#cgo LDFLAGS: -L/home/guilherme/Downloads/mysql-binary-log-events-1.0.2-labs/libbinlogevents/lib -lbinlogevents

#include "driver.h"
#include "binarylog.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/RideLink-carshare/db2mb"
)

type EventLog struct {
	driver unsafe.Pointer
	binlog unsafe.Pointer
}

func handleError(resp C.struct_resp_t) (bool, error) {
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

func (eventlog *EventLog) GetPosition() uint {
	position := C.BinaryLog_GetPosition(eventlog.binlog)

	return uint(position)
}

func (eventlog *EventLog) SetPosition(pos uint) error {
	resp := C.BinaryLog_SetPosition(eventlog.binlog, C.ulong(pos))

	ok, err := handleError(resp)
	if !ok {
		return err
	}

	return nil
}

func (eventlog *EventLog) WaitForEvent() (db2mb.Event, error) {
	resp := C.Driver_GetNextEvent(eventlog.driver)

	ok, err := handleError(resp)
	if !ok {
		return nil, err
	}

	if resp.data == nil {
		return eventlog.WaitForEvent()
	}

	event := NewEvent(resp.data)

	return event, nil
}
