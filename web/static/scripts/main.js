var apiserverURL = "http://127.0.0.1:9543/";

$(document).ready(function(){

    $('#fpeople').submit(function(event){
        event.preventDefault();

        $.ajax({
            method: 'GET',
            url: apiserverURL,
            contentType: false,
        }).then(function(data) {
            ParsePeople(data);
        });
    });


    $('#fget').submit(function(event){
        event.preventDefault();

        var formData = $(this).serialize();

        $.ajax({
            method: 'GET',
            url: apiserverURL + "person",
            data: formData,
            success: function(data) {
                ParsePerson(data);
            },
            statusCode: {
                404: function() {
                    $('.peoprow').remove();
                    Logging("not found person!!!");
                }
              }
        });
    });

    $('#fpost').submit(function(event){
        event.preventDefault();

        var person = {
            name: $('#pname').val(),
            surname: $('#psurname').val(),
        }

        var formData = JSON.stringify(person);

        $.ajax({
            method: 'POST',
            url: apiserverURL + "person",
            data: formData,
        }).then(function(data) {
            ParsePeople(data);
        });

    });

    $('#fput').submit(function(event){
        event.preventDefault();

        var person = {
            id: Number($('#putid').val()),
            name: $('#putname').val(),
            surname: $('#putsurname').val(),
        }

        var formData = JSON.stringify(person);

        $.ajax({
            method: 'PUT',
            url: apiserverURL + "person",
            data: formData,
        }).then(function(data) {
            ParsePeople(data);
        });
    });

    $('#fdelete').submit(function(event){
        event.preventDefault();

        var person = {
            id: Number($('#putid').val()),
        }

        var formData = JSON.stringify(person);

        $.ajax({
            method: 'DELETE',
            url: apiserverURL + "person",
            data: formData,
        }).then(function(data) {
            ParsePeople(data);
        });
    });

});



function ParsePerson(data){

    $('.err').remove();
    $('.peoprow').remove();

    var person = JSON.parse(data);
    
    $('#per-tab').append(
        "<tr class='peoprow'><th scope='col'>" + person.id + 
        "</th><th scope='col'>" + person.name + "</th><th scope='col'>" + 
        person.surname + "</th><tr>");
}

// фукнция парсинга пиплов в таблицу
function ParsePeople(data) {
    
    $('.peoprow').remove(); // удаляем данные из таблицы
    $('.err').remove();

    var people = JSON.parse(data);

    Object.values(people).map(item => {
        item.id
        $('#per-tab').append(
        "<tr class='peoprow'><th scope='col'>" + item.id + 
        "</th><th scope='col'>" + item.name + "</th><th scope='col'>" + 
        item.surname + "</th><tr>");

    });
} 

function Logging(log) {
    $('.err').remove();
    $('.log').append("<p class='err' style='danger'>"+log+"</p>");
}
