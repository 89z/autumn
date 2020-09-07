---
title: Java
tags: [reject]
---

## Boilerplate

Here is a minimal program:

{{< r "a.java" >}}

1. `class` is required:

~~~
error: class, interface, or enum expected
~~~

2. `public` is required:

~~~
error:  languagemain' method is not declared 'public static'
~~~

3. `static` is required:

~~~
error:  languagemain' method is not declared 'public static'
~~~

4. `void` is required:

~~~
error: invalid method declaration; return type required
~~~

5. `main` is required

~~~
error: cant find main(String[]) method in class
~~~

6. `String[]` is required

~~~
error: cant find main(String[]) method in class
~~~

<https://google.com/search?tbs=qdr:m&q=java+hello+world>

## Graal

If I try to use `native-image`, I get this:

~~~
> type HelloWorld.java
public class HelloWorld {
  public static void main(String[] args) {
    System.out.println("Hello, World!");
  }
}

> bin\javac HelloWorld.java

> bin\native-image HelloWorld
[helloworld:2924]    classlist:     960.11 ms,  1.19 GB
[helloworld:2924]        (cap):      84.78 ms,  1.19 GB
[helloworld:2924]        setup:     327.89 ms,  1.19 GB
Error: Unable to compile C-ABI query code. Make sure native software development
toolchain is installed on your system.
~~~

Running again gives this:

~~~
> bin\native-image -H:+ReportExceptionStackTraces HelloWorld
[...]
Caused by: java.io.IOException: Cannot run program "CL" (in directory
"C:\Users\Steven\AppData\Local\Temp\SVM-8913641224192734383"): CreateProcess
error=2, The system cannot find the file specified
   at java.lang.ProcessBuilder.start(ProcessBuilder.java:1048)
   at com.oracle.svm.hosted.c.codegen.CCompilerInvoker.
   startCommand(CCompilerInvoker.java:158)
~~~

Its true, I dont have `CL.exe` on my system, but I do have GCC. It seems
currently, with Windows only the MSVC compiler is allowed:

- <https://docs.microsoft.com/cpp/build/reference/compiling-a-c-cpp-program>
- <https://github.com/oracle/graal/blob/fb9b1d3b/substratevm/src/com.oracle.svm.hosted/src/com/oracle/svm/hosted/c/codegen/CCompilerInvoker.java#L106-L108>

I think people should be able to choose what compiler they use.

## Setup

<https://jdk.java.net>
