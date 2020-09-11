import std.path, std.stdio;

void main() {
   auto s = "C:\\Windows\\write.exe";
   auto s2 = s.baseName;
   writeln(s2 == "write.exe");
}
