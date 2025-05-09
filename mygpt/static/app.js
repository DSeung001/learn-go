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
        try {
            const obj = JSON.parse(lastJsonResponse);
            // 채팅방에 제목, 영문스토리, 한글스토리, 영문요약, 한글요약 모두 표시
            const responseEl = document.getElementById("current-response");
            if (responseEl) {
                let html = "";
                if (obj.title) html += `<div style='font-weight:bold;font-size:1.1rem;color:#4f46e5;margin-bottom:6px'>${obj.title}</div>`;
                if (obj.englishStory) html += `<div style='margin-bottom:10px'><b>영어 본문</b><br>${obj.englishStory}</div>`;
                if (obj.koreanStory) html += `<div style='color:#06b6d4;margin-bottom:10px'><b>한글 본문</b><br>${obj.koreanStory}</div>`;
                if (obj.englishSummary) html += `<div style='color:#64748b;margin-bottom:6px'><b>영문 요약</b><br>${obj.englishSummary}</div>`;
                if (obj.koreanSummary) html += `<div style='color:#64748b'><b>한글 요약</b><br>${obj.koreanSummary}</div>`;
                responseEl.innerHTML = html;
            }
        } catch {}
        return;
    }
    // 메시지 완료 신호 확인
    if (data.includes("__MESSAGE_COMPLETE__")) {
        isReceiving = false;
        if (chatHistory.length > 0 && chatHistory[chatHistory.length - 1].role === "user") {
            chatHistory.push({ role: "assistant", content: currentResponse });
        }
        const responseEl = document.getElementById("current-response");
        if (responseEl) {
            responseEl.removeAttribute("id");
        }
        return;
    }
    if (!isReceiving) {
        isReceiving = true;
        currentResponse = data;
        const responseEl = document.createElement("div");
        responseEl.className = "message assistant";
        responseEl.id = "current-response";
        responseEl.textContent = currentResponse;
        messagesDiv.appendChild(responseEl);
    } else {
        currentResponse += data;
        const responseEl = document.getElementById("current-response");
        responseEl.textContent = currentResponse;
    }
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
    let lastEnglishStory = null;
    let lastTitle = null;
    let lastKoreanStory = null;
    let lastEnglishSummary = null;
    let lastKoreanSummary = null;
    for (let i = chatHistory.length - 1; i >= 0; i--) {
        if (chatHistory[i].role === "user") {
            lastKeyword = chatHistory[i].content;
            break;
        }
    }
    if (lastJsonResponse) {
        try {
            const obj = JSON.parse(lastJsonResponse);
            lastEnglishStory = obj.englishStory;
            // **로 감싸진 부분을 title로 추출
            const match = lastEnglishStory && lastEnglishStory.match(/\*\*(.*?)\*\*/);
            lastTitle = match ? match[1].trim() : "Story";
            lastKoreanStory = obj.koreanStory;
            lastEnglishSummary = obj.englishSummary;
            lastKoreanSummary = obj.koreanSummary;
        } catch {
            lastEnglishStory = null;
            lastTitle = null;
            lastKoreanStory = null;
            lastEnglishSummary = null;
            lastKoreanSummary = null;
        }
    }
    if (lastKeyword && lastEnglishStory && lastTitle) {
        await fetch("/save-story", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ keyword: lastKeyword, title: lastTitle, englishStory: lastEnglishStory, koreanStory: lastKoreanStory, englishSummary: lastEnglishSummary, koreanSummary: lastKoreanSummary })
        });
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
    popupStory.innerHTML =
        (item.title ? `<div style='font-weight:bold;font-size:1.1rem;color:#4f46e5;margin-bottom:6px'><b>Title</b>: ${item.title}</div><br/>` : "") +
        (item.englishStory ? `<div style='margin-bottom:10px'><b>영어 본문</b><br>${item.englishStory}</div>` : "") +
        (item.koreanStory ? `<div style='color:#06b6d4;margin-bottom:10px'><b>한글 본문</b><br>${item.koreanStory}</div>` : "") +
        (item.englishSummary ? `<div style='color:#64748b;margin-bottom:6px'><b>영문 요약</b><br>${item.englishSummary}</div>` : "") +
        (item.koreanSummary ? `<div style='color:#64748b'><b>한글 요약</b><br>${item.koreanSummary}</div>` : "");
    popup.classList.remove("hidden");
}

closePopupBtn.onclick = function() {
    popup.classList.add("hidden");
};