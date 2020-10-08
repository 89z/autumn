import std.csv, std.stdio;

void main() {
   auto s = "Month,Day
January,Sunday
February,Monday";
   auto a = csvReader!(string[string])(s, null);
   foreach (m; a) {
      m["Day"].writeln;
   }
}
