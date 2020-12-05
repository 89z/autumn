import io = std.stdio;
import std.process: environment;

void main() {
   auto s = environment.get("USERPROFILE");
   io.writeln(s == `C:\Users\Steven`);
}
