<?php
$old_n = microtime(true);

while (true) {
   usleep(10_000);
   $diff_n = microtime(true) - $old_n;
   $min_n = (int)($diff_n / 60);
   $sec_n = $diff_n % 60;
   $mil_n = fmod($diff_n, 1) * 1000;
   printf("\r%d m %02d s %03d ms", $min_n, $sec_n, $mil_n);
}
