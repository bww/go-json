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

type singularExample struct {
  A int `json:"a" role:"public,private"`
  B int `json:"b" role:"private"`
  C int `json:"c"`
}

func TestBasicRoles(t *testing.T) {
  var data []byte
  var err error
  
  ex1 := rolesExample{1, 2, 3}
  ex2 := singularExample{1, 2, 3}
  
  data, err = MarshalRole("public", ex1)
  if assert.Nil(t, err, fmt.Sprintf("%v", err)) {
    assert.Equal(t, `{"a":1,"c":3}`, string(data))
  }
  
  data, err = MarshalRole("public", ex2)
  if assert.Nil(t, err, fmt.Sprintf("%v", err)) {
    assert.Equal(t, `{"a":1,"c":3}`, string(data))
  }
  
  data, err = MarshalRole("private", ex1)
  if assert.Nil(t, err, fmt.Sprintf("%v", err)) {
    assert.Equal(t, `{"a":1,"b":2,"c":3}`, string(data))
  }
  
  data, err = MarshalRole("private", ex2)
  if assert.Nil(t, err, fmt.Sprintf("%v", err)) {
    assert.Equal(t, `{"a":1,"b":2,"c":3}`, string(data))
  }
  
}

