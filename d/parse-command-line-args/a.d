import std.getopt, std.stdio;

int main(string[] a) {
   string s_start, s_end;

   auto o = getopt(a,
      "s", "start", &s_start,
      "e", "end", &s_end
   );

   if (a.length != 2) {
      defaultGetoptPrinter("cat [flags] <input>", o.options);
      return 1;
   }

   string s_in = a[1];
   writeln(s_start ~ s_in ~ s_end);
   return 0;
}
