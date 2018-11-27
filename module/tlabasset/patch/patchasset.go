package patch

import (
  "os/exec"
  "encoding/json"

  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getPatchAssets() ([]*esmodels.PatchType, error) {
  m, err := getEsModelPatchAssetType()
  if err != nil {
    println(err)
    return nil, err
  }

  return m, nil
}

func getPatchAssetType() (string, error) {

  m := new(esmodels.PatchType)

  qry, fn := m.GetPsWmiQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return "", err
  }

  buf := getContents(fn) 

  ma := make([]*esmodels.PatchType, 0)

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

func getEsModelPatchAssetType() ([]*esmodels.PatchType, error) {

  m := new(esmodels.PatchType)

  qry, fn := m.GetPsWmiQuery()
  
  execcmd := "PowerShell"
  println("################# :", qry)
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return nil, err
  }

  buf := getContents(fn) 
  println(string(buf))
  ma := make([]*esmodels.PatchType, 0)

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
