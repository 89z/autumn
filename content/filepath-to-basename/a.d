import std.path, std.stdio;

void main() {
   auto s = "C:\\Windows\\notepad.exe";
   auto s1 = s.baseName;
   writeln(s1 == "notepad.exe");
}
