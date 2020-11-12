---
title: Python
---

## Python

<https://python.org/downloads/windows>

## Always include the .html

<https://docs.python.org/library/email.compat32-message.html>

## Pip

Even if explicitly stated that the embeddable [version of Python does not
support Pip][1], it is possible with care. You need to:

1. Download and unzip Python [embeddable zip][2] file.

2. In the file `python39._pth` or similar, uncomment the `import` command.
Result should look similar to this:

   ~~~
   python39.zip
   .
   import site
   ~~~

3. [Download get-pip.py][3] to the Python install folder

4. Run `get-pip.py`. this installs Pip into the `Scripts` directory:

   ~~~
   python get-pip.py
   ~~~

5. Run Pip directly from command line as Pip is a executable program (this
example is to install Pandas):

   ~~~
   .\Scripts\pip install pandas
   ~~~

You could find more information about this in the [Pip issue 4207][4]

[1]://docs.python.org/using/windows.html#windows-embeddable
[2]://python.org/downloads
[3]://pip.pypa.io/en/stable/installing
[4]://github.com/pypa/pip/issues/4207

<https://stackoverflow.com/questions/42666121#48906746>
