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


