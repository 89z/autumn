import
   std.csv,
   std.stdio;

void main() {
   auto s = "Month,Day
January,Sunday
February,Monday";

   auto a = s.csvReader!(string[string])(null);

   foreach (m; a) {
      m.writeln;
   }
}
