package service

import (
	"github.com/sashabaranov/go-openai"
	"go-gin/internal/config"
	"go-gin/internal/model"
	"go-gin/pkg/util"
)

type ReqRole string

const (
	PROMPT   ReqRole = "promptForPrompt"
	GENERATE ReqRole = "generate"
	CHECK    ReqRole = "check"
)

const (
	promptForPrompt string = "# 任务描述 \n" +
		"## 角色设定 \n" +
		"你是一名优秀的项目经理，你要帮你的用户梳理需求（见\\#\\#用户需求），生成ai的prompt指导另一个本地的ai编写脚本。\n" +
		"## 任务步骤 \n" +
		"1. 思考需要用那些具体步骤 \n" +
		"2. 指出哪里有易错点 \n" +
		"3. 记得备注运行的具体环境是linux的bash \n" +
		"4. remember to ask ai to only generate **bash shell scirpt**" +
		"5.整体检查 \n" +
		"## 输出格式\n" +
		"输出为文本\n" +
		"参考markdown格式\n" +
		"## 用户需求: \n"

	promptForQR string = "# 任务描述 \n" +
		"## 角色设定 \n" +
		"你是程序员，您需要根据用户需求写脚本\n" +
		"## 任务步骤 \n" +
		"1. 阅读用户需求 \n" +
		"2. 记得运行的具体环境是linux的bash \n" +
		"3. remember to only generate **bash shell scirpt**" +
		"4.整体检查 \n" +
		"5.看看有没有格式错误或者多余字符，如```bash" +
		"## 输出格式\n" +
		"输出为bash script only\n" +
		"## 用户需求: \n"
)

func AskAI(cfgForAi *config.AIConfig, msgList *[]openai.ChatCompletionMessage, role ReqRole) (string, error) {
	var response string
	var err error

	client := util.GenOwnClient(cfgForAi)

	switch role {
	case PROMPT:

		lastContentIdx := len(*msgList) - 2
		msgListLast := (*msgList)[lastContentIdx].Content

		(*msgList)[lastContentIdx] = util.GenMessage(util.USER, promptForPrompt+msgListLast)
		response, err = client.RequestWithSdkMessage(msgList)
		if err != nil {
			return "", err
		}
	case GENERATE:
		(*msgList)[0].Content = promptForQR + (*msgList)[0].Content
		response, err = client.RequestWithSdkMessage(msgList)
		response = response + " "
	case CHECK:
		(*msgList)[0].Content = promptForQR + (*msgList)[0].Content
		response, err = client.RequestWithSdkMessage(msgList)
		response = response + " "

	}
	return response, nil
}

func AskAiWithMessage(cfgForAi *config.AIConfig, _msgList *[]model.Message, role ReqRole) (string, error) {
	msgList := util.MessageTransformer(_msgList)

	return AskAI(cfgForAi, msgList, role)

}

func AskAIWithStr(cfgForAi *config.AIConfig, _message string, role ReqRole) (string, error) {
	message :=
		&([]openai.ChatCompletionMessage{
			util.GenMessage(util.USER, _message),
		})
	return AskAI(cfgForAi, message, role)

}
