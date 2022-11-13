package util

import jsoniter "github.com/json-iterator/go"

func TransferMap2Struct(ori map[string]interface{}, des interface{}) error {
	marshal, err := jsoniter.Marshal(ori)
	if err != nil {
		return err
	}
	err = jsoniter.Unmarshal(marshal, des)
	if err != nil {
		return err
	}
	return nil
}
