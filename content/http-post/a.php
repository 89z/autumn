<?php
extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
# GET
$r1 = curl_init();
curl_setopt($r1, CURLOPT_COOKIEFILE, '');
curl_setopt($r1, CURLOPT_RETURNTRANSFER, true);
curl_setopt($r1, CURLOPT_URL, 'https://github.com');
$s_log = curl_exec($r1);
# POST
preg_match('/data-csrf="true" value="([^"]*)"/', $s_log, $a_auth);
$m_data['authenticity_token'] = $a_auth[1];
$m_data['value'] = 'aa540';
curl_setopt($r1, CURLOPT_POSTFIELDS, $m_data);
curl_setopt($r1, CURLOPT_RETURNTRANSFER, false);
curl_setopt($r1, CURLOPT_URL, 'https://github.com/signup_check/username');
curl_exec($r1);
