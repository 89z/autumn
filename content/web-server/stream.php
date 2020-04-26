<?php
$s_sock = 'localhost:10';
$s_msg = <<<eof
HTTP/1.1 200 OK

Sunday
eof;
$r_ser = stream_socket_server('tcp://' . $s_sock);
echo $s_sock, "\n";
while (true) {
   $r_acc = stream_socket_accept($r_ser);
   fwrite($r_acc, $s_msg);
   stream_socket_shutdown($r_acc, STREAM_SHUT_WR);
}
