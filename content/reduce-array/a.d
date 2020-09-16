import std.algorithm, std.stdio;

void main() {
   auto a = ["May", "June"];
   auto f = (string s_acc, string s_cur) => s_acc ~ s_cur;
   auto s = a.reduce!(f);
   writeln(s == "MayJune");
}
