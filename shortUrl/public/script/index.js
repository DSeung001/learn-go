$("document").ready(function () {
    reload()

    $("form").submit(function(event) {
        postUrl();
        event.preventDefault();
    });

    // on을 통해 추가될 li에 이벤트 처리를 할 수 있음
    $("#urlUl").on("click", ".editButton", function (e) {
        e.preventDefault()
        editUrl($(this).data("id"))
    })

    // on을 통해 추가될 li에 이벤트 처리를 할 수 있음
    $("#urlUl").on("click", ".deleteButton", function (e) {
        e.preventDefault()
        deleteUrl($(this).data("id"))
    })
})

function reload () {
    getUrl()
}

function getUrl() {
    $("#urlUl").empty()

    $.ajax({
        type: "GET",
        url: "/url",
        success: function (response) {
            response.forEach(data => {
                $("#urlUl").append(`<li>
                        <a href='${data.fullUrl}' target="_blank" >${data.aliasUrl}</a>
                        <button class="editButton" data-id='${data.id}'>Edit</button>
                        <button class="deleteButton" data-id='${data.id}'>Delete</button>
                    </li>`)
            })
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR)
            console.log("API 요청 실패:", textStatus, errorThrown);
        }
    });
}

function postUrl() {
    console.log("postUrl")

    // 입력 데이터 가져오기
    let postData = {
        aliasUrl: $("#aliasUrl").val(),
        fullUrl: $("#fullUrl").val()
    }

    // POST 요청 보내기
    $.ajax({
        type: "POST",
        url: "/url",
        data: postData,
        success: function(response) {
            console.log("post success")
            reload()
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR)
            console.log("API 요청 실패:", textStatus, errorThrown);
        }
    });
}

function editUrl(id){
    $.ajax({
        type: "PATCH",
        url: "url/" + id,
        contentType: "application/json",
        data: JSON.stringify({
            aliasUrl: $("#aliasUrl").val(),
            fullUrl: $("#fullUrl").val()
        }),
        success: function(response) {
            console.log("edit success")
            reload()
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR)
            console.log("API 요청 실패:", textStatus, errorThrown);
        }
    });
}

function deleteUrl(id){
    $.ajax({
        type: "DELETE",
        url: "url/" + id,
        success: function() {
            console.log("delete success")
            reload()
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR)
            console.log("API 요청 실패:", textStatus, errorThrown);
        }
    });
}

