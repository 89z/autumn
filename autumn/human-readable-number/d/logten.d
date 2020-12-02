import
   fmt = std.format,
   io = std.stdio,
   math = std.math;

string numberFormat(real n) {
   auto n2 = cast(int) math.log10(n) / 3;
   return fmt.format("%.3f", n / 1e3 ^^ n2) ~ ["", " k", " M", " G"][n2];
}

void main() {
   auto s = numberFormat(9012345678);
   io.writeln(s == "9.012 G");
}
