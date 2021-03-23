import std.json, std.stdio;

auto src = `{"U2": {"Boy": ["Twilight", "I Will Follow"]}}`;

void main() {
   auto m = src.parseJSON;
   auto dst = m["U2"]["Boy"][0].str;
   writeln(dst == "Twilight");
}
