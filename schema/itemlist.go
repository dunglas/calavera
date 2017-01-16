package schema

type element struct {
	Id   string `json:"@id"`
	Type string `json:"@type"`
}

type context struct {
	Vocab   string  `json:"@vocab"`
	Element element `json:"element"`
}

// ItemList stores the list of documents exposed by the API.
// It is useful to generate the API entry point.
type ItemList struct {
	Id      string   `json:"@id"`
	Context context  `json:"@context"`
	Element []string `json:"element"`
}

// NewPerson initializes a new ItemList instance with some sensitive default values.
func NewItemList() *ItemList {
	return &ItemList{
		Id: "_index.jsonld",
		Context: context{
			Vocab: "http://schema.org/",
			Element: element{
				Id:   "itemListElement",
				Type: "@id",
			},
		},
	}
}
