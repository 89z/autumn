import
   std.json,
   std.stdio;

void main() {
   auto s = `{"month": 12, "day": 31}`;
   auto m = s.parseJSON;
   writeln(m["day"].integer == 31);
}
