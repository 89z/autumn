---
title: Elixir
---

## Download

If I try to download a file, the average speed is 421k:

~~~
$ curl -O erlang.org/download/otp_win64_22.2.exe
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
 28 90.4M   28 25.7M    0     0   421k      0  0:03:39  0:01:02  0:02:37  433k
~~~

## Native

It seems HiPE is not currently available for Windows:

~~~
$ bin/erlc +native hello.erl
hello.erl: Warning: this system is not configured for native-code compilation.
~~~

<https://stackoverflow.com/questions/38708708/erlang-hipe-on-windows-10>

## Setup

- <https://elixir-lang.org/install#installing-erlang>
- <https://erlang-solutions.com/resources/download.html>
