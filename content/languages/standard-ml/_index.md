---
title: Standard ML
stars: 16722
---

## CakeML

If I compile like this:

~~~
$ make CC=x86_64-w64-mingw32-gcc
x86_64-w64-mingw32-gcc    -c -o cake.o cake.S
x86_64-w64-mingw32-gcc    -c -o basis_ffi.o basis_ffi.c
x86_64-w64-mingw32-gcc  cake.o basis_ffi.o   -o cake.exe
~~~

I get this result in a Cygwin console:

~~~
$ ./cake
sh: ./cake: cannot execute binary file: Exec format error
~~~

and this result in a Windows console:

~~~
cake.exe is not a valid Win32 application.
Access is denied.
~~~

<https://github.com/CakeML/cakeml>

## MLton

The project has no recent release.

<https://github.com/MLton/mlton>

## Poly/ML

This page says I can use `polyc` to build stand-alone application:

<https://polyml.org/FAQ#standalone>

but this release does not have that file:

<https://github.com/polyml/polyml/releases/download/v5.8/PolyML5.8-64bit.msi>

## Moscow ML

The project has no recent release.

<https://github.com/kfl/mosml>

## MLKit

A Windows version is not available.

<https://github.com/melsman/mlkit>

## Introduction

<https://cs.lmu.edu/~ray/notes/introml>

## Stars

<https://github.com/komeiji-satori/Dress>
