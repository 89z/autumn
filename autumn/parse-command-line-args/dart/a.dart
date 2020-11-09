import 'dart:io';
import 'package:args/args.dart';

void main(a) {
   var o = ArgParser();
   o.addOption('prefix', abbr: 'p', defaultsTo: '');
   o.addOption('suffix', abbr: 's', defaultsTo: '');
   var m = o.parse(a);

   if (m.rest.length != 1) {
      print('add [flags] <stem>\n' + o.usage);
      exit(1);
   }

   var s = m['prefix'] + m.rest[0] + m['suffix'];
   print(s);
}
