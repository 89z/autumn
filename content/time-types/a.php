<?php
# example 1
$o1 = date_create();
date_timestamp_set($o1, 86400);
# example 2
$o2 = date_create();
date_date_set($o2, 1970, 1, 2);
date_time_set($o2, 0, 0);
# print
var_dump($o1, $o2);
