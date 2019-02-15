var zTree;
var setting = {
	view: {//表示tree的显示状态
		dblClickExpand: false,
		showLine: true,
		selectedMulti: true//表示是否允许多选
	},
	data: {
		simpleData: {
			enable:true,
			idKey: "id",
			pIdKey: "pId",
			rootPId: "0"
		}
	},
	check:{
		enable:true,
		checkStyle:"checkbox",
		chkboxType: { "Y":"s","N":"s" }
	},
};


$(document).ready(function(){
	$.ajax({
		url:"permission/ztree",
		dataType:"json",
		data:{},
		type:"post",
		async:false,
		contentType:"application/json",
		success:function (result) {
			console.log(result)
			zNodes = result
		},
		error:function (xhr) {
			alert("error")
		}
	})
	var treeobj = $.fn.zTree.init($('#tree'), setting, zNodes);//获取ztree对象，三个参数一次分别是容器(zTree 的容器 className 别忘了设置为 "ztree")、参数配置、数据源
	var zTree = $.fn.zTree.getZTreeObj("tree");//根据 treeId 获取 zTree 对象的方法,必须在初始化 zTree 以后才可以使用此方法,有了这个方法，用户不再需要自己设定全局变量来保存 zTree 初始化后得到的对象了，而且在所有回调函数中全都会返回 treeId 属性，用户可以随时使用此方法获取需要进行操作的 zTree 对象
	zTree.selectNode(zTree.getNodeByParam("id", 0));//支持同时选中多个节点,通过 zTree 对象执行此方法


});


$(function () {
	$("#roleadd").click(function () {
		$.ajax({
			url:"permission/add",
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

$(function(){
	$("#updatebtn").click(function(){
		var val=$('input:radio[name="bedStatus"]:checked').val();
		console.log(val)
		if (val == "transfer"){
			document.getElementById("higher_rolename").style.display="none";
		}else{
			document.getElementById("higher_rolename").style.display="inline";
		}

	})
})

$(function(){
	$(":radio").click(function(){
		if ($(this).val() == "transfer"){
			document.getElementById("higher_rolename").style.display="none";
		}else{
			document.getElementById("higher_rolename").style.display="inline";
		}
	});
});



$(function () {
	$("#roleupdate").click(function () {
		$.ajax({
			url:"permission/update",
			type:"POST",
			dataType: "html",
			//将form表单序列化进行传值
			data:$("#updateForm").serialize(),
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


$(function () {
	$("#deletebtn").click(function () {
		$.ajax({
			url:"permission/delete",
			type:"POST",
			dataType: "html",
			//将form表单序列化进行传值
			data:$("#deleteForm").serialize(),
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

