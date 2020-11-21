import std.format, std.stdio;

string numberFormat(real n) {
   real n2 = n;
   int n3 = 0;
   while (n2 >= 1e3) {
      n2 /= 1e3;
      n3++;
   }
   return format("%.3f", n2) ~ ["", " k", " M", " G"][n3];
}

void main() {
   auto s = numberFormat(9012345678);
   writeln(s == "9.012 G");
}