$("document").ready(function () {
    reload()

    $(".post-container").submit(function(event) {
        postUrl();
        event.preventDefault();
    });

    // on을 통해 추가될 li에 이벤트 처리를 할 수 있음
    $("#urlUl").on("click", ".editButton", function (e) {
        console.log("실행")
        console.log($(this).parent().find(".edit-container"))

        let listContainer = $(this).parent()
        let editContainer = $(this).parent().parent().find(".edit-container")

        listContainer.addClass("invisibility")
        editContainer.removeClass("invisibility")

        e.preventDefault()
    })

    $("#urlUl").on("click", ".saveButton", function (e) {
        editUrl($(this).data("id"), $(this).parent().parent().parent())
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
                $("#urlUl").append(`
                    <li>
                        <div class="edit-container invisibility">
                            <form class="edit-form">  
                                <label for="aliasUrl">Alias Url</label>
                                <input type="text" class="aliasUrl" value="${data.aliasUrl}" required>
                                <label for="fullUrl">Full Url</label>
                                <input type="url" class="fullUrl" value="${data.fullUrl}" required>
                                <button type="button" class="saveButton" data-id='${data.id}'>save</button>
                            </form>
                        </div>
                        
                        <div class="url-list-container">
                            <a href='${data.fullUrl}' target="_blank" >${data.aliasUrl}</a>
                            <button class="editButton" data-id='${data.id}'>Edit</button>
                            <button class="deleteButton" data-id='${data.id}'>Delete</button>
                        </div>
                        
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

function editUrl(id, listElement){
    let editContainer = listElement.find(".edit-container")
    let data = JSON.stringify({
        aliasUrl: editContainer.find(".aliasUrl").val(),
        fullUrl: editContainer.find(".fullUrl").val()
    })
    console.log(data)

    $.ajax({
        type: "PATCH",
        url: "url/" + id,
        contentType: "application/json",
        data: data,
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