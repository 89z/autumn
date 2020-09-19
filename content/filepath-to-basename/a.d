import std.path, std.stdio;

void main() {
   auto s = "C:\\Windows\\write.exe";
   auto s1 = s.baseName;
   writeln(s1 == "write.exe");
}
