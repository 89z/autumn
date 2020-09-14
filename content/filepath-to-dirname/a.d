import std.path, std.stdio;

void main() {
   auto s = `C:\Windows\write.exe`;
   auto s2 = s.dirName;
   writeln(s2 == `C:\Windows`);
}
