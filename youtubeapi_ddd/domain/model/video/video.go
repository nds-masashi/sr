// Youtube の情報を保持するパッケージ
package video

/*
Youtube の情報を保持するモデル
*/
type Model struct {
	url   string
	title string
}

func New(
	URL string,
	Title string,
) Model {
	return Model{
		url:   URL,
		title: Title,
	}
}

func (m Model) GetURL() string {
	return m.url
}

func (m Model) GetTitle() string {
	return m.title
}
