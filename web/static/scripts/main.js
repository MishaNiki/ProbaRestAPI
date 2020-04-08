var apiserverURL = "http://127.0.0.1:9543/";

var request = new XMLHttpRequest();

request.open("GET", apiserverURL)
//request.setRequestHeader('Content-type', 'application/json; charset=utf-8');

$(document).ready(function(){

    $("#fpeople").submit(function() {
        alert("people");
    });

    $("#fget").submit(function() {
        alert("get");
    });

    $("#fpost").submit(function() {
        alert("post");
    });

    $("#fput").submit(function() {
        alert("put");        
    });

    $("#fdelete").submit(function() {
        alert("delete");
    });

});