#include "api-accessor.h"

bool c_accessor_callback(Tvg_Paint paint, void *data)
{
  return go_accessor_callback(paint, data);
}
