import
   io = std.stdio,
   json = std.json;

auto in_s = `{"U2": {"Boy": ["Twilight", "I Will Follow"]}}`;

void main() {
   auto m = json.parseJSON(in_s);
   auto out_s = m["U2"]["Boy"][0].str;
   io.writeln(out_s == "Twilight");
}
