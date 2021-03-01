void main() {
   var s = '1z141z3';
   var n = int.parse(s, radix: 36);
   print(n == 0xFFFFFFFF);
}
