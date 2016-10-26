#ifndef _MYSQL_BINARYLOG_H_
#define _MYSQL_BINARYLOG_H_

#include "driver.h"
#include "util.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef void BinaryLog;

BinaryLog *NewBinaryLog(Driver *driver);

struct resp_t BinaryLog_Connect(BinaryLog *self);
struct resp_t BinaryLog_Disconnect(BinaryLog *self);
struct resp_t BinaryLog_SetPosition(BinaryLog *self, unsigned long position);

#ifdef __cplusplus
}
#endif
#endif
