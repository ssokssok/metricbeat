package tlabasset

import (
  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getPatchs() []*esmodels.PatchType {

  // PatchType
  list, err := getPatchAssets()
  if err != nil {
    return nil
  }
  println("printer length: ", len(list))
  return list
}

