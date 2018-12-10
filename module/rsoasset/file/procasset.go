package file

import (
  "os/exec"
  "encoding/json"

  "github.com/ssokssok/metricbeat/module/rsoasset/utils"
)

func getProcAssets() ([]*ProcessType, error) {
  m, err := getEsModelProcessType()
  if err != nil {
    println(err)
    return nil, err
  }

  return m, nil
}

func getEsModelProcessType() ([]*ProcessType, error) {

  m := new(ProcessType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  println("################# :", qry)
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return nil, err
  }

  buf := utils.GetContents(fn) 
  //println(string(buf))
  ma := make([]*ProcessType, 0)

  if buf[0] == '[' {
    err = json.Unmarshal(buf, &ma)
  } else {
    err = json.Unmarshal(buf, m)
    ma = append(ma, m)
  }

  if err != nil {
    println(err.Error())
    return nil, err
  }

  return ma, nil  
}
