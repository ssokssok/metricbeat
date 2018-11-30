package file

import (
  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getFileAssets() ([]*esmodels.FileType, error) {
  m, err := getEsModelFileType()
  if err != nil {
    println(err)
    return nil, err
  }

  return m, nil
}


func getEsModelFileType() ([]*esmodels.FileType, error) {

  ma := make([]*esmodels.FileType, 0)

  prca, errprc := getProcAssets()
  if errprc != nil {
    return nil, errprc
  }

  println("Length of Process : ", len(prca))
  for _, prc := range prca {

    if prc.ExecutablePath == nil {
      continue
    }

    m := new(esmodels.FileType)

    err := getVersionInfo(*prc.ExecutablePath, m ) 
    if err != nil {
      continue
    }

    ma = append(ma, m)
  }

  return ma, nil  
}
