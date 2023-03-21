package messages

import (
)

// Information about Organization
type OrgInfo struct {
  Name          string          `json:"name"`
  Logo          string          `json:"logo"`
  URL           string          `json:"url"`
}

func NewOrgInfo() *OrgInfo {
  return &OrgInfo{}
}
