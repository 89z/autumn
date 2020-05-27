---
title: Common Lisp language
tags: [draft]
---

## Distros

**Clozure**:

<https://github.com/Clozure/ccl>

**Portacle**:

<https://github.com/portacle/portacle>

**Steel Bank**:

<https://sourceforge.net/projects/sbcl>

## Examples

{{< r "bool.lisp" >}}
{{< r "str.lisp" >}}

## HTTP

~~~
(ql:quickload "dexador")
~~~

<https://github.com/fukamachi/dexador>

## Issues

QuickList init speed

<https://github.com/quicklisp/quicklisp-client/issues/197>

## Setup

~~~
wx86cl -l quicklisp.lisp
(quicklisp-quickstart:install)
(ql:add-to-init-file)
~~~

<https://www.quicklisp.org>
