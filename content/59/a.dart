import 'dart:convert';
main() {
   var m = {'📗/📕': 10};
   // example 1
   var s1 = jsonEncode(m);
   print(s1);
   // example 2
   var s2 = json.encode(m);
   print(s2);
   // example 3
   var o = new JsonEncoder.withIndent('   ');
   var s3 = o.convert(m);
   print(s3);
}
