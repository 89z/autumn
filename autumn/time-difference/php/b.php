<?php
$old_o = date_create();

while (true) {
   usleep(10_000);
   $new_o = date_create();
   $diff_o = date_diff($new_o, $old_o);
   printf("%d m %5.2f s\r", $diff_o->i, $diff_o->s + $diff_o->f);
}
