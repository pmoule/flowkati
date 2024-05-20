package flowkati

type InputData map[string]*InputItem

type InputItem struct {
	Label   string
	Prompt  string
	Value   any
	Options []Option
}

func NewInputData(key string) InputData {
	inputData := InputData{}
	inputData[key] = &InputItem{}

	return inputData
}

type Option struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Info  string `json:"info"`
}

func CreateOptions(values []string) []Option {
	options := []Option{}

	for _, value := range values {
		options = append(options, Option{Key: value, Value: value})
	}

	return options
}
