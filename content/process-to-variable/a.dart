import 'dart:io';

void main() {
   Process.run('less', ['-V']).then((ProcessResult results) {
      stdout.write(results.stdout);
   });
}
