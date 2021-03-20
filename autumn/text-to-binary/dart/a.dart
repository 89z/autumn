import 'dart:convert';

void main() {
   var s = 'CgsMDQ==';
   var a = base64Decode(s);
   print(a);
}
