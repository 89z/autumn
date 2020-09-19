import os
let s = "C:\\Windows\\write.exe"
let (s1, s1a, s1b) = s.splitFile
echo s1 == "C:\\Windows" and s1a == "write" and s1b == ".exe"
