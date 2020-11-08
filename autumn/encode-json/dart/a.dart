import 'dart:convert';

void main() {
   var a = ['/', '📗'];
   // example 1
   var s1 = jsonEncode(a);
   print(s1);
   // example 2
   var s2 = json.encode(a);
   print(s2);
   // example 3
   var o = new JsonEncoder.withIndent('   ');
   var s3 = o.convert(a);
   print(s3);
}
