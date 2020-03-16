package main
import (
   "image"
   "image/jpeg"
   "os"
)
func main() {
   // files
   f1, _ := os.Open("p80eaiz8t0n41.jpg")
   f2, _ := os.Create("out.jpg")
   // rectangle
   r1 := image.Rect(0, 0, 800, 800)
   // images
   i1, _ := jpeg.Decode(f1)
   i2 := i1.(*image.YCbCr).SubImage(r1)
   // write
   jpeg.Encode(f2, i2, nil)
}
