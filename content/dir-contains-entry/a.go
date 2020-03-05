import "os"

fi, err := os.Stat("/tmp/foo")
if os.IsNotExist(err) {
  fmt.Printf("Does not exit\n")
} else {
  fmt.Printf("Exists\n")
  fm := fi.Mode()
  if fm.IsRegular() {
    fmt.Printf("Is Regular")
  } else {
    fmt.Printf("Is Not Regular")
  }
}
