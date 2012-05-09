// Copyright (c) 2012 Guillermo Estrada. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package gravatar provides access to gravatar API as
// defined in:
// 
// http://en.gravatar.com/site/implement/
//
// Check github for more details:
// http://github.com/phrozen/gravatar
package gravatar

import (
  "crypto/md5"
  "image"
  "image/jpeg"
  "io/ioutil"
  "fmt"
  "net/http"
  //"net/url"
  "strings"
)

const (
  base_url        = "http://www.gravatar.com"
  secure_base_url = "https://secure.gravatar.com"
)

var Default string = ""

type Gravatar struct {
  hash string
}

func NewGravatar(email string) *Gravatar {
  grvtr := new(Gravatar)
  email = strings.ToLower(strings.TrimSpace(email))
  hash := md5.New()
  hash.Write([]byte(email))
  grvtr.hash = fmt.Sprintf("%x", hash.Sum(nil))
  return grvtr
}

func (t Gravatar) Email(email string) {
  email = strings.ToLower(strings.TrimSpace(email))
  hash := md5.New()
  hash.Write([]byte(email))
  t.hash = fmt.Sprintf("%x", hash.Sum(nil))
}

func (t Gravatar) Hash() string {
  return t.hash
}

func (t Gravatar) Exist() (bool, error) {
  url := fmt.Sprintf("%s/avatar/%s?d=404", base_url, t.hash)
  response, err := http.Get(url)
  if err != nil {
    return false, err
  }
  return response.StatusCode != 404, nil
}

func (t Gravatar) Avatar(size int) (image.Image, error) {
  url := fmt.Sprintf("%s/avatar/%s?s=%v&d=%s", base_url, t.hash, size, Default)
  response, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer response.Body.Close()
  img, err := jpeg.Decode(response.Body)
  if err != nil {
    return nil, err
  }
  return img, nil
}

func (t Gravatar) AvatarURL(size int) string {
  return fmt.Sprintf("%s/avatar/%s?s=%v&d=%s", base_url, t.hash, size, Default)
}

func (t Gravatar) AvatarSecureURL(size int) string {
  return fmt.Sprintf("%s/avatar/%s?s=%v&d=%s", secure_base_url, t.hash, size, Default)
}

func (t Gravatar) Profile(format string) (string, error) {
  url := fmt.Sprintf("%s/%s.%s", base_url, t.hash, format)
  response, err := http.Get(url)
  if err != nil {
    return "", err
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return "", err
  }
  return string(body), nil
}

func (t Gravatar) ProfileURL() string {
  return fmt.Sprintf("%s/%s", base_url, t.hash)
}

