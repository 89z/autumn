import 'dart:io';

void main() async {
   print('BEGIN');
   var o = await Process.start('pipe', []);
   o.stdout.pipe(stdout);
}
