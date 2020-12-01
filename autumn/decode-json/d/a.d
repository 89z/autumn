import
   std.json,
   std.stdio;

auto in_s = `{"U2": {"Boy": ["Twilight", "I Will Follow"]}}`;

void main() {
   auto m = parseJSON(in_s);
   auto out_s = m["U2"]["Boy"][0].str;
   writeln(out_s == "Twilight");
}
