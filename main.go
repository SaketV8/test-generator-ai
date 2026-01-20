package main

import (
	"fmt"
	"time"

	"github.com/SaketV8/test-generator-ai/internal/utils"
)

func main() {
	fmt.Println("Test Generator LLM")
	// out, err := llm.ChatWithLLM("", "", false)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(out)

	spinner := utils.NewDefaultSpinner()
	// step 1:
	// build the system prompt + file_content from the project folder
	// TODO:
	// replace internal with the example-go-project folder
	spinner.Suffix = " reading the files..."
	spinner.Start()

	// intentional delay
	time.Sleep(2 * time.Second)
	files, err := utils.ReadFolder("internal")
	if err != nil {
		panic(err)
	}
	err = utils.WriteTemplateToFile(files)
	if err != nil {
		panic(err)
	}
	spinner.Stop()

	// step 2:
	// ask the llm to generate the test with structure output based on the prompt provided

	// step 3:
	// parse the llm response and save the generated test

	// step 4:
	// run the go test against the generated test

	// step 5:
	// if all passed then stop
	// else feedback loop with 3-4 retry

}
