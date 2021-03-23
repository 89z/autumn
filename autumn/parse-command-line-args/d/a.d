import std.getopt, std.stdio;

int main(string[] a) {
   string pre, suf;
   auto opt = getopt(a, 'p', "prefix", &pre, 's', "suffix", &suf);

   if (a.length != 2) {
      defaultGetoptPrinter("add [flags] <stem>", opt.options);
      return 1;
   }

   string stem = a[1];
   writeln(pre ~ stem ~ suf);
   return 0;
}
