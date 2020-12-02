import
   algo = std.algorithm,
   io = std.stdio;

void main() {
   auto s = "March";
   auto b = algo.startsWith(s, "Ma");
   io.writeln(b);
}
