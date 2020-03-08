---
title: C#
stars: 43133
---

## DotNet

<https://docs.microsoft.com/dotnet/api?view=netframework-4.8>

## Mono

~~~
mcs.bat a.cs
~~~

<https://mono-project.com/download>

## Roslyn

Using this file:

{{< r "a.cs" >}}

I can compile like this:

~~~
tools/csc /r:tools/Microsoft.CodeAnalysis.CSharp.dll \
/r:C:/Windows/Microsoft.NET/Framework64/v4.0.30319/netstandard.dll a.cs
~~~

whats interesting is the resultant file requires these to be in the same
directory:

~~~
Microsoft.CodeAnalysis.dll
Microsoft.CodeAnalysis.CSharp.dll
~~~

even just having them in the PATH doesnt seem to do it.

<https://nuget.org/packages/Microsoft.Net.Compilers>

## Stars

<https://github.com/shadowsocks/shadowsocks-windows>
