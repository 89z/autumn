import
   std.path,
   std.stdio;

void main() {
   auto s = dirName(`C:\Windows\notepad.exe`);
   writeln(s == `C:\Windows`);
}
