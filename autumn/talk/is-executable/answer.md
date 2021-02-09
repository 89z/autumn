I have seen three ways of doing this:

## GetBinaryTypeW

Used by [PHP] and [Rust].

[rust]://github.com/fitzgen/is_executable/blob/v1.0.1/src/lib.rs#L146
[php]://github.com/php/php-src/blob/php-8.0.2/win32/ioutil.c#L673

## GetFileAttributesW

Used by [Python] and [Nim]. Nim says all files are readable and executable, and
only differentiates on if a file is writable.

[nim]://github.com/nim-lang/Nim/blob/version-1-4/lib/pure/os.nim#L1632
[python]://github.com/python/cpython/blob/v3.9.1/Modules/posixmodule.c#L2904

## GetFileInformationByHandle

Used by [Ruby].

[ruby]://github.com/ruby/ruby/blob/v3_0_0/win32/win32.c#L5492
