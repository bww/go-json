package json

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

type rolesExample struct {
  A int `json:"a" roles:"public,private"`
  B int `json:"b" roles:"private"`
  C int `json:"c"`
}

func TestBasicRoles(t *testing.T) {
  var data []byte
  var err error
  
  ex := rolesExample{1, 2, 3}
  
  data, err = MarshalRole("public", ex)
  if assert.Nil(t, err, fmt.Sprintf("%v", err)) {
    assert.Equal(t, `{"a":1,"c":3}`, string(data))
  }
  
  data, err = MarshalRole("private", ex)
  if assert.Nil(t, err, fmt.Sprintf("%v", err)) {
    assert.Equal(t, `{"a":1,"b":2,"c":3}`, string(data))
  }
  
}

