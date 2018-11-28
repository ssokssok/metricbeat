package device

import (
  "os/exec"
  "encoding/json"

  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getBios() (*esmodels.PcBiosType, error) {
  m, err := getEsModelPcBiosType()
  if err != nil {
    println(err)
    return nil, err
  }

  return m, nil
}

func getPcBiosType() (string, error) {

  m := new(esmodels.PcBiosType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return "", err
  }

  buf := getContents(fn) 

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

func getEsModelPcBiosType() (*esmodels.PcBiosType, error) {

  m := new(esmodels.PcBiosType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return nil, err
  }

  buf := getContents(fn) 
  err = json.Unmarshal(buf, m)
  if err != nil {
    return nil, err
  }

  return m, nil  
}
