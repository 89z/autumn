import io = std.stdio;
import opt = std.getopt;

int main(string[] a) {
   string pre_s, suf_s;
   auto o = opt.getopt(a, 'p', "prefix", &pre_s, 's', "suffix", &suf_s);

   if (a.length != 2) {
      opt.defaultGetoptPrinter("add [flags] <stem>", o.options);
      return 1;
   }

   string stem_s = a[1];
   io.writeln(pre_s ~ stem_s ~ suf_s);
   return 0;
}
