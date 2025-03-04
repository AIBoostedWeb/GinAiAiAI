## 数据库格式

数据库user-list中的存储方式类似于：
```go
package main

type User struct {
	Username string
	Password string
	Love LoveStruct
	Commit CommitStruct
}

type LoveStruct struct {
    
}
type CommitStruct struct {
	
}

```

```ts
{
  "username": "aaa",
  "password": "123",
  "love": [],
  "commit": []
}
```

例如：

```ts
{
  "username": "wanghan443",
  "password": "aaaaaa",
  "love": [
    {
      "_id": "67b16d8ede2aa1f4dcd6773c",
      "id": "google-scholar",
      "name": "Google Scholar",
      "url": "https://scholar.google.com ",
      "logo": "https://tse4-mm.cn.bing.net/th/id/OIP-C.6eJ9QJxJu2XxfjqeCVPYjgHaHa?w=179&h=184&c=7&r=0&o=5&dpr=1.5&pid=1.7",
      "description": "Google Scholar 是一个广受欢迎的免费学术搜索引擎，它索引了数百万篇学术文章，包括期刊论文、学位论文、书籍、预印本和报告等。这个平台特别适合研究人员、学生和教育工作者，他们可以利用这个工具来获取最新的研究成果和历史文献。此外，Google Scholar 还提供了引文功能，用户可以查看某篇文章被引用了多少次，以及被哪些文章引用，这对于文献综述和研究的深度分析非常有帮助。",
      "category": "research",
      "tags": [
        "学术",
        "文献",
        "搜索"
      ],
      "rating": 4.8,
      "views": 15000,
      "isPaid": false,
      "language": [
        "English"
      ],
      "accessSpeed": "快速"
    },
    {
      "_id": "67b16d8ede2aa1f4dcd6773f",
      "id": "leetcode",
      "name": "LeetCode",
      "url": "https://leetcode.com ",
      "logo": "https://leetcode.com/favicon.ico ",
      "description": "LeetCode 是一个在线编程学习平台，它提供了大量的算法题和数据结构题，供程序员练习和提高编程技能。这个平台特别适合准备技术面试和希望提高算法能力的人。LeetCode 的题目覆盖了各种难度级别，从简单到困难，用户可以根据自己的水平选择合适的题目进行练习。此外，LeetCode 还提供了一个讨论区，用户可以在其中讨论解题策略、分享知识，以及与其他程序员建立联系。",
      "category": "learning",
      "tags": [
        "算法",
        "编程",
        "练习"
      ],
      "rating": 4.9,
      "views": 17000,
      "isPaid": true,
      "language": [
        "English"
      ],
      "accessSpeed": "快速"
    },
    {
      "_id": "67b16d8ede2aa1f4dcd67740",
      "id": "bilibilibili",
      "name": "Bilibili（哔哩哔哩）",
      "url": "https://www.bili.com ",
      "logo": "https://p1.ssl.qhimg.com/t011cc0073d3c813d6b.png ",
      "description": "Bilibili不仅是一个年轻人喜欢的二次元视频平台，也是一个充满技术分享内容的学习平台。在Bilibili上，你可以找到许多编程学习视频，这些视频多由技术达人和编程专家分享，内容通俗易懂，适合新手入门。此外，Bilibili还提供了许多编程相关的直播和互动课程，可以与讲师进行实时互动，解答学习中的疑问。",
      "category": "learning",
      "tags": [
        "编程学习",
        "视频教程",
        "互动"
      ],
      "rating": 4.5,
      "views": 10000,
      "isPaid": false,
      "language": [
        "中文",
        "English"
      ],
      "accessSpeed": "快速"
    },
    {
      "_id": "67b16d8ede2aa1f4dcd6773b",
      "id": "sci-hub",
      "name": "Sci-Hub",
      "url": "https://sci-hub.se ",
      "logo": "https://sci-hub.se/favicon.ico ",
      "description": "提供免费的学术论文获取服务。",
      "category": "research",
      "tags": [
        "学术",
        "论文",
        "免费"
      ],
      "rating": 4.9,
      "views": 12000,
      "isPaid": false,
      "language": [
        "English"
      ],
      "accessSpeed": "需要科学上网"
    }
  ],
  "commit": [
    {
      "name": "是",
      "url": "https:wx.com",
      "description": "对对对",
      "category": "learning",
      "reason": "对对对"
    },
    {
      "name": "是",
      "url": "https:wx.com",
      "description": "对对对",
      "category": "learning",
      "reason": "对对对"
    },
    {
      "name": "百度",
      "url": "https:wx.com",
      "description": "读到",
      "category": "research",
      "reason": "读到"
    }
  ]
}
```

##