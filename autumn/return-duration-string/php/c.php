<?php

function format(float $n): string {
   $min_n = (int)($n / 60);
   $sec_n = $n % 60;
   $mil_n = fmod($n, 1) * 1000;
   return sprintf('%d m %02d s %03d ms', $min_n, $sec_n, $mil_n);
}

$old_n = microtime(true);

while (true) {
   usleep(10_000);
   $new_n = microtime(true);
   echo "\r", format($new_n - $old_n);
}
