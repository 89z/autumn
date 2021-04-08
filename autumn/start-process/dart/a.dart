import 'dart:convert';
import 'dart:io';

void main() {
   var p = Process.runSync('dust', [], stdoutEncoding: utf8);
   stdout.write(p.stdout);
}
