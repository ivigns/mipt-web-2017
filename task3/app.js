$(document).ready(function(){
    $("#root").append("<ul></ul>");
    $("ul").append("<li><span>Сделать задание #3 по web-программированию</span><button onclick='this.parentNode.remove()'>Удалить</button></li>");
    $("#root").append("<input id='add_task_input'>");
    $("#root").append("<button id='add_task'>Добавить</button>");

    $("#add_task").click(function(){
        var txt = "Сделать задание #3 по web-программированию";
        if($("#add_task_input").val() !== "") {
            txt = $("#add_task_input").val();
        }
        var new_li = document.createElement("li");
        new_li.innerHTML = "<span>"+txt+"</span>"+"<button onclick='this.parentNode.remove()'>Удалить</button>";
        $("ul").append(new_li);
    });
});