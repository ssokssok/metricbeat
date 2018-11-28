package device

import (
  "os/exec"
  "encoding/json"

  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getDrives() ([]*esmodels.DriveType, error) {
  m, err := getEsModelDriveType()
  if err != nil {
    println(err)
    return nil, err
  }

  return m, nil
}

func getDriveType() (string, error) {

  m := new(esmodels.DriveType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return "", err
  }

  buf := getContents(fn) 

  ma := make([]*esmodels.DriveType, 0)

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

func getEsModelDriveType() ([]*esmodels.DriveType, error) {

  m := new(esmodels.DriveType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  out, err := cmd.CombinedOutput()
  if err != nil {
    println(out, err, qry)
    return nil, err
  }

  buf := getContents(fn)
  
  ma := make([]*esmodels.DriveType, 0)

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
