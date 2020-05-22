package main
import (
   "image"
   "image/jpeg"
   "log"
   "os"
)
func main() {
   f1, e := os.Open("a.jpg")
   if e != nil {
      log.Fatal(e)
   }
   f2, e := os.Create("b.jpg")
   if e != nil {
      log.Fatal(e)
   }
   r1 := image.Rect(800, 800, 1400, 1400)
   i1, e := jpeg.Decode(f1)
   if e != nil {
      log.Fatal(e)
   }
   i2 := i1.(*image.YCbCr).SubImage(r1)
   jpeg.Encode(f2, i2, nil)
}
