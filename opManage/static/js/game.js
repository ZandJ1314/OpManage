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
