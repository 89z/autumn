import std.path, std.stdio;

void main() {
   auto s = `C:\Windows\notepad.exe`;
   auto t = s.baseName;
   writeln(t == "notepad.exe");
}
