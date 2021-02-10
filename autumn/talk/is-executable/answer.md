I have seen four ways of doing this:

## [AccessCheck][check]

Based on this comment:

> Windows has `AccessCheck` with which you can check for `FILE_GENERIC_EXECUTE`
> access. First you have to call `OpenThreadToken` (or, if that fails,
> `OpenProcessToken` and `DuplicateToken`) to get an impersonation token and
> `GetNamedSecurityInfoW` to get the file security (owner, label, attribute, and
> DACL).

I found a Go package with the function [UserHasPermission][winox]. I will try to
work up an example, as I think this is the correct method to do this.

[winox]://pkg.go.dev/github.com/itchio/ox/winox#UserHasPermission
[check]://docs.microsoft.com/windows/win32/api/securitybaseapi/nf-securitybaseapi-accesscheck

## [GetBinaryTypeW][bin]

Used by [PHP] and [Rust]. This checks if a file if of executable **type**, but
does check executable **permission**.

[bin]://docs.microsoft.com/windows/win32/api/winbase/nf-winbase-getbinarytypew
[php]://github.com/php/php-src/blob/php-8.0.2/win32/ioutil.c#L673
[rust]://github.com/fitzgen/is_executable/blob/v1.0.1/src/lib.rs#L146

## [GetFileAttributesW][attr]

Used by [Go] and [Python] and [Nim]. The logic here is all files are readable and
executable, and only differentiates on if a file is writable.

[attr]://docs.microsoft.com/windows/win32/api/fileapi/nf-fileapi-getfileattributesw
[go]://github.com/golang/go/blob/go1.15.8/src/syscall/syscall_windows.go#L624
[nim]://github.com/nim-lang/Nim/blob/version-1-4/lib/pure/os.nim#L1632
[python]://github.com/python/cpython/blob/v3.9.1/Modules/posixmodule.c#L2904

## [GetFileInformationByHandle][info]

Used by [Ruby]. The logic here is all files are readable and executable, and only
differentiates on if a file is writable.

[info]://docs.microsoft.com/windows/win32/api/fileapi/nf-fileapi-getfileinformationbyhandle
[ruby]://github.com/ruby/ruby/blob/v3_0_0/win32/win32.c#L5492
