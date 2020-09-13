package main

import (
   "log"
   "os"
)

func main() {
   s := `C:\go\bin\go.exe`
   oa := os.ProcAttr{Files: []*os.File{nil, os.Stdout, nil}}
   op, e := os.StartProcess(s, []string{s, "version"}, &oa)
   if e != nil {
      log.Fatal(e)
   }
   op.Wait()
}
