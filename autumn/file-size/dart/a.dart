var x = File(path).lengthSync();
import 'dart:io';
final x = (await File(path).readAsBytes()).length;
