# Is executable

## GetBinaryTypeW

php:

~~~
return GetBinaryTypeW(path, &type) ? 0 : -1;
~~~

- https://github.com/php/php-src/blob/php-8.0.2/win32/ioutil.c#L673
- https://www.php.net/manual/en/function.is-executable.php

rust:

- https://doc.rust-lang.org/std/fs/struct.Permissions.html
- https://doc.rust-lang.org/std/os/windows/fs/trait.MetadataExt.html#tymethod.file_attributes
- https://github.com/fitzgen/is_executable/blob/v1.0.1/src/lib.rs#L146
- https://hyperpolyglot.org/rust#files

## getFileAttributesW

python:

~~~
os.access('/etc/hosts', os.X_OK)
~~~

- https://docs.python.org/library/os.html#os.access
- https://github.com/python/cpython/blob/v3.9.1/Modules/posixmodule.c#L2904
- https://hyperpolyglot.org/scripting2#readable-writable-executable

d:

https://dlang.org/library/std/file/get_attributes.html

nim:

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

- https://github.com/nim-lang/Nim/blob/version-1-4/lib/pure/os.nim#L1632
- https://nim-lang.org/docs/os.html#FilePermission
- https://nim-lang.org/docs/os.html#getFilePermissions,string

## GetFileInformationByHandle

ruby:

~~~
File.executable?("/etc/hosts")
~~~

- https://docs.ruby-lang.org/en/master/File.html#method-c-executable-3F
- https://github.com/ruby/ruby/blob/v3_0_0/win32/win32.c#L5492
- https://hyperpolyglot.org/scripting2#readable-writable-executable

## none

go:

- https://github.com/golang/go/blob/724d0720/src/os/exec/lp_windows.go#L18-L27
- https://stackoverflow.com/questions/60128401
- https://stackoverflow.com/questions/63911706
