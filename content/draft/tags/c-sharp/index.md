---
title: C# language
tags: [draft]
---

## Mono

~~~
mcs.bat a.cs
~~~

Could not load type System.Numerics.Vector from assembly mscorlib

- <https://github.com/mono/mono/issues/19808>
- <https://mono-project.com/download>

## .NET Core

Users need 52 MB runtime, else they get:

~~~
The required library hostfxr.dll could not be found.
~~~

<https://dotnet.microsoft.com/download>

## .NET Framework

~~~
csc a.cs
~~~

Visual C# Command Line Compiler has stopped working

<https://github.com/dotnet/runtime/issues/36723>

## Roslyn

Could not load file or assembly

- <https://github.com/dotnet/roslyn/issues/42330>
- <https://nuget.org/packages/Microsoft.Net.Compilers>
