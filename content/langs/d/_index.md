---
title: D
---

## Build Tools for Visual Studio

First go here:

<https://visualstudio.microsoft.com/downloads>

Click **Tools for Visual Studio**. Then under **Build Tools for Visual Studio**,
click **Download**. Then run the following command. Prefer Command Prompt over
other shells as this will take a while. Warning, this takes about 45 minutes,
and size is around 22.1 GB:

~~~
vs_BuildTools --lang en-us --layout out
~~~

Extract:

~~~
lessmsi x netfx_NETCoreUWP.msi
~~~

First find the files:

~~~
fd msi -t f
~~~

Need to specify files, as some directories are named MSI as well:

~~~
Microsoft.VisualC.140.EspC.Msi,version=14.0.24245
~~~

## Issues

MinGW:

<https://github.com/ldc-developers/ldc/issues/3385>

Small build tools:

<https://github.com/microsoft/msbuild/issues/5197>

## Setup

<https://dlang.org/download.html>
