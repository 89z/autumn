<?php
$s_f_green = "\e[32m";
$s_f_red = "\e[31m";
$s_end = "\e[m";
$s1 = 'Sunday';
# example 1
echo $s_f_green;
echo $s1;
echo $s_f_red;
echo $s1;
echo $s_end;
echo $s1, "\n";
# example 2
$s2 = $s_f_green . $s1;
$s3 = $s_f_red . $s1;
$s4 = $s_f_end . $s1;
echo $s2, $s3, $s4, "\n";
