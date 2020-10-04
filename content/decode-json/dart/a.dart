import 'dart:convert';

void main() {
   var s = '{"year": 2019, "month": 12}';
   var m = jsonDecode(s);
   print(m);
}
