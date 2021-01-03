<?php
declare(strict_types = 1);

function id_encode(int $year): string {
   $t = new DateTime;
   $t->setDate($year, 1, 1);
   $x = time() - $t->getTimestamp();
   return base_convert((string) $x, 10, 36);
}

echo id_encode(2020), "\n";
