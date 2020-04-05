---
title: Python
stars: 77460
---

The Cygwin version is garbage:

<https://sourceware.org/pipermail/cygwin/2020-March/244118.html>

Here is another option:

<https://python.org/downloads/windows>

With the Zip File, `pip` is missing. Here is fix:

~~~
curl -O https://bootstrap.pypa.io/get-pip.py
python get-pip.py
~~~

<https://packaging.python.org/tutorials/installing-packages>

After that `python/Scripts` will need to be added to the `PATH`. Alternatively,
we could just use the installer, but `PATH` would still need to be modified.

## Docs

Always include the `.html`:

<https://docs.python.org/library/email.compat32-message.html>

## Stars

<https://github.com/donnemartin/system-design-primer>
