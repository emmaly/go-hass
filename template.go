package hass

// RenderTemplate sends a template string to be rendered by the server.
func (a *Access) RenderTemplate(template string) (string, error) {
	var resultBody string
	templateBody := struct {
		Template string `json:"template"`
	}{
		Template: template,
	}

	err := a.httpPost(PathTypeAPI, "template", &templateBody, &resultBody)
	if err != nil {
		return "", err
	}

	return resultBody, nil
}
