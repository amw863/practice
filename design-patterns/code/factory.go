/*
 * @Author : wb
 * @Date : 2022/2/3 10:39 上午
 */

package code

type Parser interface {
	parse()
}

type JsonParse struct {
}

func (j JsonParse) parse() {

}

type XmlParse struct {
}

func (p XmlParse) parse() {

}

type RuleConfigSource struct {
}

//
func (s RuleConfigSource) createParse(configFormat string) Parser {
	if configFormat == "json" {
		return JsonParse{}
	} else if configFormat == "xml" {
		return XmlParse{}
	}

	return nil
}
func (s RuleConfigSource) reload() {
	//parse := s.createParse("")
	parse := RuleConfigParserFactory("")
	parse.parse()
}

// 工厂方法
type JsonParseFactory struct {
}

func (f JsonParseFactory) parse() {

}

type XmlParseFactory struct {
}

func (x XmlParseFactory) parse() {

}

// 工厂的工厂
func RuleConfigParserFactory(cf string) Parser {
	m := make(map[string]Parser)
	m["json"] = &JsonParseFactory{}
	m["xml"] = &XmlParseFactory{}

	return m[cf]
}
