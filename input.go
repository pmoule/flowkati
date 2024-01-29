package flowkati

type InputData map[string]*InputItem

type InputItem struct {
	Label   string
	Prompt  string
	Value   any
	Options []string
}

func NewInputData(key string) InputData {
	inputData := InputData{}
	inputData[key] = &InputItem{}

	return inputData
}
