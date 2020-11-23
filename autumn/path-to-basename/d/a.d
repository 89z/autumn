import
   std.path,
   std.stdio;

void main() {
   auto s = baseName(`C:\Windows\notepad.exe`);
   writeln(s == "notepad.exe");
}
