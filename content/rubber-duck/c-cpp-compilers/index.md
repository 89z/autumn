---
title: C and C++ compilers
---

{{< r "a.c" >}}

## musl GCC

This includes the GCC C compiler `gcc.exe` and C++ compiler `g++.exe`. It also
includes the musl C standard library. However it does not include a linker.

- <https://git.zv.io/toolchains/musl-cross-make/issues/1>
- <https://musl.cc>

## LLVM

This includes the GCC compatible C compiler `clang.exe` and C++ compiler
`clang++.exe`. It also includes the MSVC compatible C and C++ compiler
`clang-cl.exe`.

<https://releases.llvm.org/download.html>

You will need to add these to be able to compile:

- <https://packages.msys2.org/package/mingw-w64-x86_64-crt>
- <https://packages.msys2.org/package/mingw-w64-x86_64-headers>

~~~sh
# example 1
clang -c -I include a.c
lld-link -entry:main -defaultlib:lib/libmsvcrt.a a.o
# example 2
clang-cl /c /Zl /I include a.c
lld-link /entry:main /defaultlib:lib/libmsvcrt.a a.obj
~~~

## LLVM-MinGW

This includes the GCC compatible C compiler `clang.exe` and C++ compiler
`clang++.exe`. You can also copy `clang.exe` to `clang-cl.exe` to get a MSVC
compatible C and C++ compiler.

<https://github.com/mstorsjo/llvm-mingw>

## MSVC

Small build tools

<https://github.com/microsoft/msbuild/issues/5197>

## References

- <https://clang.llvm.org>
- <https://stackoverflow.com/questions/58998975>
