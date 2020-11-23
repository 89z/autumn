import std.path, std.stdio;

void main() {
   auto s = `C:\Windows\notepad.exe`;
   auto s1 = s.dirName;
   writeln(s1 == `C:\Windows`);
}
