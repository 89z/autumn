import std.format: format;
import std.math: log10;
import std.stdio: writeln;

string numberFormat(real n) {
   auto n2 = cast(int)log10(n) / 3;
   return format("%.3f", n / 1e3 ^^ n2) ~ ["", " k", " M", " G"][n2];
}

void main() {
   auto s = numberFormat(9012345678);
   writeln(s == "9.012 G");
}
