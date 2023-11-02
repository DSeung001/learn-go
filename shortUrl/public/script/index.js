$("document").ready(function () {
    getUrl()

    $("form").submit(function(event) {
        postUrl();
        event.preventDefault();
    });

    // on을 통해 추가될 li에 이벤트 처리를 할 수 있음
    $("#url-ul").on("click", ".editButton", function (e) {
        e.preventDefault()
        editUrl($(this).data("id"))
    })

    // on을 통해 추가될 li에 이벤트 처리를 할 수 있음
    $("#url-ul").on("click", ".deleteButton", function (e) {
        e.preventDefault()
        deleteUrl($(this).data("id"))
    })
})

function getUrl() {
    $.get("/url")
        .done(function (response) {
            // $("#url-ul").empty()
            response.forEach(data => {
                console
                $("#url-ul").append(`<li>
                        <a href='${data.fullUrl}'>${data.aliasUrl}</a>
                        <button class="editButton" data-id='${data.id}'>Edit</button>
                        <button class="deleteButton" data-id='${data.id}'>Delete</button>
                    </li>`)
            })
        })
        .fail(function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR)
            console.log("API 요청 실패:", textStatus, errorThrown);
        });
}

function editUrl(id){
    console.log("edit")
}

function deleteUrl(id){
    console.log("delete")
}

function postUrl() {
    console.log("test")
}