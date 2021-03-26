<?php

function id_encode(int $year): string {
   $t = new DateTime;
   $d = $t->setDate($year, 1, 1)->setTime(0, 0)->diff(new DateTime);
   $s = (string) $t->setTimestamp(0)->add($d)->getTimestamp();
   return base_convert($s, 10, 36);
}

echo id_encode(2021), "\n";
