import 'dart:convert';
main() {
   var s = '{"Sunday": 10}';
   var m = jsonDecode(s);
   print(m);
}
