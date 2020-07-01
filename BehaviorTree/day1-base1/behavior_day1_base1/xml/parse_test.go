package xml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func data() string {
	return `<root>  <!--入口-->
	<selector>
		<sequence>  <!--顺序节点-->
			<condhp para="70,100"/> <!--条件节点 判断血量百分比范围;参数:[最低值,最高值);不小于:返回true,小于:返回flase-->
			<selector> 
				<skill para="10001,30"/> <!--行为节点 释放技能 参数:技能ID,消耗mp值 -->
				<skill para="10002,0"/>
			</selector>
		</sequence>
		<sequence> 
			<condhp para="30,70"/>
			<selector> 
				<skill para="10003,70"/>
				<skill para="10004,0"/>
			</selector>
		</sequence>
		<sequence> 
			<condhp para="0,30"/>
			<parallel> 
				<escape/>
				<eat para="101"/>
			</parallel>
		</sequence>
	</selector>
</root>`
}

func TestParseXMLDoc(t *testing.T) {
	ParseXMLDoc(xml.NewDecoder(strings.NewReader(data())))
}
