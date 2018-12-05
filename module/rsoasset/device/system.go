package device

import (
  "os/exec"
  "encoding/json"

  "github.com/ssokssok/metricbeat/module/rsoasset/utils"
  "bitbucket.org/realsighton/rso/servers/common/esmodels"
)

func getSystem() (*esmodels.SystemType, error) {
  m, err := getEsModelSystemType()
  if err != nil {
    println(err)
    return nil, err
  }

  return m, nil
}

func getSystemType() (string, error) {

  m := new(esmodels.SystemType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return "", err
  }

  buf := utils.GetContents(fn) 

  err = json.Unmarshal(buf, m)
  if err != nil {
    return "", err
  }

  s := ""

  if b, err := json.Marshal(m); err == nil {
    s = string(b)
  } 

  return s, nil
}

func getEsModelSystemType() (*esmodels.SystemType, error) {

  m := new(esmodels.SystemType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return nil, err
  }

  buf := utils.GetContents(fn) 
  //println(string(buf))
  err = json.Unmarshal(buf, m)
  if err != nil {
    return nil, err
  }

  return m, nil  
}
