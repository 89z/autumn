import std.stdio;
void main() {
   int n1 = 10;
   int n2;
   if (n1 > 30) {
      n2 = 30;
   } else if (n1 > 20) {
      n2 = 20;
   } else {
      n2 = n1;
   }
   n2.writeln;
}
