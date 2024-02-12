const socket = new WebSocket('ws://localhost:6969/ws');

socket.onopen = function(event) {
  console.log('WebSocket connection established.');
};

socket.onmessage = function(event) {
  console.log('Received message from server:', event.data);
  appendMessage(event.data);
};

socket.onerror = function(error) {
  console.error('WebSocket error:', error);
};

socket.onclose = function(event) {
  console.log('WebSocket connection closed:', event.code, event.reason);
};

function sendMessage() {
  const messageInput = document.getElementById('messageInput');
  const message = messageInput.value;
  socket.send(message);
  messageInput.value = ''; // Clear input field after sending
}

function appendMessage(message) {
  const messagesList = document.getElementById('messages');
  const li = document.createElement('li');
  li.textContent = message;
  messagesList.appendChild(li);
}
