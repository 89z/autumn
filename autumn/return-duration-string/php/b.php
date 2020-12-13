<?php

function format(int $n): string {
   $n = intdiv($n, 1_000_000);
   $mil_n = $n % 1000;
   $n = intdiv($n, 1000);
   $sec_n = $n % 60;
   $min_n = intdiv($n, 60);
   return sprintf('%d m %02d s %03d ms', $min_n, $sec_n, $mil_n);
}

$old_n = hrtime(true);

while (true) {
   usleep(10_000);
   $new_n = hrtime(true);
   echo "\r", format($new_n - $old_n);
}
