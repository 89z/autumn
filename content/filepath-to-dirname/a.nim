import os
let s = "C:\\Windows\\write.exe"
let (s1, s2, s3) = s.splitFile
echo [s1 == "C:\\Windows", s2 == "write", s3 == ".exe"]
