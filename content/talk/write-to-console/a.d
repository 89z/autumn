import std.stdio;
void main() {
   auto s_f_red = "\x1b[31m";
   auto s_end = "\x1b[m";
   writeln(s_f_red ~ "Sunday" ~ s_end);
}
