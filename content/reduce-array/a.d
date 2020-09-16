import std.algorithm, std.stdio;

void main() {
   auto a = ["May", "June"];
   auto f = (string s_acc, string s_cur) => s_acc ~ s_cur;
   // example 1
   auto s = a.reduce!(f);
   // example 2
   string s2;
   foreach (s_cur; a) {
      s2 = f(s2, s_cur);
   }
   // print
   writeln(s == "MayJune" && s2 == "MayJune");
}
