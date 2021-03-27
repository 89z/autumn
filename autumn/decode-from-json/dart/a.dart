import 'dart:convert';

void main() {
   var s = '{"month": 12, "day": 31}';
   var m = jsonDecode(s);
   var n = m['day'];
   print(n == 31);
}
