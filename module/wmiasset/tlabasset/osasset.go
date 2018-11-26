package tlabasset

import (
  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getOsAsset() *esmodels.OsAssetType {

  gos := new(esmodels.OsAssetType)

  // OsType
  los, errlos := getOs()
  if errlos == nil {
    gos.Os = los
  }

  // Timezone
  tz, errtz := getTZ()
  if errtz == nil {
    gos.Timezone = tz
  }

  // Share
  sharea, errshare := getShares()
  if errshare == nil {
    gos.Shares = sharea
  }

  // UserAccount
  uaa, errua := getUserAccounts()

  if errua == nil {
    gos.UserAccounts = uaa
  }

  return gos
}

