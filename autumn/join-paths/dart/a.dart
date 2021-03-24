import 'package:path/path.dart';

void main() {
   var s = join('south', 'north');
   print(s == r'south\north');
}
