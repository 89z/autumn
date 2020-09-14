---
title: 'Talk:Variable to console'
---

Some languages have not set `ENABLE_VIRTUAL_TERMINAL_PROCESSING`:

<https://bugs.python.org/issue29059>

which means we would need to set:

~~~
[HKEY_CURRENT_USER\Console]
"VirtualTerminalLevel"=dword:00000001
~~~

but this breaks PowerShell 5.1 when doing `git status`. Instead we can use this:

<https://github.com/microsoft/terminal>

**Reference**:

<https://superuser.com/questions/413073>
