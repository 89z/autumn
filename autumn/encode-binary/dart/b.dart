import 'dart:convert';

void main() {
   var a = [10, 11, 12, 13];
   var s = base64Encode(a);
   print(s == 'CgsMDQ==');
}
