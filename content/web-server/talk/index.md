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
