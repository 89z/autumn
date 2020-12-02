import
   array = std.array,
   io = std.stdio;

void main() {
   auto s = "M a r c h";
   // example 1
   auto a1 = array.split(s, " ");
   // example 2
   auto a2 = array.split(s);
   // print
   io.writeln(a1, a2);
}
