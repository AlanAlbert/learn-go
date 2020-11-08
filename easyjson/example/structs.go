package easyjson_example

type Person struct {
	Name string `json:"name"`
	Age uint8 `json:"age"`
}

type Student struct {
	Person
	Class string `json:"class"`
	Score float32 `json:"score"`
}

type Teacher struct {
	Person
	Class string `json:"class"`
	Title string `json:"title"`
}
