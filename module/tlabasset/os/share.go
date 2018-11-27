package os

import (
  "os/exec"
  "encoding/json"

  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getShares() ([]*esmodels.ShareType, error) {
  m, err := getEsModelShareType()
  if err != nil {
    println(err)
    return nil, err
  }

  return m, nil
}

func getShareType() (string, error) {

  m := new(esmodels.ShareType)

  qry, fn := m.GetPsWmiQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return "", err
  }

  buf := getContents(fn) 

  ma := make([]*esmodels.ShareType, 0)

  if buf[0] == '[' {
    err = json.Unmarshal(buf, &ma)
  } else {
    err = json.Unmarshal(buf, m)
    ma = append(ma, m)
  }

  if err != nil {
    return "", err
  }

  s := ""

  if b, err := json.Marshal(ma); err == nil {
    s = string(b)
  } 

  return s, nil
}

func getEsModelShareType() ([]*esmodels.ShareType, error) {

  m := new(esmodels.ShareType)

  qry, fn := m.GetPsWmiQuery()
  
  execcmd := "PowerShell"
  println("######### qry: ", qry)
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return nil, err
  }

  buf := getContents(fn) 
  println(string(buf))
  ma := make([]*esmodels.ShareType, 0)

  if buf[0] == '[' {
    err = json.Unmarshal(buf, &ma)
  } else {
    err = json.Unmarshal(buf, m)
    ma = append(ma, m)
  }

  if err != nil {
    return nil, err
  }

  return ma, nil  
}
