import std.format, std.stdio;

string numberFormat(double n) {
   int n2 = 0;
   while (n >= 1e3) {
      n /= 1e3;
      n2++;
   }
   return format("%.3f", n) ~ ["", " k", " M", " G"][n2];
}

void main() {
   auto s = numberFormat(9012345678);
   writeln(s == "9.012 G");
}
