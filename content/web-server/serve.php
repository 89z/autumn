<?php
extension_loaded('sockets') or die('php-sockets');
$s_addr = 'localhost';
$n_port = 10;
$s_msg = <<<eof
HTTP/1.1 200 OK
Content-Length: 7

Sunday

eof;
$r_cr = socket_create(AF_INET, SOCK_STREAM, 0);
socket_bind($r_cr, $s_addr, $n_port);
socket_listen($r_cr);
echo $s_addr, ':', $n_port, "\n";
while (true) {
   $r_ac = socket_accept($r_cr);
   if ($r_ac != false) {
      socket_write($r_ac, $s_msg);
   } else {
      usleep(100000);
   }
}
