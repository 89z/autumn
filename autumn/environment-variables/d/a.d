import
   io = std.stdio,
   std.process: environment;

void main() {
   auto s = environment.get("USERPROFILE");
   io.writeln(s == `C:\Users\Steven`);
}
