import std.format, std.stdio;

string numberFormat(real d) {
   auto e = d, f = 0;
   while (e >= 1e3) {
      e /= 1e3;
      f++;
   }
   return format("%.3f", e) ~ ["", " k", " M", " G"][f];
}

void main() {
   auto s = numberFormat(9012345678);
   writeln(s == "9.012 G");
}
