package os

import (
  "os/exec"
  "encoding/json"

  "github.com/ssokssok/metricbeat/module/rsoasset/utils"
  "bitbucket.org/realsighton/rso/servers/common/esmodels"
)

func getTZ() (*esmodels.TimeZoneType, error) {
  m, err := getEsModelTimezoneType()
  if err != nil {
    println(err)
    return nil, err
  }

  return m, nil
}

func getTimezoneType() (string, error) {

  m := new(esmodels.TimeZoneType)

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

func getEsModelTimezoneType() (*esmodels.TimeZoneType, error) {

  m := new(esmodels.TimeZoneType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  println("############ qry:", qry)
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
