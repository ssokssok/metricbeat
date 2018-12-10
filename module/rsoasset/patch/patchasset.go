package patch

import (
  "os/exec"
  "encoding/json"

  "github.com/ssokssok/metricbeat/module/rsoasset/utils"
  "bitbucket.org/realsighton/rso/servers/common/esmodels"
)

var (
  old []*esmodels.PatchType
)

func statusInit() {
  old = make([]*esmodels.PatchType, 0)
}


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

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return "", err
  }

  buf := utils.GetContents(fn) 

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

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  println("################# :", qry)
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return nil, err
  }

  buf := utils.GetContents(fn) 
  println(string(buf))
  cur := make([]*esmodels.PatchType, 0)

  if buf[0] == '[' {
    err = json.Unmarshal(buf, &cur)
  } else {
    err = json.Unmarshal(buf, m)
    cur = append(cur, m)
  }

  if err != nil {
    println(err.Error())
    return nil, err
  }

  ma := make([]*esmodels.PatchType, 0)

  for _, v := range cur {
     f := findExistList(v)
     if !f {
       ma = append(ma, v)
     }
  }

  for _, mi := range ma {
    old = append(old, mi)
  }
  
  return ma, nil  
}


func findExistList(v *esmodels.PatchType) bool {
  if len(old) == 0 {
    return false
  }

  for _, ov := range old {
    if ov.Equals(v) {
      return true
    }
  }

  return false
}