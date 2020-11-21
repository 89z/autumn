<?php
$old_o = new DateTime;

while (true) {
   usleep(10_000);
   $new_o = new DateTime;
   $diff_o = $new_o->diff($old_o);
   printf("%d m %5.2f s\r", $diff_o->i, $diff_o->s + $diff_o->f);
}
