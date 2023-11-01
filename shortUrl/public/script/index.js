$(function () {

    console.log("Start")
    getUrl()

    function getUrl(){
        $.get("/url")
            .done(function (response) {
                console.log(response)
                // response.forEach(name =>{
                //     console.log(response)
                // })
                //     if (name !== ""){
                //         $("#participants").append(`<li>${name}</li>`)
                //     }
                // })
            })
            .fail(function (jqXHR, textStatus, errorThrown) {
                console.log(jqXHR)
                console.log("API 요청 실패:", textStatus, errorThrown);
            });
    }
})