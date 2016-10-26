#ifndef _MYSQL_DRIVER_H_
#define _MYSQL_DRIVER_H_

#ifdef __cplusplus
extern "C" {
#endif

typedef void Driver;

Driver *NewDriver(const char *url);

#ifdef __cplusplus
}
#endif
#endif
