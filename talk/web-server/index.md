---
title: 'Talk:Web Server'
---

If you have a repository like this:

~~~
css/
fonts/
image/
js/
index.html
~~~

You can just run a normal server and be done with it:

~~~
php -S localhost:10
~~~

However in some cases, you will want to move the site into a subfolder:

~~~
docs/css/
docs/fonts/
docs/image/
docs/js/
docs/index.html
license.txt
readme.md
~~~

This allows you to separate website files, from other files. Then you can choose
**master branch /docs folder** with GitHub. The issue is, that the publish site
will have files like this:

~~~
/umber/js/umber.js
~~~

but locally, that file will be:

~~~
/js/umber.js
~~~

So the server needs to strip the prefix, in order for the request to succeed.
Go can do this pretty easily:

{{< r "../strip.go" >}}

but with PHP, you need to use a route script:

{{< r "../route.php" >}}

~~~
php -S localhost:10 route.php
~~~

The issue is, that likely the route script will be tucked away somewhere, and
not in the current directory. If that is the case, you get this:

~~~
[Sun Apr 26 09:57:29 2020] PHP Fatal error:  Unknown: Failed opening required
'route.php' (include_path='.:/usr/share/pear:/usr/share/php/php') in Unknown on
line 0
~~~

Since I am only using the route with one repository, I could just put it in the
local directory. Or I could put it in `/usr/share/php/php`. However I would
rather just call something like `serve.php` instead of
`php -S localhost:10 route.php`. Alternatively I could pull the command from
the shell history with `Ctrl + R`, but I still feel the best option is calling
a single script. In order to do that, we could need something like this:

{{< r "a.php" >}}

I dislike the idea of calling PHP from PHP. So another option would be to
reimplement the PHP server, in PHP:

{{< r "../serve.php" >}}

This example only responds with `Sunday` regardless of the request. A full
solution would need to read the request, and redirect if need be.
