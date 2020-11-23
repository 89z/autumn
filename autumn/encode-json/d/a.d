import
   std.json,
   std.stdio;

void main() {
   auto o = ["/", "📗"].JSONValue;
   // example 1
   auto s1 = o.toJSON;
   // example 2
   auto s2 = o.toJSON(true);
   // example 3
   auto s3 = o.toJSON(true, JSONOptions.doNotEscapeSlashes);
   // print
   writeln(s1, '\n', s2, '\n', s3);
}
