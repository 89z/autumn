<?php
# files
$s1 = 'p80eaiz8t0n41.jpg';
$s2 = 'out.jpg';
# rectangle
$m1 = ['x' => 0, 'y' => 0, 'width' => 800, 'height' => 800];
# images
$r1 = imagecreatefromjpeg($s1);
$r2 = imagecrop($r1, $m1);
# write
imagejpeg($r2, $s2);
