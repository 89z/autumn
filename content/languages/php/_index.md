---
title: PHP
stars: 55336
---

If we have a file like this:

<https://github.com/login>

part of it looks like this:

~~~
name="authenticity_token" value="WwmEvmIUMYN6m1cxUDh7GgeCKEwkAnGPRYcE2N/HAv5g7A"
~~~

what are different ways to get the value? The obvious choice would be HTML:

{{< r "a.php" >}}

but you get this:

~~~
PHP Warning: DOMDocument::loadHTMLFile(): Tag svg invalid
~~~

If you try this you get the same warning:

{{< r "a2.php" >}}

This was supposedly fixed in 2017:

<https://bugs.php.net/bug.php?id=74004>

Another option is `strtok`, but syntax is awkward:

{{< r "a3.php" >}}

Another option is to `foreach` an array:

{{< r "a4.php" >}}

Another option is to `array_search`:

{{< r "a5.php" >}}

Another option is regular expression:

{{< r "a6.php" >}}

Interestingly, regular expression is shortest code and fastest:

~~~
0.015600204467773 preg_match
0.029999971389771 strtok
0.03000020980835 foreach
0.031199932098389 array_search
~~~

## Reference

- <https://php.net/funcref>
- <https://stitcher.io/blog/short-closures-in-php>

## Stars

<https://github.com/laravel/laravel>
