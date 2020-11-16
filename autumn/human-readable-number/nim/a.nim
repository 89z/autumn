import strformat

proc filesize(size: float): string =
   var left = size
   var out_n = 0
   while left > 1024:
      left /= 1024
      out_n += 1
   return fmt"{left:.3f}" & ["", " k", " M", " G"][out_n]

echo filesize(1264)
