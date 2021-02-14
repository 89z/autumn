<?php

function format(float $n): string {
   $min = (int)($n / 60);
   $sec = $n % 60;
   $mil = fmod($n, 1) * 1000;
   return sprintf('%d m %02d s %03d ms', $min, $sec, $mil);
}

$old = microtime(true);

while (true) {
   usleep(10_000);
   $new = microtime(true);
   echo "\r", format($new - $old);
}
