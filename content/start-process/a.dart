import 'dart:io';

void main() {
   // example 1
   var a = ['run.dart'];
   Process.runSync('notepad', a);
   // example 2
   var a2 = ['--version'];
   var f = (ProcessResult o) => stdout.write(o.stderr);
   Process.run('dart', a2).then(f);
}
