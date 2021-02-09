<?php
declare(strict_types = 1);
symlink("/etc/hosts", "/tmp/hosts");
is_link("/etc/hosts")
readlink("/tmp/hosts")
