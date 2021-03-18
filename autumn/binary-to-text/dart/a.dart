import 'dart:convert';

void main() {
   var a = [10, 11, 12];
   var s = base64.encode(a);
   print(s == 'CgsM');
}
