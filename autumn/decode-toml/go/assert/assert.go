package assert
type Map map[string]interface{}

func (m Map) M(s string) Map {
   return m[s].(map[string]interface{})
}

func (m Map) S(s string) string {
   return m[s].(string)
}
