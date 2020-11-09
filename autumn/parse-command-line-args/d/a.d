import std.getopt, std.stdio;

int main(string[] a) {
   string s_pre, s_suf;

   auto o = getopt(a,
      'p', "prefix", &s_pre,
      's', "suffix", &s_suf
   );

   if (a.length != 2) {
      defaultGetoptPrinter("cat [flags] <stem>", o.options);
      return 1;
   }

   string s_stem = a[1];
   writeln(s_pre ~ s_stem ~ s_suf);
   return 0;
}
