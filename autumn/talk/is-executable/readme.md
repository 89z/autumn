# Is executable

## php

~~~
return GetBinaryTypeW(path, &type) ? 0 : -1;
~~~

- https://github.com/php/php-src/blob/php-8.0.2/win32/ioutil.c#L673
- https://www.php.net/manual/en/function.is-executable.php

## go

- https://github.com/golang/go/blob/724d0720/src/os/exec/lp_windows.go#L18-L27
- https://stackoverflow.com/questions/60128401
- https://stackoverflow.com/questions/63911706

## d

https://dlang.org/library/std/file/get_attributes.html

## dart

https://api.dart.dev/dart-io/FileSystemEntity-class.html

## nim

~~~
READONLY
fpUserExec,               ## execute access for the file owner
fpUserRead,               ## read access for the file owner
fpGroupExec,              ## execute access for the group
fpGroupRead,              ## read access for the group
fpOthersExec,             ## execute access for others
fpOthersRead               ## read access for others

ELSE
fpUserExec,               ## execute access for the file owner
fpUserWrite,              ## write access for the file owner
fpUserRead,               ## read access for the file owner
fpGroupExec,              ## execute access for the group
fpGroupWrite,             ## write access for the group
fpGroupRead,              ## read access for the group
fpOthersExec,             ## execute access for others
fpOthersWrite,            ## write access for others
fpOthersRead               ## read access for others
~~~

- https://nim-lang.org/docs/os.html#FilePermission
- https://nim-lang.org/docs/os.html#getFilePermissions,string

## python

~~~
os.access('/etc/hosts', os.X_OK)
~~~

https://hyperpolyglot.org/scripting2#readable-writable-executable

## ruby

~~~
File.executable?("/etc/hosts")
~~~

https://hyperpolyglot.org/scripting2#readable-writable-executable

## rust

- https://doc.rust-lang.org/std/fs/struct.Permissions.html
- https://doc.rust-lang.org/std/os/windows/fs/trait.MetadataExt.html#tymethod.file_attributes
- https://hyperpolyglot.org/rust#files
