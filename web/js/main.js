function getCustomersList() {
    searchString = $('input[name="search-string"]').val()

    $.ajax({
        url: '/api/customer/all',
        type: 'GET',
        dataType: 'json',
        data: {
            'searchString': searchString,
            'limit': 10,
            'offset': 0
        },
        success: function (products) {
            productListSuccess(products);
        }
    });
}

$(document).ready(function(){

    getTutorials();

    $("#newTutBtn").on("click", function(e){
        $("#newForm").toggle();
    });

    function getTutorials(){
        $('#tutorialsBody').html('');
        $.ajax({
            url: 'http://localhost:8080/api/customer/all',
            method: 'get',
            dataType: 'json',
            data: {
                test: 'test data'
            },
            success: function(data) {
                $(data).each(function(i, tutorial){
                    $('#tutorialsBody').append($("<tr>")
                        .append($("<td>").append(tutorial.tutorialNumber))
                        .append($("<td>").append(tutorial.title))
                        .append($("<td>").append(tutorial.author))
                        .append($("<td>").append(tutorial.type))
                        .append($("<td>").append(tutorial.id))
                        .append($("<td>").append(`
                                                <i class="far fa-edit editTut" data-tutid="`+tutorial.id+`"></i> 
                                                <i class="fas fa-trash deleteTut" data-tutid="`+tutorial.id+`"></i>
                                            `)));
                });
                loadButtons();
            }
        });
    }

    function getOneTutorial(id){
        $.ajax({
            url: 'http://localhost:3000/api/tutorials/' + id,
            method: 'get',
            dataType: 'json',
            success: function(data) {
                $($("#updateForm")[0].tutId).val(data._id);
                $($("#updateForm")[0].updateNum).val(data.tutorialNumber);
                $($("#updateForm")[0].updateTitle).val(data.title);
                $($("#updateForm")[0].updateAuthor).val(data.author);
                $($("#updateForm")[0].updateType).val(data.type);
                $("#updateForm").show();
            }
        });
    }

    $("#submitTutorial").on("click", function(e) {
        let data = {
            tutorialNumber: $($("#newForm")[0].tutNum).val(),
            title: $($("#newForm")[0].title).val(),
            author: $($("#newForm")[0].author).val(),
            type: $($("#newForm")[0].type).val()
        }

        postTutorial(data);
        $("#newForm").trigger("reset");
        $("#newForm").toggle();
        e.preventDefault();

    });


    function postTutorial(data) {
        $.ajax({
            url: 'http://localhost:3000/api/tutorials',
            method: 'POST',
            dataType: 'json',
            data: data,
            success: function(data) {
                console.log(data);
                getTutorials();
            }
        });
    }

    function loadButtons() {
        $(".editTut").click(function(e){
            getOneTutorial($($(this)[0]).data("tutid"));
            e.preventDefault();
        });

        $(".deleteTut").click(function(e){
            deleteTutorial($($(this)[0]).data("tutid"));
            e.preventDefault();
        })
    }

    function putTutorial(id, data){
        $.ajax({
            url: 'http://localhost:3000/api/tutorials/' + id,
            method: 'PUT',
            dataType: 'json',
            data: data,
            success: function(data) {
                console.log(data);
                getTutorials();
            }
        });
    }

    $("#updateTutorial").on("click", function(e) {
        let data = {
            tutorialNumber: $($("#updateForm")[0].updateNum).val(),
            title: $($("#updateForm")[0].updateTitle).val(),
            author: $($("#updateForm")[0].updateAuthor).val(),
            type: $($("#updateForm")[0].updateType).val()
        }

        putTutorial($($("#updateForm")[0].tutId).val(), data);
        $("#updateForm").trigger("reset");
        $("#updateForm").toggle();
        e.preventDefault();

    });



    function deleteTutorial(id){
        $.ajax({
            url: 'http://localhost:3000/api/tutorials/' + id,
            method: 'DELETE',
            dataType: 'json',
            success: function(data) {
                console.log(data);
                getTutorials();
            }
        });
    }

});




