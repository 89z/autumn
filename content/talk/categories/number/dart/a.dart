import 'dart:math';

void main() {
   var nX = 7;
   var nY = 19;
   // example 1
   var n1 = pow(nX, nY);
   // example 2
   var n2 = 1;
   for (var n = 0; n < 19; n++) {
      n2 *= 7;
   }
   // print
   print([n1, n2]);
}
