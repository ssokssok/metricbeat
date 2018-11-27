package device

import (
  "fmt"
  "os"
  "io/ioutil"
)



func getContents(fn string) []byte {
  path := fmt.Sprintf("%s", fn)

  buf, err := ioutil.ReadFile(path)
  if err != nil {
    println(err)
    return nil
  }
 
  os.Remove(path)
  
  return buf[3:]  // remove BOM
}
