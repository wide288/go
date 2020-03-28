package converter_complex

import "encoding/json"

// 注意结构属性要大字首字母
type Diy struct {
	Test1 string
	Test2 string
	Arrs []string
}

func relation2json(rel Diy) (string, error) {
	data, err := json.Marshal(rel)
	if err != nil {
		return "", nil
	}
	return string(data), nil
}

func json2relation(data []byte) (Diy, error) {
	var result Diy

	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}