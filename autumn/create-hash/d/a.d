import std.digest.md;
import std.stdio;

void main() {
   auto s = "south north";
   auto t = s.md5Of.toHexString;
   writeln(t == "F53B1396FE2736A7E5C44629EE1A3AF5");
}
