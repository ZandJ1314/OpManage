$(document).ready(function () {
    var trigger = $('.hamburger'),
        overlay = $('.overlay'),
        isClosed = false;

    trigger.click(function () {
        hamburger_cross();
    });

    function hamburger_cross() {

        if (isClosed == true) {
            overlay.hide();
            trigger.removeClass('is-open');
            trigger.addClass('is-closed');
            isClosed = false;
        } else {
            overlay.show();
            trigger.removeClass('is-closed');
            trigger.addClass('is-open');
            isClosed = true;
        }
    }

    $('[data-toggle="offcanvas"]').click(function () {
        $('#wrapper').toggleClass('toggled');
    });
});

// function getname() {
//     var name = $(".name-a").val()
//     $("#showname").html(name)
// }

//在模态窗口使用 window.dialogArguments 表示父窗口对象
$("#getname").on("shown.bs.modal",function (event) {
    var name = window.dialogArguments.document.getElementById("getname").value;
    $("#showname").html(name)
})





// document.getElementById("getname").click(function () {
//     var name = $(".name-a").val()
//     console.log(name)
//     $("#showname").html(name)
// });


function checknumber(){
    var phone= $("#heigherrole").val();
    if(/^1[34578]\d{9}$/.test(phone)){
        return true;
    }else{
        alert('电话号码格式不正确，请重新填写!');
        $("#heigherrole").val('');
        return false;
    }
}


function checkemail() {
    var email = $("#rolelevel").val();
    if(/^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/.test(email)){
        return true
    }else{
        alert("邮箱格式书写不正确，请重新填写!");
        $("#rolelevel").val('');
        // $("#showemail").html("<font color='red'>手机号码格式不正确！请重新输入！</font>")
        return false;
    }
}

function checkimg() {
    var fileType = $('#headphoto').val();
    console.log(fileType)
    var filestart = fileType.lastIndexOf(".")
    var length = fileType.length
    var finalType = fileType.substring(filestart,length);
    console.log(finalType);
    if (finalType != ".jpg" && finalType != ".png" && finalType != ".jpeg"){
        alert("请输入正确格式的图片,要以jpg/jpeg/png结尾");
        return false
    }
    return true
}


// function checkfun() {
//     if (checknumber() && checkemail() && checkimg()){
//         return true;
//     }
//     alert("请检查您的填写格式");
//     return false;
// }

