<?php

# natural int
$s = bcdiv('7.0', '2');
var_dump($s === '3');

# natural int
$s = bcdiv('7', '2');
var_dump($s === '3');
