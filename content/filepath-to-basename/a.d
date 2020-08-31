import std.path, std.stdio;
void main() {
   auto s1 = "C:\\Windows\\write.exe";
   auto s2 = s1.baseName;
   writeln(s2 == "write.exe");
}
