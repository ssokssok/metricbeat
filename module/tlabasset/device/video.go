package device

import (
  "os/exec"
  "encoding/json"

  "github.com/ssokssok/metricbeat/module/tlabasset/utils"
  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getVideoController() ([]*esmodels.VideoControllerType, error) {
  m, err := getEsModelVideoControllerType()
  if err != nil {
    println(err)
    return nil, err
  }

  return m, nil
}

func getVideoControllerType() (string, error) {

  m := new(esmodels.VideoControllerType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return "", err
  }

  buf := utils.GetContents(fn) 

  ma := make([]*esmodels.VideoControllerType, 0)

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

func getEsModelVideoControllerType() ([]*esmodels.VideoControllerType, error) {

  m := new(esmodels.VideoControllerType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return nil, err
  }

  buf := utils.GetContents(fn) 

  ma := make([]*esmodels.VideoControllerType, 0)

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
