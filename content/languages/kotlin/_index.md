---
title: Kotlin
stars: 35066
---

## Write to console

~~~java
fun main() {
   println("sunday")
}
~~~

## Issues

I tried running a command like this:

~~~
kotlinc.bat a.kt
~~~

and I get this result:

~~~
Downloading native dependencies (LLVM, sysroot etc). This is a one-time action 
performed only on the first run of the compiler.
Downloading dependency: https://download.jetbrains.com/kotlin/native/
   msys2-mingw-w64-x86_64-clang-llvm-lld-compiler_rt-8.0.1.zip
   (206.7 MiB/900.5 MiB).
~~~

This raises a couple of questions. First, why is this dependency 900 MB? Other
options are under 100 MB:

<https://musl.cc>

Second, should I choose to use my own C compiler, why cant I? It seems no option
is currently available to do that.
Choice of C compiler

<https://github.com/JetBrains/kotlin-native>

## Setup

- <https://github.com/JetBrains/kotlin/releases/latest>
- <https://jdk.java.net>

## Stars

<https://github.com/android/architecture-samples>
