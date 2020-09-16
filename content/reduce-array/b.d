import  std.stdio;

void main() {
   auto a = ["May", "June"];
   auto f = (string s_acc, string s_cur) => s_acc ~ s_cur;
   string s;
   foreach (s_cur; a) {
      s = f(s, s_cur);
   }
   writeln(s == "MayJune");
}
