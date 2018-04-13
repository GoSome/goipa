package ipa

import (
	"log"
	"encoding/json"

)

func (c *Client)ListAllGroup() ([]string, error) {
	options := map[string]interface{}{
		"sizelimit":0,
		"pkey_only": true,
	}
	var ret []string
	res,err := c.rpc("group_find",[]string{""},options)
	if err != nil {
		log.Fatal(err)
	}

	var Resp []interface{}
	err = json.Unmarshal(res.Result.Data, &Resp)
	if err != nil {
		return nil, err
	}
	for _,i := range Resp {
		for k,v :=range i.(map[string]interface{}){
			if k == "cn" {
				ret = append(ret,v.([]interface{})[0].(string))
			}
		}
	}
	return ret, nil

}