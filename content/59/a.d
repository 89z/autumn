import std.json, std.stdio;

void main() {
   auto o = ["📗/📕": 10].JSONValue;
   // example 1
   auto s = o.toJSON;
   // example 2
   auto s2 = o.toJSON(true);
   // example 3
   auto s3 = o.toJSON(true, JSONOptions.doNotEscapeSlashes);
   // print
   writeln(s, s2, s3);
}
