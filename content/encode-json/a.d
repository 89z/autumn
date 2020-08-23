import std.json, std.stdio;
void main() {
   auto m = ["📕/📙": 10];
   auto o = m.JSONValue;
   // example 1
   auto s1 = o.toJSON;
   // example 2
   auto s2 = o.toJSON(true);
   // example 3
   auto s3 = o.toJSON(true, JSONOptions.doNotEscapeSlashes);
   // print
   writeln(s1, s2, s3);
}
