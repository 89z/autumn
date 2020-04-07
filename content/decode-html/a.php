<?php
$s1 = file_get_contents('https://github.com/login');
# example 1
$d1 = new DOMDocument;
$d1->loadHTMLFile('a.html');
# exmaple 2
$d2 = new DOMDocument;
libxml_use_internal_errors(true);
$d2->loadHTML($s1);
# example 3
$t1 = new tidy;
$t1->parseString($s1);
$d3 = new DOMDocument;
$d3->loadHTML($t1);
# print
var_dump($d1, $d2, $d3);
