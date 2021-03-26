import 'dart:io';

void main() {
   var a = ['google.com/search?tbm=vid&q=squarepusher'];
   Process.runSync('waterfox', a);
}
