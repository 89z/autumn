import 'dart:convert';
import 'package:crypto/crypto.dart';

void main() {
   var b = utf8.encode('south north');
   var s = md5.convert(b).toString();
   print(s == 'f53b1396fe2736a7e5c44629ee1a3af5');
}
