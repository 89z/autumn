import std.conv, std.format, std.stdio;

string numberFormat(double n) {
   int n2 = 0;
   while (n > 1024) {
      n /= 1024;
      n2++;
   }
   return format("%.3f", n) ~ ["", " k", " M", " G"][n2];
}

void main() {
   auto s = numberFormat(1264);
   writeln(s == "1.234 k");
}
