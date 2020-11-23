import
   std.datetime,
   std.file,
   std.stdio;

void main() {
   auto in_o = Clock.currTime;
   setTimes("a.d", in_o, in_o);
   auto out_o = timeLastModified("a.d");
   writeln(in_o == out_o);
}
