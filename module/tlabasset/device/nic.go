package device

import (
  "os/exec"
  "encoding/json"

  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getNics() ([]*esmodels.NicType, error) {
  m, err := getEsModelNicType()
  if err != nil {
    println(err)
    return nil, err
  }

  return m, nil
}

func getNicType() (string, error) {

  m := new(esmodels.NicType)

  qry, fn := m.GetPsWmiQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return "", err
  }

  buf := getContents(fn) 

  ma := make([]*esmodels.NicType, 0)

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

func getEsModelNicType() ([]*esmodels.NicType, error) {

  m := new(esmodels.NicType)

  qry, fn := m.GetPsWmiQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return nil, err
  }

  buf := getContents(fn) 

  ma := make([]*esmodels.NicType, 0)

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
