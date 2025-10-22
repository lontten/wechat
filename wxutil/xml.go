package wxutil

import "encoding/xml"

func ToCDataXml(v any) (string, error) {
	bytes, err := xml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type CDATA struct {
	Text string `xml:",cdata"`
}
