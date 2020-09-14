import std.conv, std.stdio;

void main() {
   auto a = [10, 11];
   auto s = "";
   foreach (n; a) {
      if (s > "") {
         s ~= ",";
      }
      s ~= n.text;
   }
   writeln(s == "10,11");
}
