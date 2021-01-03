<?php

function id_encode(int $year): string {
   $t = new DateTime;
   $d = $t->setDate($year, 1, 1)->setTime(0, 0)->diff(new DateTime);
   $x = (string) $t->setTimestamp(0)->add($d)->getTimestamp();
   return base_convert($x, 10, 36);
}

echo id_encode(2021), "\n";
