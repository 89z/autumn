import std.format, std.math, std.stdio;

string numberFormat(real n) {
   auto n2 = cast(int)log10(n) / 3;
   auto n3 = n / pow(1000, n2);
   return format("%.3f", n3) ~ ["", " k", " M", " G"][n2];
}

void main() {
   auto s = numberFormat(9012345678);
   writeln(s == "9.012 G");
}
