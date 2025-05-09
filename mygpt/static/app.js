const ws = new WebSocket("ws://localhost:8081/ws");
const messagesDiv = document.getElementById("messages");
const input = document.getElementById("input");
const sendBtn = document.getElementById("send");
const saveBtn = document.getElementById("save");
const savedList = document.getElementById("saved-list");
const popup = document.getElementById("popup");
const popupKeyword = document.getElementById("popup-keyword");
const popupStory = document.getElementById("popup-story");
const closePopupBtn = document.getElementById("close-popup");

// 채팅 기록 저장
const chatHistory = [];
let currentResponse = "";
let isReceiving = false;
let lastJsonResponse = null;

// 임시 저장 스토리 목록
const savedStories = [];

ws.onopen = () => console.log("WebSocket 연결 성공");
ws.onmessage = ({ data }) => {
    // JSON 결과 메시지 분리 처리
    if (data.startsWith("__STORY_JSON__")) {
        lastJsonResponse = data.replace("__STORY_JSON__", "");
        // 번역본도 채팅방에 표시
        try {
            const obj = JSON.parse(lastJsonResponse);
            // 마지막 assistant 메시지(영문)에 번역본 추가
            const responseEl = document.getElementById("current-response");
            if (responseEl && obj.koeanStory) {
                responseEl.innerHTML += `<div style='margin-top:12px;color:#06b6d4;font-size:0.98rem'><b>한국어 번역</b><br>${obj.koeanStory}</div>`;
            }
        } catch {}
        return;
    }
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

// 저장 버튼 기능
saveBtn.onclick = async function() {
    let lastKeyword = null;
    let lastStory = null;
    let lastTitle = null;
    let lastKorean = null;
    for (let i = chatHistory.length - 1; i >= 0; i--) {
        if (chatHistory[i].role === "user") {
            lastKeyword = chatHistory[i].content;
            break;
        }
    }
    if (lastJsonResponse) {
        try {
            const obj = JSON.parse(lastJsonResponse);
            lastStory = obj.story;
            lastTitle = obj.title;
            lastKorean = obj.koeanStory;
        } catch {
            lastStory = null;
            lastTitle = null;
            lastKorean = null;
        }
    }
    if (lastKeyword && lastStory && lastTitle) {
        await fetch("/save-story", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ keyword: lastKeyword, title: lastTitle, story: lastStory, koeanStory: lastKorean })
        });
        // 저장 후 목록 새로고침
        await loadSavedStories();
    } else {
        alert("저장할 스토리가 없습니다.");
    }
};

// 저장된 스토리 목록 서버에서 불러오기
async function loadSavedStories() {
    const res = await fetch("/list-stories");
    const files = await res.json();
    savedStories.length = 0;
    for (const file of files) {
        try {
            const storyRes = await fetch("/story/" + encodeURIComponent(file));
            const storyObj = await storyRes.json();
            savedStories.push(storyObj);
        } catch {}
    }
    renderSavedList();
}

// 페이지 로드시 저장된 스토리 목록 불러오기
window.addEventListener("DOMContentLoaded", loadSavedStories);

function renderSavedList() {
    savedList.innerHTML = "";
    savedStories.forEach((item, idx) => {
        const li = document.createElement("li");
        li.textContent = item.title || item.keyword;
        li.onclick = () => showPopup(item);
        savedList.appendChild(li);
    });
}

function showPopup(item) {
    popupKeyword.textContent = item.title || "(제목 없음)";
    popupStory.innerHTML = `<div style='color:#6b7280;font-size:0.98rem;margin-bottom:10px'><b>키워드:</b> ${item.keyword}</div>` +
        `<div style='margin-bottom:16px'><b>영어 스토리</b><br>${item.story}</div>` +
        (item.korean ? `<div><b>한국어 번역</b><br>${item.korean}</div>` : "");
    popup.classList.remove("hidden");
}

closePopupBtn.onclick = function() {
    popup.classList.add("hidden");
};