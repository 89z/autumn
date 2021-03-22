<?php

# example 1
$s = strtok('May,June', ',');
var_dump($s == 'May');

# example 2
$s = strtok(',');
var_dump($s == 'June');
