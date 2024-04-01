package task

import (
	"fmt"
	"os"
	"strings"
	"syj/hope/chat"
	"syj/hope/entity"
	"syj/hope/image"
	"syj/hope/service"
	"time"

	"github.com/google/uuid"
)

var prompts = []string{
	"创造一个基于当下热门电影角色的互联网梗，让它既幽默又能引起共鸣。",
	"描述一种只有程序员才会懂的幽默情境，用互联网流行语言包装。",
	"构思一个与当前社会热点事件相关的网络热句，让它带有一丝讽刺和幽默感。",
	"创建一个围绕环保主题的流行网络用语，让它既激发思考也容易传播。",
	"设计一个结合古典文学和现代网络文化的热句，展示时代的融合。",
	"想象一个由未来科技趋势激发的网络热句，让它听起来既科幻又接地气。",
	"发明一个新的网络热句，用来形容那种只有在2024年才会出现的特定情感或现象。",
	"提出一个围绕宠物和动物的可爱互联网热句，让人既心温又能笑出声。",
	"构想一个融合了多国文化元素的网络流行语，展现全球化的幽默感。",
	"创造一个能够引起全年龄层共鸣的、基于日常生活小确幸的网络热句。",
	"创作一个关于虚拟现实世界的搞笑网络热句，让它既反映现实又充满想象。",
	"想出一个以宇宙探索为主题的网络流行语，让它带有未知探索的魅力和幽默感。",
	"构思一个描述现代职场文化的网络热句，通过幽默展现工作生活的平衡。",
	"创建一个关于数字货币或区块链的热门网络热句，使其既专业又通俗易懂。",
	"设计一个围绕家庭生活小趣事的网络热句，让它温馨又富有幽默感。",
	"发明一个反映年轻人社交趋势的网络热句，体现他们的语言风格和生活态度。",
	"提出一个基于最新科技发明的网络流行语，展现科技的趣味性和影响力。",
	"构想一个与健康生活方式相关的热门热句，鼓励人们更加关注身心健康。",
	"创造一个描绘旅行和探险精神的网络热句，激发人们对未知世界的好奇心。",
	"想出一个轻松应对生活压力的网络热句，传递积极向上的生活态度。",
	"设计一个以环境变化为核心的网络热句，既有教育意义又不失幽默感。",
	"发明一个基于人工智能和机器人日益普及的幽默网络热句，展现与机器共生的未来。",
	"创作一个关于跨文化交流和理解的网络热句，体现多样性和包容性的价值。",
	"想出一个以美食探索为主题的网络流行语，既表达对美食的热爱也包含幽默元素。",
	"构思一个围绕音乐节或演唱会体验的热门热句，让它充满活力和共鸣。",
	"创建一个描述超级英雄文化影响的网络热句，既有创意又能激发人们的正义感。",
	"设计一个反映远程工作和数字游牧生活方式的网络热句，既现实又富有想象。",
	"构想一个围绕精神健康和自我关怀的网络流行语，鼓励人们关注内心世界。",
	"发明一个关于学习新技能和个人成长的激励性网络热句，既正面提升又充满智慧。",
	"创作一个关于游戏文化和电子竞技的网络热句，展现游戏世界的魅力和竞技精神。",
}

var index = 0
var imagePre = "" //网络访问的文件夹前缀，如 "https://www.allalive.cn/images/krs/"
var reNamePre = "" //本地访问的文件夹前缀路径 如 "images/krs/"

// 每隔 interval 毫秒执行一次
func StartAutoGenerateSay(interval int) {
	go func() {
		for {
			func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in safeCall:", r)
						time.Sleep(time.Second)
					}
				}()
				if index >= len(prompts) {
					index = 0
				}
				chatRes := chat.GetChatRes(prompts[index] + "只包含句子，不超过30个字，没有引号")
				index++
				strs := strings.Split(chatRes, "|")
				imageName := image.GenerateImage(chatRes + " 给这段文字画一个配图")
				updatedImageName := image.UpdateImage(imageName)
				imageName2 := image.GenerateImage(chatRes + " 给这个画一个头像")
				avatarName := image.UpdateImage(imageName2)
				os.Rename(imageName, reNamePre+imageName)
				os.Rename(imageName2, reNamePre+imageName2)
				os.Rename(avatarName, reNamePre+avatarName)
				os.Rename(updatedImageName, reNamePre+updatedImageName)
				id := uuid.New().String()
				service.InsertSay(entity.Say{
					Id:         id,
					PersonName: "A哥",
					Content:    strs[0],
					Avatar:     imagePre + avatarName,
					Image:      imagePre + updatedImageName,
				})
				time.Sleep(time.Millisecond * time.Duration(interval))
			}()
		}
	}()
}
