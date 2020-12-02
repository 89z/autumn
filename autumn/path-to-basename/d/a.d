import std.path: baseName;
import std.stdio: writeln;

void main() {
   auto s = baseName(`C:\Windows\notepad.exe`);
   writeln(s == "notepad.exe");
}
