#ifndef _MYSQL_UTIL_H_
#define _MYSQL_UTIL_H_

#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

struct resp_t {
    bool ok;
    const char *message;
    void *data;
};

#ifdef __cplusplus
}
#endif
#endif
