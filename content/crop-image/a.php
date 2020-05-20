<?php
extension_loaded('gd') or die('gd');
# files
$s1 = 'a.jpg';
$s2 = 'b.jpg';
# rectangle
$m1 = ['x' => 0, 'y' => 0, 'width' => 800, 'height' => 800];
# images
$r1 = imagecreatefromjpeg($s1);
$r2 = imagecrop($r1, $m1);
# write
imagejpeg($r2, $s2);
