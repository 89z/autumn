import 'dart:io';

void main() async {
   { // example 1
      var p = await Process.start('dust', []);
      stdout.addStream(p.stdout);
   }
   { // example 2
      var p = await Process.start('dust', [], workingDirectory: '..');
      stderr.addStream(p.stdout);
   }
}
