import 'dart:io';

void main() async {
   var p = await Process.start('dust', [], workingDirectory: '..');
   stdout.addStream(p.stdout);
}
