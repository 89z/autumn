<?php
$s_red = "\e[31m";
$s_gre = "\e[32m";
$s_end = "\e[m";
# example 1
echo $s_red;
echo 'Sunday';
echo $s_gre;
echo 'Monday';
echo $s_end;
echo "Tuesday\n";
# example 2
$s_in = $s_red . 'Sunday' . $s_gre . 'Monday' . $s_end . 'Tuesday';
echo $s_in, "\n";
