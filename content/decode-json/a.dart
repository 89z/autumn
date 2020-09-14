import 'dart:convert';

void main() {
   var s = '{"year": 2019, "month": 12, "day": 31}';
   var m = jsonDecode(s);
   print(m);
}
