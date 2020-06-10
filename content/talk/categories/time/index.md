---
title: 'Category talk:Time'
---

current date string:

~~~
cove/git-train/git-board.php
43:$s_curr = strftime('%F');
~~~

current date number:

~~~
cove/youtube-insert/youtube-insert.php
78:$n_id_1 = time();
~~~

date string to object:

~~~
cove/youtube-playlist/youtube-playlist.php
43:   $o_then = date_create($s_then);
~~~

- <https://docs.python.org/library/time.html#time.strptime>
- <https://docs.python.org/library/datetime.html#datetime.datetime.strptime>

date difference:

~~~
cove/youtube-playlist/youtube-playlist.php
44:   $o_diff = date_diff($o_now, $o_then);
~~~

number to date string:

~~~
umber/docs/js/date.js
14:   const o = new Date(n_id * 1000);
15:   const s_day = o.toLocaleString(0, {weekday: 'short'});
~~~

not needed:

- current date object
- date object to number
- date object to string
- date string to number
- number to date object
