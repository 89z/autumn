<?php
# example 1
$n1 = time();
# example 2
$n2 = strtotime('now');
# example 3
$o = new DateTime('now');
$n3 = $o->getTimestamp();
# example 4
$o = date_create('now');
$n3 = date_timestamp_get($o);
