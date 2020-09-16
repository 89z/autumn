import os
let s = "C:\\Windows\\write.exe"
let (s2, s3, s4) = s.splitFile
echo s2 == "C:\\Windows" and s3 == "write" and s4 == ".exe"
