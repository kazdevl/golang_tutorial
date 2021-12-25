package fileRead

type Content struct {
	Page string `json:"page"`
}

type SliceContent struct {
	Pages []string `json:"pages"`
}

func ReadTextFileContent() []string {
	return nil
}

func ReadJsonFileContentWithContents() []Content {
	return nil
}

func ReadJsonFileContentWithSliceContent() []string {
	return nil
}
