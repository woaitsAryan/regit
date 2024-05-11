package initializers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	OPENAI_API_KEY string `mapstructure:"OPENAI_API_KEY"`
}

var CONFIG Config

func LoadEnv() {
	homeDir, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }

    configPath := filepath.Join(homeDir, ".config", "regit")

	viper.SetConfigName(CONFIG_FILE_NAME) 
	viper.SetConfigType("env") 
	viper.AddConfigPath(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("OpenAI key not found, recommit uses GPT 3.5 to rewrite commit messages.")
		addOpenAIKey(configPath)
		os.Exit(0)
	}

	err = viper.Unmarshal(&CONFIG)
	if err != nil {
		log.Fatal(err)
	}
	requiredKeys := getRequiredKeys(CONFIG)
	missingKeys := checkMissingKeys(requiredKeys, CONFIG)

	if len(missingKeys) > 0 {
		fmt.Println("OpenAI key not found, recommit uses GPT 3.5 to rewrite commit messages.")
		addOpenAIKey(configPath)
		os.Exit(0)
	}

}

func getRequiredKeys(config Config) []string {
	requiredKeys := []string{}
	configType := reflect.TypeOf(config)

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		tag := field.Tag.Get("mapstructure")
		if tag != "" {
			requiredKeys = append(requiredKeys, tag)
		}
	}

	return requiredKeys
}

func checkMissingKeys(requiredKeys []string, config Config) []string {
	missingKeys := []string{}

	configValue := reflect.ValueOf(config)
	for _, key := range requiredKeys {
		value := configValue.FieldByName(key).String()
		if value == "" {
			missingKeys = append(missingKeys, key)
		}
	}

	return missingKeys
}


func addOpenAIKey(configPath string) {
    var openAIKey string

    fmt.Print("Enter your OpenAI key: ")
    _, err := fmt.Scanln(&openAIKey)
    if err != nil {
        log.Fatal(err)
    }

    err = os.MkdirAll(configPath, 0755)
    if err != nil {
        log.Fatal(err)
    }

    filePath := filepath.Join(configPath, "config")
    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    if _, err := file.WriteString("OPENAI_API_KEY=" + openAIKey + "\n"); err != nil {
        log.Fatal(err)
    }

    fmt.Println("OpenAI key saved successfully. Kindly run the command again")
}