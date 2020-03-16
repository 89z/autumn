<?php
# files
$f1 = 'p80eaiz8t0n41.jpg';
$f2 = 'out.jpg';
# rectangle
$r1 = ['x' => 0, 'y' => 0, 'width' => 800, 'height' => 800];
# images
$i1 = imagecreatefromjpeg($f1);
$i2 = imagecrop($i1, $r1);
# write
imagejpeg($i2, $f2);
