import 'dart:convert';

void main() {
   var a = ['/', '📗'];
   { // example 1
      var s = jsonEncode(a);
      print(s);
   }
   { // example 2
      var s = json.encode(a);
      print(s);
   }
   { // example 3
      var e = new JsonEncoder.withIndent(' ');
      var s = e.convert(a);
      print(s);
   }
}
