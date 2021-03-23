import std.format, std.math, std.stdio;

string numberFormat(real d) {
   auto e = cast(int) log10(d) / 3;
   return format("%.3f", d / 1e3 ^^ e) ~ ["", " k", " M", " G"][e];
}

void main() {
   auto s = numberFormat(9012345678);
   writeln(s == "9.012 G");
}
