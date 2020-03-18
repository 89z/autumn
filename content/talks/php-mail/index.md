---
title: PHP mail
---

PHP mail:

<https://php.net/function.mail>

seems pretty desirable upon first glance. Doing this:

{{< r "a.php" >}}

is much simpler than doing this:

{{< r "b.php" >}}

but what that page doesnt says is that `mail` requires `sendmail`:

<https://en.wikipedia.org/wiki/Sendmail>

and Sendmail is not cheap requirements wise. Here are requirements:

~~~
   63,280 cygrunsrv
   60,044 cyrus-sasl
    4,132 libcom_err2
   20,804 libcrypt0
  883,388 libdb5.3
   59,248 libgc1
  102,828 libgssapi_krb5_2
2,235,304 libguile2.0_22
   73,344 libk5crypto3
  236,076 libkrb5_3
   14,424 libkrb5support0
   13,988 libltdl7
  149,684 libopenldap2_4_2
      108 libopenssl100
   23,112 libprocps7
  131,228 libsasl2_3
  385,472 libunistring2
   12,228 libwrap0
  231,416 m4
  459,228 make
  122,568 procmail
       32 procps
  304,620 procps-ng
  509,556 sendmail
   58,912 tcp_wrappers

       25 packages
6,155,024 bytes
~~~

and here is PHP cURL, not including packages already required by PHP:

~~~
   54,360 libbrotlicommon1
   18,272 libbrotlidec1
    4,132 libcom_err2
   20,804 libcrypt0
  233,752 libcurl4
  883,388 libdb5.3
  102,828 libgssapi_krb5_2
   50,980 libidn2_0
   73,344 libk5crypto3
  236,076 libkrb5_3
   14,424 libkrb5support0
   52,312 libnghttp2_14
  149,684 libopenldap2_4_2
      108 libopenssl100
   50,452 libpsl5
  131,228 libsasl2_3
  157,744 libssh4
   15,608 libssh-common
  385,472 libunistring2
   27,156 php-curl
   38,224 publicsuffix-list-dafsa

       21 packages
2,700,348 bytes
~~~

So if the starting point is PHP, then just using PHP cURL is the better
choice.
