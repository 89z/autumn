import 'dart:io';
main() {
   Process.run('less', ['-V']).then((ProcessResult results) {
      stdout.write(results.stdout);
   });
}
