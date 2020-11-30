import 'dart:convert';

void main() {
   var s = '[10, 11]';
   var a = jsonDecode(s);
   print(a[0] == 10);
}
