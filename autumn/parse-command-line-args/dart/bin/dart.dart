import 'dart:io';
import 'package:args/args.dart';

void main(a) {
   var p = ArgParser();
   p.addOption('prefix', abbr: 'p', defaultsTo: '');
   p.addOption('suffix', abbr: 's', defaultsTo: '');
   var m = p.parse(a);

   if (m.rest.length != 1) {
      print('add [flags] <stem>\n' + p.usage);
      return;
   }

   var s = m['prefix'] + m.rest[0] + m['suffix'];
   print(s);
}
