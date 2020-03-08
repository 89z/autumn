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

what are different way to get the value? The obvious choice would be HTML:

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

Another option is to split into array:

{{< r "a4.php" >}}

## Reference

- <https://php.net/funcref>
- <https://stitcher.io/blog/short-closures-in-php>

## Stars

<https://github.com/laravel/laravel>
