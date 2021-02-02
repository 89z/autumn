<?php
$m = parse_url('https://example.com/one?two=even');
$s1 = $m['host'];
$s1 = parse_url($s, PHP_URL_HOST);
$s2 = $m['path'];
$s2 = parse_url($s, PHP_URL_PATH);
$s3 = $m['query'];
$s3 = parse_url($s, PHP_URL_QUERY);
