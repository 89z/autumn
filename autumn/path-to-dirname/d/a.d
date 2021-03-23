import std.path, std.stdio;

void main() {
   auto s = `C:\Windows\notepad.exe`;
   auto t = s.dirName;
   writeln(t == `C:\Windows`);
}
