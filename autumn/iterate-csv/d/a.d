import
   csv = std.csv,
   io = std.stdio;

auto s = "Month,Day
January,Sunday
February,Monday";

void main() {
   auto a = csv.csvReader!(string[string])(s, null);
   foreach (m; a) {
      io.writeln(m);
   }
}
