import std.csv, std.stdio;

string[string][] f(string s_in) {
   auto a_in = csvReader!(string[string])(s_in, null);
   string[string][] a_out;
   foreach (m_row; a_in) {
      a_out ~= m_row;
   }
   return a_out;
}

void main() {
   auto s = "Month,Day
January,Sunday
February,Monday";
   auto a = s.f;
   a.writeln;
}
