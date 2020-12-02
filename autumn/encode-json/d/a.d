import std.json: JSONOptions, JSONValue, toJSON;
import std.stdio: writeln;

void main() {
   auto o = ["/", "📗"].JSONValue;
   // example 1
   auto s1 = toJSON(o);
   // example 2
   auto s2 = toJSON(o, true);
   // example 3
   auto s3 = toJSON(o, true, JSONOptions.doNotEscapeSlashes);
   // print
   writeln(s1, '\n', s2, '\n', s3);
}
