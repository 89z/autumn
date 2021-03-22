<?php

# example 1
$s = getenv('USERPROFILE');
var_dump($s == 'C:\Users\Steven');

# example 2
$s = getenv()['USERPROFILE'];
var_dump($s == 'C:\Users\Steven');
