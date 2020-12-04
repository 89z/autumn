<?php
$old_n = microtime(true);

while (true) {
   usleep(10_000);
   $diff_n = microtime(true) - $old_n;
   $min_n = (int)($diff_n / 60);
   $sec_n = fmod($diff_n, 60);
   printf("%d m %5.2f s\r", $min_n, $sec_n);
}
