import 'dart:io';

void main() {
   var s = Platform.environment['BROWSER'];
   Process.runSync(s, ['https://dart.dev']);
}
