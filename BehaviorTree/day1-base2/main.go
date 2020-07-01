package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"
)

//xml解析

func getData() string {
	return `<extension name="rtp_multicast_page">` +
		`<condition field="destination_number" expression="^pagegroup$|^7243$">` +
		`<!-- comment -->` +
		`<action application="answer">raw text</action>` +
		`<action application="esf_page_group"/>` +
		`</condition>` +
		`</extension>`
}

func main() {

	dec := xml.NewDecoder(strings.NewReader(getData()))

	for {
		token, err := dec.Token()
		if err != nil {
			//是否读到结尾
			if errors.Is(err, io.EOF) {
				break
			}
			return
		}

		switch t := token.(type) {
		case xml.StartElement:
			fmt.Println("===> 1", t.Name.Local)
			fmt.Println("========================== start ==============================")
			fmt.Println("name ", t.Name.Local, " space ", t.Name.Space)
			for _, v := range t.Attr {
				fmt.Println("attr ", v.Name.Local, " space ", v.Name.Space, " value ", v.Value)
			}
			fmt.Println("===============================================================")
		case xml.EndElement:
			fmt.Println("===> 2", t.Name.Local)
			fmt.Println("========================== end ==============================")
			fmt.Println("name ", t.Name.Local, " space ", t.Name.Space)
			fmt.Println("=============================================================")
		case xml.CharData:
			fmt.Println("===> 3", string(t))
			fmt.Println("========================== data ==============================")
			fmt.Println("text ", string(t))
			fmt.Println("==============================================================")
		case xml.Comment:
			fmt.Println("===> 4", string(t))
			fmt.Println("========================== comment ==============================")
			fmt.Println("text ", string(t))
			fmt.Println("=================================================================")
		}

	}

}
