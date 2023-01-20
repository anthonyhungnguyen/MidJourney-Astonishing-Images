package main

type Data struct {
	Props struct {
		PageProps struct {
			Event struct {
				Height uint16 `json:"height"`
				Width  uint16 `json:"width"`
			} `json:"event"`
			Jobs []struct {
				ImagePaths []string `json:"image_paths"`
			} `json:"jobs"`
			Prompt string `json:"prompt"`
			Type   string `json:"type"`
		} `json:"pageProps"`
	} `json:"props"`
}

type ParsedInfo struct {
	Height     uint16
	Width      uint16
	ImagePaths []string
	Prompt     string
	Type       string
}
