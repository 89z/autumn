import fmt = std.format;
import io = std.stdio;

string numberFormat(real n) {
   auto n2 = n, n3 = 0;
   while (n2 >= 1e3) {
      n2 /= 1e3;
      n3++;
   }
   return fmt.format("%.3f", n2) ~ ["", " k", " M", " G"][n3];
}

void main() {
   auto s = numberFormat(9012345678);
   io.writeln(s == "9.012 G");
}
