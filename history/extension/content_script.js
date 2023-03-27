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
  
    const rootNodeId = Object.keys(mapping).find(id => !mapping[id].parent);
    processNode(rootNodeId);
    return markdownText;
  }
  