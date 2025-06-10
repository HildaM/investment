package utility

import (
	"context"
	"investment/server/conf"

	"github.com/cloudwego/eino-ext/components/model/gemini"
	"google.golang.org/api/option"

	"github.com/cloudwego/eino-ext/components/model/openai"
	aclopenai "github.com/cloudwego/eino-ext/libs/acl/openai"
	genai "github.com/google/generative-ai-go/genai"
)

// 硅基流动的ds模型
func GetSFDeepSeekChatModel() *openai.ChatModel {
	result, err := openai.NewChatModel(context.Background(), &openai.ChatModelConfig{
		APIKey:     conf.Get().System.SFAPIkey,
		Model:      "deepseek-ai/DeepSeek-R1",
		BaseURL:    "https://api.siliconflow.cn/v1",
		HTTPClient: NewDebugHTTPClient(),
	})
	if err != nil {
		panic(err)
	}
	return result
}

// 腾讯的ds模型
func GetTTDeepSeekChatModel() *openai.ChatModel {
	result, err := openai.NewChatModel(context.Background(), &openai.ChatModelConfig{
		APIKey:     conf.Get().System.TTAPIkey,
		Model:      "deepseek-r1",
		BaseURL:    "https://api.lkeap.cloud.tencent.com/v1",
		HTTPClient: NewDebugHTTPClient(),
	})
	if err != nil {
		panic(err)
	}
	return result
}

// 阿里千问模型
func GetQwenChatModel() *openai.ChatModel {
	temperature := Of[float32](1.0)
	result, err := openai.NewChatModel(context.Background(), &openai.ChatModelConfig{
		APIKey:      conf.Get().System.QWENAPIkey,
		Model:       "qwen-plus-latest",
		BaseURL:     "https://dashscope.aliyuncs.com/compatible-mode/v1",
		HTTPClient:  NewDebugHTTPClient(),
		Temperature: temperature,
	})
	if err != nil {
		panic(err)
	}
	return result
}

// gemini
func GetGeminiChatModel() *gemini.ChatModel {
	client, err := genai.NewClient(
		context.Background(),
		option.WithAPIKey(conf.Get().System.GeminiAPIkey),
		option.WithHTTPClient(NewDebugHTTPClient(conf.Get().System.GeminiAPIkey)),
	)
	if err != nil {
		panic(err)
	}
	chatmodel, err := gemini.NewChatModel(context.Background(), &gemini.Config{
		Client: client,
		Model:  "gemini-2.0-flash",
	})
	if err != nil {
		panic(err)
	}
	return chatmodel
}

const (
	EmbeddingModelTextEmbeddingAda002 string = "text-embedding-ada-002"
	EmbeddingModelTextEmbedding3Small string = "text-embedding-3-small"
	EmbeddingModelTextEmbedding3Large string = "text-embedding-3-large"
)

func GetDeepSeekEmbedder() (result *aclopenai.EmbeddingClient) {
	i := Of(32)
	result, err := aclopenai.NewEmbeddingClient(context.Background(), &aclopenai.EmbeddingConfig{
		APIKey: conf.Get().System.SFAPIkey,
		//Model:  "BAAI/bge-large-zh-v1.5",
		Model:      "BAAI/bge-large-en-v1.5",
		BaseURL:    "https://api.siliconflow.cn/v1",
		HTTPClient: NewDebugHTTPClient(),
		Dimensions: i,
	})
	if err != nil {
		panic(err)
	}
	return
}

func GetTTDeepSeekEmbedder() (result *aclopenai.EmbeddingClient) {
	i := Of(32)
	result, err := aclopenai.NewEmbeddingClient(context.Background(), &aclopenai.EmbeddingConfig{
		APIKey:     conf.Get().System.TTAPIkey,
		Model:      EmbeddingModelTextEmbeddingAda002,
		BaseURL:    "https://api.lkeap.cloud.tencent.com/v1",
		HTTPClient: NewDebugHTTPClient(),
		Dimensions: i,
	})
	if err != nil {
		panic(err)
	}
	return
}

func GetGeminiEmbedder() (result *aclopenai.EmbeddingClient) {
	result, err := aclopenai.NewEmbeddingClient(context.Background(), &aclopenai.EmbeddingConfig{
		APIKey:     conf.Get().System.GeminiAPIkey,
		Model:      "text-embedding-004",
		BaseURL:    "https://generativelanguage.googleapis.com/v1beta/openai/",
		HTTPClient: NewDebugHTTPClient(),
		User:       new(string),
	})
	if err != nil {
		panic(err)
	}
	return
}

// 阿里千问嵌入
func GetQwenEmbedder() (result *aclopenai.EmbeddingClient) {
	result, err := aclopenai.NewEmbeddingClient(context.Background(), &aclopenai.EmbeddingConfig{
		APIKey:     conf.Get().System.QWENAPIkey,
		Model:      "text-embedding-v3",
		BaseURL:    "https://dashscope.aliyuncs.com/compatible-mode/v1",
		HTTPClient: NewDebugHTTPClient(),
		User:       new(string),
	})
	if err != nil {
		panic(err)
	}
	return result
}
