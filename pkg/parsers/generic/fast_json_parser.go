
package generic

import (
  "github.com/tidwall/gjson"
  "sync"
  "time"
  "fmt"
)

type JSONParser struct {
      schema        []abstract.ColSchema
      resultPool    sync.Pool
      auxOpts       *AuxParserOpts
}

type ParseResult struct {
      Data map[string] interface{}
      Err error
}

func NewJSONParser(schema []abstract.ColSchema, opts *AuxParserOpts) *JSONParser {   // This create a new JSONParser instance
        return &JSONParser {
                  schema : schema,
                  resultPool : sync.Pool{  // Here reusing the same object
                      New : func () interface {
                          return &ParseResult {
                                  Data : make(map[string]interface{}, len(schema))
                          }
                      } 
                  },
                  auxOpts : opts,
        }
}

func (p *JSONParser) Parse(input []byte) (*ParseResult, err) {
        result := p.resultPool.Get().(*ParseResult)
        result.Data = make(map[string]interface{}, len(p.schema)) // This resets the data before parsing.
        result.Err = nil

        if !gjson.ValidBytes(input) {
                  return nil, fmt.Errorf("invalid json input")
        }

        for _, col := range p.schema {                                // This function is to convert values to expected data types usign ConvertValue
                  value := gjson.GetBytes(input, col.ColPath())
                  if !value.Exists(){
                              if col.Required {
                                          return nil, fmt.Errorf("required field %s is missing", col.ColPath())
                              }
                              continue
                  }
                  
                  parsed, err := p.ConvertValue(value, col.DataType)
                  if err != nil {
                          return nil, fmt.Errorf("field %s : %w", col.ColPath(), err)
                  }
                  result.Data[col.columnName] = parsed

        }
        
        if p.auxOpts.AddRest {
                rest := make(map[string]interface{})
                gjson.ParseBytes(input).ForEach(func(key, value gjson.Result) bool {
                              if !p.isSchemaField(key.String()){
                                          rest[key.string()] = value.Value()
                              }
                              return true
                })
                result.Data["_rest"] = rest
        }
        return result, nil 
}


func (p *JSONParser) ConvertValue(value gjson.Result, targetType string) (interface{}, error) {
              switch targetType {
                    case "string" : 
                            return value.String(), nil
                    case "int64" :
                            return value.Int(), nil
                    case "float64" :
                            return value.Float(), nil
                    case "bool" : 
                            return value.Bool(), nil
                    case "timestamp" :  
                    if value.Type == gjson.String {                                // This check is the timestamp is string : "2025-02-14T10:26:00Z"
                                      return time.Parse(time.RFC3339, value.String()) 
                            }
                            return time.Unix(value.Int(), 0), nil                 // Here we return if the timestamp if of Int type (UNIX TS): 1739145600
                    default : 
                            return value.Value(), nil
              } 
}


func (p *JSONParser) isSchemaField(field string) bool {             // This function is to check if the field is explicitly defined in the schema.
      for _, col = range p.schema {
              if col.ColPath() == field {
                          return true
              }
      }
      return false
}
