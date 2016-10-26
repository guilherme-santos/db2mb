#ifndef _MYSQL_DRIVER_H_
#define _MYSQL_DRIVER_H_

#include "util.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef void Driver;

Driver *NewDriver(const char *url);

struct resp_t Driver_GetNextEvent(Driver *self);

#ifdef __cplusplus
}
#endif
#endif
