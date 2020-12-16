package main
type Slice []interface{}
type Map map[string]interface{}

func (m Map) M(s string) Map {
   return m[s].(map[string]interface{})
}

func (m Map) A(s string) Slice {
   return m[s].([]interface{})
}
