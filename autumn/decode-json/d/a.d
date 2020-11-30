import
   std.json,
   std.stdio;

void main() {
   auto a = parseJSON("[10, 11]");
   auto n = a[0].integer;
   writeln(n == 10);
}
