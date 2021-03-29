<?php

# example 1
$s = strtok('west,east', ',');
var_dump($s == 'west');

# example 2
$s = strtok(',');
var_dump($s == 'east');
