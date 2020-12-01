import 'dart:io';

void main() {
   final x = (await File(path).readAsBytes()).length;
}
