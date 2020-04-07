<?php
$d1 = new DOMDocument;
$d1->loadHTMLFile('a.html');
$x1 = new DOMXPath($d1);
$q1 = $x1->query('//img[@src="a.jpg"]');
$s1 = $q1[0]->getAttribute('alt');
var_dump($s1);
