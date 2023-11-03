$("document").ready(function () {
    getUrl()

    // 단축하기 Form에 submit 이벤트가 발생하면 postUrl 함수를 실행
    $(".post-container").submit(function (event) {
        event.preventDefault();
        postUrl();
    });

    // on을 통해 동적으로 추가될 li에 이벤트 처리를 할 수 있음
    // 단축 URL 수정 이벤트
    $("#urlUl").on("click", ".editButton", function (e) {
        let listContainer = $(this).parent()
        let editContainer = $(this).parent().parent().find(".edit-container")

        listContainer.addClass("invisibility")
        editContainer.removeClass("invisibility")

        e.preventDefault()
    })

    // 단축 URL 수정 완료 이벤트
    $("#urlUl").on("click", ".saveButton", function (e) {
        editUrl($(this).data("id"), $(this).parent().parent().parent())
    })

    // 단축 URL 삭제 이벤트
    $("#urlUl").on("click", ".deleteButton", function (e) {
        e.preventDefault()
        deleteUrl($(this).data("id"))
    })
})

// 단축 URL 목록 조회
// 만약 http://localhost:4000가 아니면 수정 필요
function getUrl() {
    $("#urlUl").empty()

    $.ajax({
        type: "GET",
        url: "http://localhost:4000/url",
        success: function (response) {
            if (response === null) {
                $("#urlUl").append("<li>등록된 URL이 없습니다.</li>")
            } else {
                response.forEach(data => {
                    $("#urlUl").append(`
                    <li>
                        <div class="edit-container invisibility">
                            <form class="edit-form">  
                                <label for="aliasUrl">Alias Url</label>
                                <input type="text" class="aliasUrl" value="${data.aliasUrl}" required>
                                <label for="fullUrl">Full Url</label>
                                <input type="url" class="fullUrl" value="${data.fullUrl}" required>
                                <button type="button" class="saveButton" data-id='${data.id}'>Save</button>
                            </form>
                        </div>
                        
                        <div class="url-list-container">
                            <a href='${data.fullUrl}' target="_blank" >${data.aliasUrl}</a>
                            <button class="editButton" data-id='${data.id}'>Edit</button>
                            <button class="deleteButton" data-id='${data.id}'>Delete</button>
                        </div>
                        
                    </li>`)
                })
            }
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR)
            console.log("API 요청 실패:", textStatus, errorThrown);
        }
    });
}

// 단축 URL 등록
function postUrl() {
    // 입력 데이터 가져오기
    let postData = {
        aliasUrl: $("#aliasUrl").val(),
        fullUrl: $("#fullUrl").val()
    }

    // POST 요청 보내기
    $.ajax({
        type: "POST",
        url: "http://localhost:4000/url",
        data: postData,
        success: function (response) {
            $("#aliasUrl").val('')
            $("#fullUrl").val('')
            getUrl()
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR)
            console.log("API 요청 실패:", textStatus, errorThrown);
        }
    });
}

// 단축 URL 수정
function editUrl(id, listElement) {
    let editContainer = listElement.find(".edit-container")
    let data = JSON.stringify({
        aliasUrl: editContainer.find(".aliasUrl").val(),
        fullUrl: editContainer.find(".fullUrl").val()
    })

    $.ajax({
        type: "PATCH",
        url: "http://localhost:4000/url/" + id,
        contentType: "application/json",
        data: data,
        success: function () {
            getUrl()
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR)
            console.log("API 요청 실패:", textStatus, errorThrown);
        }
    });
}

// 단축 URL 삭제
function deleteUrl(id) {
    $.ajax({
        type: "DELETE",
        url: "http://localhost:4000/url/" + id,
        success: function () {
            getUrl()
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR)
            console.log("API 요청 실패:", textStatus, errorThrown);
        }
    });
}