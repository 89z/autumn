import std.stdio, urllibparse;

void main() {
   auto s = "west=left&east=right";
   auto m = s.parseQS;
   // ["west":["left"], "east":["right"]]
   m.writeln;
}
