<?php
declare(strict_types = 1);

function id_encode(int $year): string {
   [$t, $u] = [new DateTime, new DateTime];
   $u->setDate($year, 1, 1);
   $x = $t->diff($u)->format('U');
   return base_convert($x, 10, 36);
}

function decode36(string $s): int {
   return (int) base_convert($s, 36, 10);
}
