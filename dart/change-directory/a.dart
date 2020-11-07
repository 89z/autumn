import 'dart:io';

void main() {
   // example 1
   Directory.current = new Directory('..');
   print(Directory.current);
   // example 2
   Directory.current = 'dart';
   print(Directory.current);
}
