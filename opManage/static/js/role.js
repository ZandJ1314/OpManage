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

