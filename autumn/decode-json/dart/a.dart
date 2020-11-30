import 'dart:convert';

void main() {
   var s = '{"month": 12, "day": 31}';
   var m = jsonDecode(s);
   print(m['day'] == 31);
}
