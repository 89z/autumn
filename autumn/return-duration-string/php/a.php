<?php

function format(object $d): string {
   $min = $d->i;
   $sec = $d->s;
   $mil = (int)($d->f * 1000);
   return sprintf('%d m %02d s %03d ms', $min, $sec, $mil);
}

$then = new DateTime;

while (true) {
   usleep(10_000);
   $now = new DateTime;
   $diff = $now->diff($then);
   echo "\r", format($diff);
}
