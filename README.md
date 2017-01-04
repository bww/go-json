# Go JSON

This is a replacement JSON marshaler/encoder which supports the conditional exclusion of properties. It is a modification of the standard library JSON encoder.

You can use this package instead of the standard marshaler in cases where you want to control which properties are included in the marshaled result.

    import (
      "github.com/bww/json"
    )
    
    type Example struct {
      A int `json:"a" roles:"public,private"`
      B int `json:"b" roles:"private"`
      C int `json:"c"`
    }
    
    func enc() {
      var data []byte
      
      ex := Example{
        1, 2, 3,
      }
      
      data, _ = json.MarshalRole("public", ex)
      fmt.Println(string(data)) // {"a":1,"C",3}
      
      data, _ = json.MarshalRole("private", ex)
      fmt.Println(string(data)) // {"a":1,"b":2,"c",3}
      
      data, _ = json.Marshal(ex) // for compatibility with encoding/json
      fmt.Println(string(data)) // {"a":1,"b":2,"c",3}
      
    }
