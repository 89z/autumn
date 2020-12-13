<?php

function format(object $o): string {
   $min_n = $o->i;
   $sec_n = $o->s;
   $mil_n = (int)($o->f * 1000);
   return sprintf('%d m %02d s %03d ms', $min_n, $sec_n, $mil_n);
}

$old_o = new DateTime;

while (true) {
   usleep(10_000);
   $new_o = new DateTime;
   $diff_o = $new_o->diff($old_o);
   echo "\r", format($diff_o);
}
