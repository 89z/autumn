import 'dart:io' as io;

File f = new io.File("/tmp/foo");
f.copySync("/tmp/bar");
f.deleteSync();
File f2 = new io.File("/tmp/bar");
f2.renameSync("/tmp/foo");
