<?php
extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
$m1['reqtype'] = 'fileupload';
$m1['fileToUpload'] = new CURLFile('index.md');
$r1 = curl_init('https://catbox.moe/user/api.php');
curl_setopt($r1, CURLOPT_HTTP_VERSION, CURL_HTTP_VERSION_1_1);
curl_setopt($r1, CURLOPT_POST, true);
curl_setopt($r1, CURLOPT_POSTFIELDS, $m1);
curl_setopt($r1, CURLOPT_VERBOSE, true);
curl_exec($r1);
