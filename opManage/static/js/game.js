$(function(){
    $("#addbtn").click(function(){
        var val=$('input:radio[name="bedStatus"]:checked').val();
        console.log(val)
        if (val == "transfer"){
            document.getElementById("addgametype").style.display="none";
            document.getElementById("gametype").style.display="inline";
        }else{
            document.getElementById("addgametype").style.display="inline";
            document.getElementById("gametype").style.display="none";
        }

    })
})
$(function(){
    $(":radio").click(function(){
        if ($(this).val() == "transfer"){
            document.getElementById("addgametype").style.display="none";
            document.getElementById("gametype").style.display="inline";
        }else{
            document.getElementById("addgametype").style.display="inline";
            document.getElementById("gametype").style.display="none";
        }
    });
});

$(function () {
    $("#gameaddbtn").click(function () {
        $.ajax({
            url:"game/add",
            type:"POST",
            dataType: "html",
            //将form表单序列化进行传值
            data:$("#addForm").serialize(),
            async: false,
            // contentType: "application/json",
            success:function (result) {
                var arre = $.parseJSON(result)
                alert(arre.message)
            },
            error:function (xhr) {
                alert("警告：请求返回数据失败！！！")
                console.log(JSON.stringify(xhr))
            }
        })
    })
})



