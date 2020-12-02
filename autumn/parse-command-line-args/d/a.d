import std.getopt: defaultGetoptPrinter, getopt;
import std.stdio: writeln;

int main(string[] a) {
   string pre_s, suf_s;

   auto o = getopt(a,
      'p', "prefix", &pre_s,
      's', "suffix", &suf_s
   );

   if (a.length != 2) {
      defaultGetoptPrinter("add [flags] <stem>", o.options);
      return 1;
   }

   string stem_s = a[1];
   writeln(pre_s ~ stem_s ~ suf_s);
   return 0;
}
