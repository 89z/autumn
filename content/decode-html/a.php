<?php
$s1 = 'https://github.com/login';
# example 1
$o1 = new DOMDocument;
libxml_use_internal_errors(true);
$o1->loadHTMLFile($s1);
# example 2
$o2 = new DOMDocument;
$o2->loadHTMLFile('https://github.com/login', LIBXML_NOERROR);
# print
var_dump($o1, $o2);
