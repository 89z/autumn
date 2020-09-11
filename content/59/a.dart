import 'dart:convert';

void main() {
   var m = {'📗/📕': 10};
   // example 1
   var s = jsonEncode(m);
   // example 2
   var s2 = json.encode(m);
   // example 3
   var o = new JsonEncoder.withIndent('   ');
   var s3 = o.convert(m);
   // print
   print([s, s2, s3]);
}
