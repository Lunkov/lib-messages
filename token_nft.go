package messages

import (
)

type NFTToken struct {
  ID          uint64  `json:"id"`
  Name        string  `json:"name"`
  Description string  `json:"description"`
  ImageURL    string  `json:"image_url"`
  Owner       string  `json:"owner"`
}
