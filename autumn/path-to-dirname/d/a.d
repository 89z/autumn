import std.path: dirName;
import std.stdio: writeln;

void main() {
   auto s = dirName(`C:\Windows\notepad.exe`);
   writeln(s == `C:\Windows`);
}
