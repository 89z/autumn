import os
let s = "C:\\Windows\\write.exe"
let (s2, s3, s4) = s.splitFile
echo [s2 == "C:\\Windows", s3 == "write", s4 == ".exe"]
