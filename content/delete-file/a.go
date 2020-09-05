import "os"

// no copy function in standard library

err := os.Remove("/tmp/foo")
if err != nil {
  fmt.Printf("Remove failed: %s\n", err)
}

err2 := os.Rename("/tmp/bar", "/tmp/foo")
if err2 != nil {
  fmt.Printf("Rename failed: %s\n", err2)
}
