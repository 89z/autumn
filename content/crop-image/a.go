package main
import (
   "image"
   "image/jpeg"
   "os"
)
func main() {
   // files
   f1, _ := os.Open("a.jpg")
   f2, _ := os.Create("b.jpg")
   // rectangle
   r1 := image.Rect(800, 800, 1400, 1400)
   // images
   i1, _ := jpeg.Decode(f1)
   i2 := i1.(*image.YCbCr).SubImage(r1)
   // write
   jpeg.Encode(f2, i2, nil)
}
