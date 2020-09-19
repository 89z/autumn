import 'dart:io';

void main() async {
   print('May');
   var o = await Process.start('test', []);
   o.stdout.pipe(stdout);
}
