<?php
declare(strict_types = 1);
$a = new DateTime('2020-01-01');
$b = new DateTime('2020-12-31 23:59:59');
$c = $a->diff($b);
$d = (string) $a->setTimestamp(0)->add($c)->getTimestamp();
$e = base_convert($d, 10, 36);
echo $e, "\n";
