import 'dart:io' show Platform;

void main() {
   var s = Platform.environment['USERPROFILE'];
   print(s == r'C:\Users\Steven');
}
