import os
var s1 = "C:\\Windows\\write.exe"
let (s2, s3, s4) = splitFile(s1)
echo [s2 == "C:\\Windows", s3 == "write", s4 == ".exe"]
