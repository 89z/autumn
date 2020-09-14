import std.conv, std.stdio;

void main() {
   auto a = [10, 11];
   auto s = "";
   foreach (n, n2; a) {
      if (n > 0) {
         s ~= ",";
      }
      s ~= n2.text;
   }
   writeln(s == "10,11");
}
