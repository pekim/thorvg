#include "api-picture.h"

bool c_picture_asset_resolver(Tvg_Paint paint, const char *src, void *data)
{
  return go_picture_asset_resolver(paint, (char *)src, data);
}
