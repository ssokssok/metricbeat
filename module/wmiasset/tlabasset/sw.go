package tlabasset

import (
  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getSws() []*esmodels.SwAssetType {

  // SystemType
  sws, err := getSwAssets()
  if err != nil {
    return nil
  }
  println("sws length: ", len(sws))
  return sws
}

