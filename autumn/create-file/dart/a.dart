import 'dart:io';

void main() {
  final filename = 'file.txt';
  new File(filename).writeAsString('some content')
    .then((File file) {
      // Do something with the file.
    });
}
