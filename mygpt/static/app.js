const ws = new WebSocket("ws://localhost:8080/ws");
const messagesDiv = document.getElementById("messages");
const input = document.getElementById("input");
const sendBtn = document.getElementById("send");

// 채팅 기록 저장
const chatHistory = [];
let currentResponse = "";
let isReceiving = false;

ws.onopen = () => console.log("WebSocket 연결 성공");
ws.onmessage = ({ data }) => {
    // 메시지 완료 신호 확인
    if (data.includes("__MESSAGE_COMPLETE__")) {
        isReceiving = false;
        
        // 완료된 대화 저장
        if (chatHistory.length > 0 && chatHistory[chatHistory.length - 1].role === "user") {
            chatHistory.push({ role: "assistant", content: currentResponse });
        }
        
        // ID 제거
        const responseEl = document.getElementById("current-response");
        if (responseEl) {
            responseEl.removeAttribute("id");
        }
        return;
    }
    
    // 응답 처리
    if (!isReceiving) {
        // 새 응답 시작
        isReceiving = true;
        currentResponse = data;
        
        // 응답을 위한 새 메시지 엘리먼트 생성
        const responseEl = document.createElement("div");
        responseEl.className = "message assistant";
        responseEl.id = "current-response";
        responseEl.textContent = currentResponse;
        messagesDiv.appendChild(responseEl);
    } else {
        // 이미 받고 있는 응답에 텍스트 추가
        currentResponse += data;
        const responseEl = document.getElementById("current-response");
        responseEl.textContent = currentResponse;
    }
    
    // 스크롤 자동 조정
    messagesDiv.scrollTop = messagesDiv.scrollHeight;
};

sendBtn.onclick = sendMessage;
input.addEventListener("keypress", e => {
    if (e.key === "Enter") sendMessage();
});

function sendMessage() {
    if (!input.value) return;
    
    // 사용자 메시지 표시
    const userMessage = input.value;
    const userMessageEl = document.createElement("div");
    userMessageEl.className = "message user";
    userMessageEl.textContent = userMessage;
    messagesDiv.appendChild(userMessageEl);
    
    // 대화 기록에 사용자 메시지 추가
    chatHistory.push({ role: "user", content: userMessage });
    
    // WebSocket으로 메시지 전송
    ws.send(userMessage);
    
    // 입력창 초기화
    input.value = "";
    
    // 응답 시작 상태 리셋
    isReceiving = false;
}
