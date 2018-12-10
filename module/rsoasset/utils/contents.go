package utils

import (
  "fmt"
  "os"
  "io/ioutil"
)


// GetContents is ...
func GetContents(fn string) []byte {
  path := fmt.Sprintf("%s", fn)

  buf, err := ioutil.ReadFile(path)
  if err != nil {
    println(err)
    return nil
  }
 
  os.Remove(path)
  
  return buf[3:]  // remove BOM
}

// GetJSONContents is ....
func GetJSONContents(fn string) []byte {
  path := fmt.Sprintf("%s", fn)

  // file, err := os.Open(path)
  // if err != nil {
  //   println("path:", path, "err:", err.Error())
  //   return nil
  // }

  // //file.Close()

  // buf, err := ioutil.ReadAll(file)
  // if err != nil {
  //   println("path:", path, "err:", err.Error())
  //   file.Close()
  //   return nil
  // }

  // file.Close()
  buf, err := ioutil.ReadFile(path)
  if err != nil {
    println("path :", path, "err:", err.Error())
    return nil
  }
 
  os.Remove(path)
  
  return buf
}

