package sample

import (
	"learn/abstractfactory/sample/factory"
	"learn/abstractfactory/sample/listfactory"
	"testing"
)

func TestListFactory(t *testing.T) {
	var f factory.Factory = &listfactory.ListFactory{}
	people := f.CreateLink("人民日报", "http://www.people.com.cn/")
	gmw := f.CreateLink("光明日报", "http://www.gmw.cn/")

	us_yahoo := f.CreateLink("Yahoo!", "http://www.yahoo.com/")
	jp_yahoo := f.CreateLink("Yahoo!Japan", "http://www.yahoo.co.jp/")
	excite := f.CreateLink("Excite", "http://www.excite.com/")
	google := f.CreateLink("Google", "http://www.google.com/")

	traynews := f.CreateTray("日报")
	traynews.Add(people)
	traynews.Add(gmw)

	trayyahoo := f.CreateTray("Yahoo!")
	trayyahoo.Add(us_yahoo)
	trayyahoo.Add(jp_yahoo)

	traysearch := f.CreateTray("检索引擎")
	traysearch.Add(trayyahoo)
	traysearch.Add(excite)
	traysearch.Add(google)

	page := f.CreatePage("LinkPage", "jizhihong")
	page.Add(traynews)
	page.Add(traysearch)

	page.Output()
}
