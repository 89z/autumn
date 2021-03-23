<?php

function format(int $n): string {
   $n = intdiv($n, 1_000_000);
   $mil = $n % 1000;
   $n = intdiv($n, 1000);
   $sec = $n % 60;
   $min = intdiv($n, 60);
   return sprintf('%d m %02d s %03d ms', $min, $sec, $mil);
}

$then = hrtime(true);

while (true) {
   usleep(10_000);
   $now = hrtime(true);
   echo "\r", format($now - $then);
}
