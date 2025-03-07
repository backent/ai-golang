package services_ai

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/backent/ai-golang/config"
	"github.com/backent/ai-golang/helpers"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AiServiceImplementation struct {
}

type Message struct {
	Role    string `json:"role"`    // "user" or "ai"
	Content string `json:"content"` // Text content
}

func NewAiServiceGemini() AiServiceInterface {
	return &AiServiceImplementation{}
}
func (implementation *AiServiceImplementation) MakeQuestionFromFile(fileURI string, amount int, chapter string, language string) (string, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GetGeminiAPIKey()))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")

	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"
	model.ResponseSchema = &genai.Schema{
		Type:     genai.TypeObject,
		Required: []string{"result"},
		Properties: map[string]*genai.Schema{
			"result": {
				Type: genai.TypeArray,
				Items: &genai.Schema{
					Type:     genai.TypeObject,
					Required: []string{"question", "options", "answer", "explanation"},
					Properties: map[string]*genai.Schema{
						"question": {
							Type: genai.TypeString,
						},
						"answer": {
							Type: genai.TypeString,
						},
						"explanation": {
							Type: genai.TypeString,
						},
						"options": {
							Type: genai.TypeArray,
							Items: &genai.Schema{
								Type: genai.TypeString,
							},
						},
					},
				},
			},
		},
	}

	session := model.StartChat()
	session.History = []*genai.Content{
		{
			Role: "user",
			Parts: []genai.Part{
				genai.FileData{URI: fileURI},
			},
		},
	}

	description := fmt.Sprintf(`
	Buatkan saya %d soal pilihan ganda dengan pertanyaan pada object 'question'
	 dan pilihannya pada object 'option'
	  dan kunci jawabannya pada object 'answer' dengan isian antara A,B,C atau D
		 dan penjelasannya pada object 'explanation'`, amount)
	if chapter != "" {
		description += ". Dan berfokus pada " + chapter
	}

	if language != "" {
		description += ". Dengan bahasa yang mudah dipahami oleh anak " + language
	}

	resp, err := session.SendMessage(ctx, genai.Text(description))
	if err != nil {
		return "", err
	}

	var textResponse string

	for _, part := range resp.Candidates[0].Content.Parts {
		text, ok := part.(genai.Text)
		if !ok {
			return "", errors.New("error while getting response ai")
		}
		textResponse = string(text)
	}
	return textResponse, nil
}

func (implementation *AiServiceImplementation) StoreFileuploadFile(file multipart.File, fileName string) (string, error) {
	ctx := context.Background()

	// Create a Gemini AI client
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GetGeminiAPIKey()))
	if err != nil {
		return "", err
	}
	defer client.Close()

	resFile, err := client.UploadFileFromPath(ctx, filepath.Join(helpers.RootDir(), "storage/pdf/"+fileName), nil)
	if err != nil {
		return "", err
	}

	// Return the uploaded file URI
	return resFile.URI, nil
}

func (implementation *AiServiceImplementation) CheckFileMaterialExists(fileURI string, text string) (string, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GetGeminiAPIKey()))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")

	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"
	model.ResponseSchema = &genai.Schema{
		Type:     genai.TypeObject,
		Required: []string{"result"},
		Properties: map[string]*genai.Schema{
			"result": {
				Type: genai.TypeBoolean,
			},
		},
	}

	session := model.StartChat()
	session.History = []*genai.Content{
		{
			Role: "user",
			Parts: []genai.Part{
				genai.FileData{URI: fileURI},
			},
		},
	}

	description := fmt.Sprintf(`
	Tolong di check apakah file tersebut mengandung materi %s`, text)

	resp, err := session.SendMessage(ctx, genai.Text(description))
	if err != nil {
		return "", err
	}

	var textResponse string

	for _, part := range resp.Candidates[0].Content.Parts {
		text, ok := part.(genai.Text)
		if !ok {
			return "", errors.New("error while getting response ai")
		}
		textResponse = string(text)
	}
	return textResponse, nil
}
