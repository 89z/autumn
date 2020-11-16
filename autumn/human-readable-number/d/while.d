import std.conv, std.format, std.stdio;

string filesize(double in_n) {
   int out_n = 0;
   while (in_n > 1024) { 
      in_n /= 1024;
      out_n++;
   }
   return format("%.3f", in_n) ~ ["", " k", " M", " G"][out_n];
}

void main() {
   auto s = filesize(1264);
   writeln(s);
}
