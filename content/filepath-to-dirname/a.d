import std.path, std.stdio;
void main() {
   auto s1 = `C:\Windows\write.exe`;
   auto s2 = s1.dirName;
   writeln(s2 == `C:\Windows`);
}
