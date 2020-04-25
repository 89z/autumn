---
title: Web server
topics: [network]
languages: [go, php, python]
---

## Go

{{< r "serve.go" >}}

<https://golang.org/pkg/net/http#ListenAndServe>

{{< r "strip.go" >}}

<https://golang.org/pkg/net/http#StripPrefix>

## PHP

~~~
php -S localhost:10
~~~

<https://php.net/commandline.webserver#example-414>

## Python

~~~
python3 -m http.server --bind localhost 10
~~~

<https://docs.python.org/library/http.server.html>

## References

- <https://hyperpolyglot.org/scripting2#serve-pwd>
- <https://rosettacode.org/wiki/Hello_world/Web_server>
