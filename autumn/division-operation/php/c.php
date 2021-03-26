<?php

# natural float
$n = 7.0 / 2;
var_dump($n == 3.5);

# natural float
$n = 7 / 2;
var_dump($n == 3.5);

# force int
$n = (int)(7 / 2);
var_dump($n == 3);
