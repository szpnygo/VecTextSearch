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
  