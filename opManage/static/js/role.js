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
		url:"role/ztree",
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
	//序列化表单转换为json字符串
	// var formarray = $("#addForm").serializeArray()
	$("#addBtn").click(function () {
		var array = {}
		var treeObj = $.fn.zTree.getZTreeObj("tree"),
			nodes = treeObj.getCheckedNodes(true)
		var rolenodes = {}
		for (var j=0;j<nodes.length;j++){
			var id = nodes[j].id
			var rolename = nodes[j].name
			rolenodes[id] = rolename
		}
		array["roleinfo"] = rolenodes
		// alert(rolenodes)
		// alert(JSON.stringify(nodes))
		var form = $("#roleForm").serializeArray()
		for (var i=0;i<form.length;i++){
			var name = form[i].name
			var value = form[i].value
			array[name] = value
			var data = JSON.stringify(array)
		}
		// alert(data)
		$.ajax({
			url:"role/add",
			type:"POST",
			dataType: "html",
			data:data,
			async: false,
			// contentType: "application/json",
			success:function (result) {
				var arre = $.parseJSON(result)
				// alert("success")
				alert(arre.message)
			},
			error:function (xhr) {
				alert("警告：请求返回数据失败！！！")
				console.log(JSON.stringify(xhr))
			}
		})
	})
})

function UserSelet(name){
	$.get(
		"/userquery",
		{"name":name},
		function (data) {
			var html = "";
			$.each(data,function(i,obj){
				html+="<option value='"+obj.name+"'>";
				html+=obj.name;
				html+="</option>";
			});
			$("#username").html(html);
		},'json');
}

$(function () {
	$('#usertype').change(function () {
		UserSelet(this.value)
	})
})
