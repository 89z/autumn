import 'dart:io';

void main() async {
   var p = await Process.start('dust', []);
   stdout.addStream(p.stdout);
}
