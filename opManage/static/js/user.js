$(function(){
    $("#addbtn").click(function(){
        var val=$('input:radio[name="bedStatus"]:checked').val();
        console.log(val)
        if (val == "transfer"){
            document.getElementById("addusertype").style.display="none";
            document.getElementById("usertype").style.display="inline";
        }else{
            document.getElementById("addusertype").style.display="inline";
            document.getElementById("usertype").style.display="none";
        }

    })
})
$(function(){
    $(":radio").click(function(){
        if ($(this).val() == "transfer"){
            document.getElementById("addusertype").style.display="none";
            document.getElementById("usertype").style.display="inline";
        }else{
            document.getElementById("addusertype").style.display="inline";
            document.getElementById("usertype").style.display="none";
        }
    });
});

$(function () {
    $("#useraddbtn").click(function () {
        $.ajax({
            url:"user/add",
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

// var trlens = $('tr #detailsBtn').length //获取元素的长度个数
// $(function () {
//     $('tr #deleteBtn').click(function () {
//         confirm("您确定要删除吗？")
//     })
// })


function detail(name) {
    var data = {"name":name}
    $.ajax({
        url:"user/detail",
        type:"POST",
        dataType: "html",
        //将form表单序列化进行传值
        data:JSON.stringify(data),
        async: false,
        // contentType: "application/json",
        success:function (result) {
            var arre = $.parseJSON(result);
            var lens = arre.length
            h4 = document.getElementById("detailname");
            h4.innerText = name;
            if (lens === undefined){
                p = document.getElementById("detailgame")
                p.innerText = arre.gamename
            }else{
                for (var i=0;i<lens;i++){
                    $('#detailgame').append('<li>' + arre[i].gamename + '</li>')
                }
            }
            // location.reload()

        },
        error:function (xhr) {
            alert("警告：请求返回数据失败！！！")
            console.log(JSON.stringify(xhr))
        }
    })
}

$(function () {
    $('#usertypename').change(function () {
        var name = $('#usertypename').val()
        $.ajax({
            url:"user/judge",
            type: "get",
            data:{name:name},
            dataType: "json",
            contentType: "application/json",
            success:function (result) {
                if (result.name === "false"){
                    alert("该组名已经存在了！！");
                    $('#usertypename').val('')
                    return false
                }else{
                    return true
                }
            },
            error:function (xhr) {
                alert(xhr)
            }
        })
    })
})

$(function () {
    $('#username').change(function () {
        var name = $('#username').val()
        $.ajax({
            url:"user/judgeuser",
            type: "get",
            data:{name:name},
            dataType: "json",
            contentType: "application/json",
            success:function (result) {
                if (result.name === "false"){
                    alert("该管理员已经存在了！！");
                    $('#username').val('')
                    return false
                }else{
                    return true
                }
            },
            error:function (xhr) {
                alert(xhr)
            }
        })
    })
})
