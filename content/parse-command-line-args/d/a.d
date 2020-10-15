import std.getopt, std.stdio;

int main(string[] a) {
   auto n_start = 0;
   auto n_len = 1;

   auto o = getopt(a,
      "start", "starting position", &n_start,
      "len", "substring length", &n_len
   );

   if (a.length != 2) {
      defaultGetoptPrinter("slice [flags] <string>", o.options);
      return 1;
   }

   auto s_in = a[1];
   auto s_out = s_in[n_start .. n_start + n_len];
   s_out.writeln;
   return 0;
}
