#include <binlog.h>
#include "driver.h"

extern "C" {

using namespace binary_log::system;

Driver *NewDriver(const char *url) {
    Binary_log_driver *driver = create_transport(url);
    return static_cast<Driver*>(driver);
}

} // extern "C"
