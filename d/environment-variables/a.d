import std.process, std.stdio;

void main() {
   auto s = environment.get("USERPROFILE");
   writeln(s == `C:\Users\Steven`);
}
