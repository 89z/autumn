package main
import (
   "image"
   "image/draw"
   "image/jpeg"
   "os"
)
func main() {
   // files
   f1, _ := os.Open("p80eaiz8t0n41.jpg")
   f2, _ := os.Create("out.jpg")
   // rectangle
   r1 := image.Rect(800, 800, 1400, 1400)
   // images
   i1, _ := jpeg.Decode(f1)
   i2 := image.NewRGBA(r1)
   // write
   draw.Draw(i2, r1, i1, r1.Min, draw.Src)
   jpeg.Encode(f2, i2, nil)
}
