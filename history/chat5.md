# ChatGPT输出Markdown

## ChatGPT



## Neo

我想写一个Chrome浏览器插件，作用是获取ChatGPT的聊天记录并输出markdown。
我想补充一些细节，当我打开类似于https://chat.openai.com/chat/*****的页面的时候，直接在当前页面展示一个下载markdown的按钮。
然后使用这个网站的cookie直接请求一个api,https://chat.openai.com/backend-api/conversation/**********
请求这个api还需要在header中添加authorization，内容是Bearer+空格+cookie中的__Secure-next-auth.session-token

获得一个json返回值：
{
	"title": "Golang \"Hello, World!\"",
	"create_time": 1679893675.866467,
	"mapping": {
		"******": {
			"id": "******",
			"message": {
				"id": "******",
				"author": {
					"role": "system",
					"metadata": {}
				},
				"create_time": 1679893675.866467,
				"content": {
					"content_type": "text",
					"parts": [""]
				},
				"end_turn": true,
				"weight": 1.0,
				"metadata": {},
				"recipient": "all"
			},
			"parent": "******",
			"children": ["******"]
		},
		"******": {
			"id": "******",
			"children": ["******"]
		},
		"******": {
			"id": "******",
			"message": {
				"id": "******",
				"author": {
					"role": "user",
					"metadata": {}
				},
				"create_time": 1679893675.86882,
				"content": {
					"content_type": "text",
					"parts": ["请给我一段hello world的代码，使用golang"]
				},
				"weight": 1.0,
				"metadata": {
					"timestamp_": "absolute"
				},
				"recipient": "all"
			},
			"parent": "******",
			"children": ["******"]
		},
		"******": {
			"id": "******",
			"message": {
				"id": "******",
				"author": {
					"role": "assistant",
					"metadata": {}
				},
				"create_time": 1679893681.004823,
				"content": {
					"content_type": "text",
					"parts": ["当然，下面是使用 Golang 编写的经典的 \"Hello, World!\" 程序:\n\n```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}\n```\n\n在上面的代码中，我们首先声明了一个 `main` 函数，然后使用 `fmt` 包中的 `Println` 函数来输出 \"Hello, World!\"。运行上面的代码将会在控制台上打印出 \"Hello, World!\"。"]
				},
				"end_turn": true,
				"weight": 1.0,
				"metadata": {
					"model_slug": "text-davinci-002-render-sha",
					"finish_details": {
						"type": "stop",
						"stop": ""
					},
					"timestamp_": "absolute"
				},
				"recipient": "all"
			},
			"parent": "******",
			"children": []
		}
	},
	"moderation_results": [],
	"current_node": "******"
}


获取每段聊天记录的parts部分，然后合并输出成一个makrdown文件下载出来，每段之前要加一个标题表明是谁说的。如果role里是user，就输出Neo，如果是assistant，就输出ChatGPT

## ChatGPT

要实现这个需求，你需要编写一个Chrome浏览器插件，首先需要创建一个manifest.json文件，该文件包含了插件的基本信息和权限。然后编写一个content script，用于在指定的页面中添加下载按钮，并获取聊天记录。最后，编写一个函数，将聊天记录转换为Markdown格式并下载。

以下是实现这个功能的基本步骤：

1. 创建一个文件夹，如"ChatGPT_Markdown_Exporter"。

2. 在该文件夹中创建一个manifest.json文件，内容如下：

```json
{
  "manifest_version": 2,
  "name": "ChatGPT Markdown Exporter",
  "version": "1.0",
  "description": "Export ChatGPT conversation to Markdown",
  "permissions": ["activeTab", "cookies", "https://chat.openai.com/*"],
  "content_scripts": [
    {
      "matches": ["https://chat.openai.com/chat/*"],
      "js": ["content_script.js"]
    }
  ],
  "icons": {
    "48": "icon48.png"
  }
}
```

3. 创建一个content_script.js文件，内容如下：

```javascript
// 添加下载按钮
const downloadButton = document.createElement('button');
downloadButton.textContent = 'Download Markdown';
downloadButton.style.position = 'fixed';
downloadButton.style.top = '10px';
downloadButton.style.right = '10px';
downloadButton.style.zIndex = 1000;
downloadButton.onclick = downloadMarkdown;
document.body.appendChild(downloadButton);

async function downloadMarkdown() {
  // 获取cookie
  const cookie = await getCookie('__Secure-next-auth.session-token');
  const conversationId = window.location.href.split('/').pop();
  const apiUrl = `https://chat.openai.com/backend-api/conversation/${conversationId}`;
  const response = await fetch(apiUrl, {
    headers: {
      'Authorization': `Bearer ${cookie}`
    }
  });

  const json = await response.json();
  const markdownText = convertToMarkdown(json);
  const blob = new Blob([markdownText], {type: 'text/plain;charset=utf-8'});
  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = url;
  link.download = 'ChatGPT-Conversation.md';
  link.click();
}

function getCookie(name) {
  return new Promise((resolve, reject) => {
    chrome.cookies.get({url: 'https://chat.openai.com', name}, (cookie) => {
      if (cookie) {
        resolve(cookie.value);
      } else {
        reject(new Error('Cookie not found'));
      }
    });
  });
}

function convertToMarkdown(json) {
  const mapping = json.mapping;
  const messages = Object.values(mapping)
    .filter(node => node.message)
    .map(node => {
      const {role} = node.message.author;
      const speaker = role === 'user' ? 'Neo' : 'ChatGPT';
      const content = node.message.content.parts.join('');
      return `## ${speaker}\n\n${content}\n\n`;
    });
  return messages.join('');
}
```

4. 为插件添加一个图标（如icon48.png），并将其放在"ChatGPT_Markdown_Exporter"文件夹中。

5. 在Chrome浏览器中，打开扩展程序页面(chrome://extensions)，打开"开发者模式"，点击"加载已解压的

## Neo

Uncaught (in promise) TypeError: Cannot read properties of undefined (reading 'get')
    at content_script.js:34:20
    at new Promise (<anonymous>)
    at getCookie (content_script.js:33:10)
    at HTMLButtonElement.downloadMarkdown (content_script.js:13:24)

## ChatGPT

抱歉，我犯了一个错误。在content script中，chrome.cookies API是不可用的。我们需要使用background script来获取cookie，然后将cookie发送给content script。下面是更新后的代码：

1. 更新manifest.json文件，添加background script和新的权限：

```json
{
  "manifest_version": 2,
  "name": "ChatGPT Markdown Exporter",
  "version": "1.0",
  "description": "Export ChatGPT conversation to Markdown",
  "permissions": ["activeTab", "cookies", "https://chat.openai.com/*", "tabs"],
  "content_scripts": [
    {
      "matches": ["https://chat.openai.com/chat/*"],
      "js": ["content_script.js"]
    }
  ],
  "background": {
    "scripts": ["background_script.js"],
    "persistent": false
  },
  "icons": {
    "48": "icon48.png"
  }
}
```

2. 创建一个新的background_script.js文件：

```javascript
chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  if (request.action === 'getCookie') {
    chrome.cookies.get({url: 'https://chat.openai.com', name: '__Secure-next-auth.session-token'}, (cookie) => {
      if (cookie) {
        sendResponse({cookie: cookie.value});
      } else {
        sendResponse({error: 'Cookie not found'});
      }
    });
    return true; // 需要这一行，以便异步发送响应
  }
});
```

3. 更新content_script.js文件，使用chrome.runtime.sendMessage来获取cookie：

```javascript
// 添加下载按钮
const downloadButton = document.createElement('button');
downloadButton.textContent = 'Download Markdown';
downloadButton.style.position = 'fixed';
downloadButton.style.top = '10px';
downloadButton.style.right = '10px';
downloadButton.style.zIndex = 1000;
downloadButton.onclick = downloadMarkdown;
document.body.appendChild(downloadButton);

async function downloadMarkdown() {
  const cookieResult = await getCookie('__Secure-next-auth.session-token');
  if (cookieResult.error) {
    alert(cookieResult.error);
    return;
  }

  const cookie = cookieResult.cookie;
  const conversationId = window.location.href.split('/').pop();
  const apiUrl = `https://chat.openai.com/backend-api/conversation/${conversationId}`;
  const response = await fetch(apiUrl, {
    headers: {
      'Authorization': `Bearer ${cookie}`
    }
  });

  const json = await response.json();
  const markdownText = convertToMarkdown(json);
  const blob = new Blob([markdownText], {type: 'text/plain;charset=utf-8'});
  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = url;
  link.download = 'ChatGPT-Conversation.md';
  link.click();
}

function getCookie(name) {
  return new Promise((resolve, reject) => {
    chrome.runtime.sendMessage({action: 'getCookie', name}, (response) => {
      resolve(response);
    });
  });
}

function convertToMarkdown(json) {
  const mapping = json.mapping;
  const messages = Object.values(mapping)
    .filter(node => node.message)
    .map(node => {
      const {role} = node.message.author;
      const speaker = role === 'user' ? 'Neo' : 'ChatGPT';
      const content = node.message.content.parts.join('');
      return `## ${speaker}\n\n${content}\n\n`;
    });
  return messages.join('');
}
```

现在

## Neo

继续

## ChatGPT

现在你的Chrome插件应该能正常工作。如果你已经安装了此插件，请确保刷新扩展程序页面，以便更新插件的更改。

在这个插件中，我们使用了background script来获取cookie，并通过`chrome.runtime.sendMessage`将其发送给content script。接着，content script使用这个cookie来请求API并获取聊天记录。将聊天记录转换为Markdown格式，并创建一个下载链接。

当你打开类似于`https://chat.openai.com/chat/****`的页面时，你应该能看到一个"Download Markdown"按钮。点击该按钮，聊天记录将以Markdown格式下载到你的计算机。

如果你遇到任何问题或需要进一步的帮助，请随时告诉我。

## Neo

请在API请求中删除header里的content-type

## ChatGPT

根据你的要求，我将在API请求中删除header里的`content-type`。这里是更新后的`downloadMarkdown`函数：

```javascript
async function downloadMarkdown() {
  const cookieResult = await getCookie('__Secure-next-auth.session-token');
  if (cookieResult.error) {
    alert(cookieResult.error);
    return;
  }

  const cookie = cookieResult.cookie;
  const conversationId = window.location.href.split('/').pop();
  const apiUrl = `https://chat.openai.com/backend-api/conversation/${conversationId}`;
  const response = await fetch(apiUrl, {
    headers: {
      'Authorization': `Bearer ${cookie}`
    }
  });

  const json = await response.json();
  const markdownText = convertToMarkdown(json);
  const blob = new Blob([markdownText], {type: 'text/plain;charset=utf-8'});
  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = url;
  link.download = 'ChatGPT-Conversation.md';
  link.click();
}
```

现在，API请求将不再包含`content-type`头。在这个示例中，我们只保留了`Authorization`头。这应该不会影响插件的功能。

如果你有其他问题或需要进一步的帮助，请随时告诉我。

## Neo

好的。我想我们有一些错误要修复，API的token是从另外一个接口中获取的。https://chat.openai.com/api/auth/session

返回的结果：
{"user":{"id":"*************","name":"neo su","email":"st2udio@gmail.com","image":"************************","picture":"**************************","groups":[]},"expires":"2023-04-26T06:18:36.985Z","accessToken":"*********************"}


所以我们现在可能需要做一些修复

## ChatGPT

谢谢你提供的信息。根据新的API接口，我们需要修改插件的代码来正确获取`accessToken`。这里是更新后的`getAccessToken`和`downloadMarkdown`函数：

```javascript
async function getAccessToken() {
  const apiUrl = 'https://chat.openai.com/api/auth/session';
  try {
    const response = await fetch(apiUrl);
    if (response.ok) {
      const json = await response.json();
      return { accessToken: json.accessToken };
    } else {
      return { error: 'Failed to fetch access token.' };
    }
  } catch (error) {
    return { error: error.message };
  }
}

async function downloadMarkdown() {
  const accessTokenResult = await getAccessToken();
  if (accessTokenResult.error) {
    alert(accessTokenResult.error);
    return;
  }

  const accessToken = accessTokenResult.accessToken;
  const conversationId = window.location.href.split('/').pop();
  const apiUrl = `https://chat.openai.com/backend-api/conversation/${conversationId}`;
  const response = await fetch(apiUrl, {
    headers: {
      'Authorization': `Bearer ${accessToken}`
    }
  });

  const json = await response.json();
  const markdownText = convertToMarkdown(json);
  const blob = new Blob([markdownText], {type: 'text/plain;charset=utf-8'});
  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = url;
  link.download = 'ChatGPT-Conversation.md';
  link.click();
}
```

现在，`getAccessToken`函数将直接从`https://chat.openai.com/api/auth/session`获取`accessToken`。`downloadMarkdown`函数已经更新，它现在使用`getAccessToken`函数获取的`accessToken`，而不是从cookie中获取。

这应该修复了之前的问题。如果你有其他问题或需要进一步的帮助，请随时告诉我。

## Neo

这两个接口都需要使用当前页面的cookie，只是第二个接口的token从第一个接口中获取

## ChatGPT

了解了，我们需要为两个接口都添加cookie。这里是更新后的`getCookie`函数，以及修改后的`getAccessToken`和`downloadMarkdown`函数：

```javascript
async function getCookie(name) {
  return new Promise((resolve) => {
    chrome.runtime.sendMessage({type: 'get_cookie', details: {name: name, url: 'https://chat.openai.com'}}, (cookie) => {
      if (cookie) {
        resolve({cookie: cookie.value});
      } else {
        resolve({error: `Failed to get cookie: ${name}`});
      }
    });
  });
}

async function getAccessToken() {
  const cookieResult = await getCookie('__Secure-next-auth.session-token');
  if (cookieResult.error) {
    return { error: cookieResult.error };
  }

  const apiUrl = 'https://chat.openai.com/api/auth/session';
  try {
    const response = await fetch(apiUrl, {
      credentials: 'include',
      headers: {
        'Cookie': `__Secure-next-auth.session-token=${cookieResult.cookie}`
      }
    });
    if (response.ok) {
      const json = await response.json();
      return { accessToken: json.accessToken };
    } else {
      return { error: 'Failed to fetch access token.' };
    }
  } catch (error) {
    return { error: error.message };
  }
}

async function downloadMarkdown() {
  const accessTokenResult = await getAccessToken();
  if (accessTokenResult.error) {
    alert(accessTokenResult.error);
    return;
  }

  const accessToken = accessTokenResult.accessToken;
  const conversationId = window.location.href.split('/').pop();
  const apiUrl = `https://chat.openai.com/backend-api/conversation/${conversationId}`;

  const cookieResult = await getCookie('__Secure-next-auth.session-token');
  if (cookieResult.error) {
    alert(cookieResult.error);
    return;
  }

  const response = await fetch(apiUrl, {
    credentials: 'include',
    headers: {
      'Authorization': `Bearer ${accessToken}`,
      'Cookie': `__Secure-next-auth.session-token=${cookieResult.cookie}`
    }
  });

  const json = await response.json();
  const markdownText = convertToMarkdown(json);
  const blob = new Blob([markdownText], {type: 'text/plain;charset=utf-8'});
  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = url;
  link.download = 'ChatGPT-Conversation.md';
  link.click();
}
```

在这个版本中，我们使用`getCookie`函数获取cookie，并将其添加到`getAccessToken`和`downloadMarkdown`函数的请求头中。这样，这两个函数都会使用当前页面的cookie。

现在，插件应该能正确地获取`accessToken`并下载聊天记录的Markdown文件。如果你有其他问题或需要进一步的帮助，请随时告诉我。

## Neo

请忽视之前提到的一些错误信息，两个API请求都携带当前网站的所有cookie，然后重新输出整个content_script

## ChatGPT

好的，这里是使用当前网站所有cookie的修改后的`content_script.js`文件：

```javascript
// Add download button to the page
const downloadButton = document.createElement('button');
downloadButton.textContent = 'Download Markdown';
downloadButton.style.position = 'fixed';
downloadButton.style.bottom = '20px';
downloadButton.style.right = '20px';
downloadButton.style.zIndex = 1000;
downloadButton.addEventListener('click', downloadMarkdown);
document.body.appendChild(downloadButton);

async function getAccessToken() {
  const apiUrl = 'https://chat.openai.com/api/auth/session';
  try {
    const response = await fetch(apiUrl, { credentials: 'include' });
    if (response.ok) {
      const json = await response.json();
      return { accessToken: json.accessToken };
    } else {
      return { error: 'Failed to fetch access token.' };
    }
  } catch (error) {
    return { error: error.message };
  }
}

async function downloadMarkdown() {
  const accessTokenResult = await getAccessToken();
  if (accessTokenResult.error) {
    alert(accessTokenResult.error);
    return;
  }

  const accessToken = accessTokenResult.accessToken;
  const conversationId = window.location.href.split('/').pop();
  const apiUrl = `https://chat.openai.com/backend-api/conversation/${conversationId}`;

  const response = await fetch(apiUrl, {
    credentials: 'include',
    headers: {
      'Authorization': `Bearer ${accessToken}`
    }
  });

  const json = await response.json();
  const markdownText = convertToMarkdown(json);
  const blob = new Blob([markdownText], {type: 'text/plain;charset=utf-8'});
  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = url;
  link.download = 'ChatGPT-Conversation.md';
  link.click();
}

function convertToMarkdown(json) {
  const mapping = json.mapping;
  let markdownText = `# ${json.title}\n\n`;

  function processNode(nodeId) {
    const node = mapping[nodeId];
    if (node.message) {
      const role = node.message.author.role;
      const name = role === 'user' ? 'Neo' : 'ChatGPT';
      const text = node.message.content.parts.join('');
      markdownText += `## ${name}\n\n${text}\n\n`;
    }
    node.children.forEach(processNode);
  }

  processNode(json.current_node);
  return markdownText;
}
```

现在这个版本的`content_script.js`文件会自动包含当前网站的所有cookie。`getAccessToken`和`downloadMarkdown`函数会使用`credentials: 'include'`选项来发送请求，以便自动携带所有的cookie。

如果你有其他问题或需要进一步的帮助，请随时告诉我。

## Neo

只获取了聊天记录的最后一段，我希望获取全部的，拼接在一个markdown文件中

## ChatGPT

对不起，我的错误。这是更新后的`convertToMarkdown`函数，它会遍历所有的聊天记录并拼接在一个markdown文件中：

```javascript
function convertToMarkdown(json) {
  const mapping = json.mapping;
  let markdownText = `# ${json.title}\n\n`;

  function processNode(nodeId) {
    const node = mapping[nodeId];
    if (node.message) {
      const role = node.message.author.role;
      const name = role === 'user' ? 'Neo' : 'ChatGPT';
      const text = node.message.content.parts.join('');
      markdownText += `## ${name}\n\n${text}\n\n`;
    }
    node.children.forEach(processNode);
  }

  const rootNodeId = Object.keys(mapping).find(id => !mapping[id].parent);
  processNode(rootNodeId);
  return markdownText;
}
```

现在，这个版本的`convertToMarkdown`函数将从根节点开始处理整个聊天记录，并将其拼接成一个markdown文件。如果你有其他问题或需要进一步的帮助，请随时告诉我。

## Neo

我会把这个项目放在Github上，请帮我写一个readme

## ChatGPT

当然可以。这是一个简单的README模板，你可以根据实际情况进行修改：

```
# ChatGPT to Markdown Chrome Extension

This Chrome extension allows you to easily download ChatGPT conversation logs from OpenAI as a markdown file. The generated markdown file will include the entire conversation with clear distinctions between user and assistant

## Neo

我会把这个项目放在Github上，请帮我写一个readme，使用中文

## ChatGPT

当然可以。以下是一个简单的中文README模板，您可以根据实际情况进行修改：

```
# ChatGPT 转 Markdown Chrome 扩展

这个 Chrome 扩展允许您轻松地将 ChatGPT 与 OpenAI 的对话记录下载为 Markdown 文件。生成的 Markdown 文件将包含整个对话内容，并在用户和助手之间进行明确区分。

## 功能

- 在 ChatGPT 对话页面添加一个 "下载 Markdown" 按钮
- 将整个对话记录转换为 Markdown 格式
- 自动生成以 "Neo"（用户）和 "ChatGPT"（助手）为标题的聊天记录段落

## 安装方法

1. 克隆或下载此仓库
2. 打开 Chrome 浏览器，进入扩展管理页面（chrome://extensions/）
3. 打开右上角的 "开发者模式" 开关
4. 点击 "加载已解压的扩展程序"，选择下载或克隆的仓库文件夹

## 使用方法

1. 在已安装扩展的 Chrome 浏览器中，打开 ChatGPT 对话页面（例如：https://chat.openai.com/chat/******）
2. 点击页面上的 "下载 Markdown" 按钮
3. 保存生成的 Markdown 文件

## 注意事项

- 请确保您已登录 OpenAI 账户，且 Chrome 浏览器中存在有效的登录 Cookie
- 本扩展仅支持 ChatGPT 对话记录的下载和转换

## 贡献

欢迎提交 Pull Requests 或创建 Issues 进行反馈和建议。

## 开源许可

本项目采用 [MIT 许可证](LICENSE)。欢迎自由使用、修改和分发。

```

## Neo

我会把这个项目同时放在另外一个项目中，请补充一段简单的介绍，markdown格式，不需要包含使用说明，说明功能即可，同时表明这个插件是又ChatGPT开发的

## ChatGPT

以下是一个简短的介绍，您可以将其添加到另一个项目的文档中：

```
## ChatGPT 转 Markdown Chrome 插件

ChatGPT 转 Markdown 是一个由 ChatGPT 开发的 Chrome 插件，旨在帮助用户轻松地将 ChatGPT 与 OpenAI 的对话记录下载为 Markdown 文件。生成的 Markdown 文件将包含整个对话内容，并在用户和助手之间进行明确区分。这款插件方便用户整理和查看聊天记录，提高工作效率。

主要功能：

- 在 ChatGPT 对话页面添加一个 "下载 Markdown" 按钮
- 将整个对话记录转换为 Markdown 格式
- 自动生成以 "Neo"（用户）和 "ChatGPT"（助手）为标题的聊天记录段落

```

您可以根据需要对此介绍进行修改。

## Neo

在这段描述中添加项目路径history/extension/readme.md

## ChatGPT

以下是修改后的简短介绍，包含了项目路径：

```
## ChatGPT 转 Markdown Chrome 插件

ChatGPT 转 Markdown 是一个由 ChatGPT 开发的 Chrome 插件，旨在帮助用户轻松地将 ChatGPT 与 OpenAI 的对话记录下载为 Markdown 文件。生成的 Markdown 文件将包含整个对话内容，并在用户和助手之间进行明确区分。这款插件方便用户整理和查看聊天记录，提高工作效率。

主要功能：

- 在 ChatGPT 对话页面添加一个 "下载 Markdown" 按钮
- 将整个对话记录转换为 Markdown 格式
- 自动生成以 "Neo"（用户）和 "ChatGPT"（助手）为标题的聊天记录段落

详细说明和使用方法，请参考 [history/extension/readme.md](history/extension/readme.md) 文件。

```

现在这段介绍已经包含了项目路径，您可以将其添加到另一个项目的文档中。

## Neo

为这个插件的开发写一个commit message，并表明是你开发的

## ChatGPT

feat: 添加ChatGPT转Markdown Chrome插件

开发一个新的Chrome插件，用于将ChatGPT聊天记录转换为Markdown格式。插件在ChatGPT对话页面上添加一个"下载Markdown"按钮，方便用户获取完整聊天记录，并以Markdown文件形式下载。本插件由OpenAI的ChatGPT模型提供开发支持。

Co-authored-by: ChatGPT <chatgpt@openai.com>

