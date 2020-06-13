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

## Roslyn

Could not load file or assembly

- <https://github.com/dotnet/roslyn/issues/42330>
- <https://nuget.org/packages/Microsoft.Net.Compilers>

## .NET

Users do need to download runtime, else they get:

~~~
The required library hostfxr.dll could not be found.
~~~

<https://dotnet.microsoft.com/download/dotnet>

However it is pretty small. Zip is 32 MB and Exe is 27 MB. It can be installed
in less than 1 minute from my testing.
