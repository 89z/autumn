package main

import (
   "flag"
   "image"
   "image/jpeg"
   "log"
   "os"
   "path/filepath"
)

func main() {
   n_bottom := flag.Int("bottom", 0, "pixels")
   n_left := flag.Int("left", 0, "pixels")
   n_right := flag.Int("right", 0, "pixels")
   n_top := flag.Int("top", 0, "pixels")

   flag.Parse()
   if flag.NArg() != 1 {
      println("crop-image [flags] <file>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s_in := flag.Arg(0)
   s_out := "crop-" + filepath.Base(s_in)

   file_in, e := os.Open(s_in)
   if e != nil {
      log.Fatal(e)
   }

   file_out, e := os.Create(s_out)
   if e != nil {
      log.Fatal(e)
   }

   img_in, e := jpeg.Decode(file_in)
   if e != nil {
      log.Fatal(e)
   }

   rect_in := img_in.Bounds()
   rect_out := image.Rect(
      *n_left,
      *n_top,
      rect_in.Max.X - *n_right,
      rect_in.Max.Y - *n_bottom,
   )

   img_out := img_in.(*image.YCbCr).SubImage(rect_out)
   jpeg.Encode(file_out, img_out, nil)
}
