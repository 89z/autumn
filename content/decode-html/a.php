<?php
$s1 = 'https://github.com/login';
$o1 = new DOMDocument;
libxml_use_internal_errors(true);
$o1->loadHTMLFile($s1);
var_dump($o1);
