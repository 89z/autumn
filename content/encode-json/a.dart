import 'dart:convert';
main() {
   var m = {'Sunday': 10};
   var s = jsonEncode(m);
   print(s);
}
