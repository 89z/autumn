import std.stdio;

string f(string sa, string sc) {
   return sa ~ sc;
}

void main() {
   auto a = ["May", "June"];
   string s;
   foreach (sc; a) {
      s = f(s, sc);
   }
   writeln(s == "MayJune");
}
