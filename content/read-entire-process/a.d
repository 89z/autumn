import std.process, std.stdio;
void main() {
   const s1 = "cal";
   const s2 = executeShell(s1).output;
   write(s2);
}
