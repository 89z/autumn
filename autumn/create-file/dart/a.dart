import 'dart:io';

void main() {
   var file = new File('a.txt');
   var sink = file.openWrite();
   sink.close();
}
