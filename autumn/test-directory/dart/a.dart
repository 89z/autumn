import 'dart:io';

void main() {
   var b = FileSystemEntity.isDirectorySync(r'C:\Users');
   print(b);
}
