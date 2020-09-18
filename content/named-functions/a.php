<?php
# example A
function fA(int $nY, int $nZ): bool {
   return $nY > $nZ;
}
# example B
function fB($nY, $nZ) {
   return $nY > $nZ;
}
# print
$bY = fA(9, 8);
$bZ = fB(9, 8);
var_dump($bY, $bZ);
