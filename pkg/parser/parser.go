package parser

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"regexp"
)

// parseKeys replaces all keys from `template`
func ParseKeys(regex string, template []byte, data map[string]interface{}, out *bytes.Buffer) {
	re := regexp.MustCompile(regex)
	matches := re.FindAllStringSubmatchIndex(string(template), -1)

	index := 0
	for _, key := range matches {
		out.Write(template[index:key[0]])
		out.Write([]byte(data[string(template[key[2]:key[3]])].(string)))
		index = key[1]
	}
	if index < len(template) {
		out.Write(template[index:])
	}
}

// readJson reads a JSON file and decodes it into a struct slice
func ReadJson(filename string) ([]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var t []interface{}
	if err = json.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return t, nil
}
