function bootstrapValidation(form, data) {

    "use strict";
    var errors = {
			fieldErrors:data.fieldErrors,
			errors:data.errors
	};
    // Clear existing errors on submit
    form.find("div.error").removeClass("error");
    form.find("span.s2_help_inline").remove();
    form.find("div.s2_validation_errors").remove();

    //Handle non field errors
    if (errors.errors && errors.errors.length > 0) {
        var errorDiv = $("<div class='alert alert-error s2_validation_errors'></div>");
        form.prepend(errorDiv);
        $.each(errors.errors, function(index, value) {
            errorDiv.append('<p>' + value + '</p>\n');
        });
    }

    //Handle field errors
    if (errors.fieldErrors) {
        $.each(errors.fieldErrors, function(index, value) {
            var element = form.find(":input[name=\"" + index + "\"]"), controlGroup, controls;
            if (element && element.length > 0) {

            // select first element
                element = $(element[0]);
            controlGroup = element.closest("div.control-group");
            controlGroup.addClass('error');
            controls = controlGroup.find("div.controls");
            if (controls) {
                controls.append("<span class='help-inline s2_help_inline'>" + value[0] + "</span>");
            }
            }
        });
    }
}
function bootstrapConfirm(message,callback){
	var div = "<div style='width:300px;margin:0 0 0 0;left:40%;top:30%;' class='modal hide' id='confirmModal'><div class='modal-body'><span><div style='text-align:center; min-height:30px; _height:30px; padding-top:20px'>"+message+"</div></span></div><div style='padding:7px;' class='modal-footer'><a href='javascript:void(0);' id='yes' class='btn'>确定</a><a href='javascript:void(0);' class='btn' id='no'>取消</a></div></div>";
	$("body").append(div);
	$("#confirmModal").modal({
		backdrop:'static'
	});
	$("#no").click(function(){
		$('#confirmModal').modal('hide');
		$('#confirmModal').remove();
	});
	$("#yes").click(function(){
		$('#confirmModal').modal('hide');
		$('#confirmModal').remove();
		callback();
	});
}
function bootstrapAlert(message){
	var div = "<div style='width:300px;margin:0 0 0 0;left:40%;top:30%;' class='modal hide' id='confirmModal'><div class='modal-body'><span><div style='text-align:center; min-height:30px; _height:30px; padding-top:20px'>"+message+"</div></span></div><div style='padding:7px;' class='modal-footer'><a href='javascript:void(0);' class='btn' id='no'>确定</a></div></div>";
	$("body").append(div);
	$("#confirmModal").modal({
		backdrop:'static'
	});
	$("#no").click(function(){
		$('#confirmModal').modal('hide');
		$('#confirmModal').remove();
	});
}

function createOperationalTable(data,title,colnum){
	if(title.length != colnum.length){
		return "";
	}else{
		var table = "<table class='table table-bordered  table-striped table-hover' ><thead><tr>";
		for (var i=0;i<title.length;i++) {
			table = table+"<th style='text-align:center'>"+title[i]+"</th>";
		}
		table += "</tr></thead><tbody>";
		for(var i = 0;i < data.length;i++){
			table += "<tr>";
			var res = data[i];
			for (var j=0;j<colnum.length;j++) {
				if(typeof(colnum[j]) == "object"){
					table += "<td style='text-align:center'>"+colnum[j].value+"</td>";
				}else{
					var value = typeof(res[colnum[j]]) == "undefined"?"":res[colnum[j]];
					table += "<td name="+colnum[j]+" style='text-align:center'>"+value+"</td>";
				}
				
			}
			table += "</tr>";
		}
		if(data.length == 0) {
			table += "<tr> <td colspan=" + title.length + ">Empty !</td></tr>";
		}
		table += "</tbody></table>";
		return table;
	}
}

function in_array(array,e){  
	for(var i=0;i<array.length;i++){  
		if(array[i] == e)  
			return true;  
		}  
		return false;  
	};  

function createCommonTable(data){
	var table;
	var attrs = new Array();
	for(var i=0;i<data.length;i++){
		for(var attr in data[i]){
			//alert(attr);
			if(!in_array(attrs,attr)){
				attrs.push(attr);
			}
		}
	}
	table = "<table class='table table-bordered  table-striped table-hover' ><thead><tr class='info'>";
	for(var i=0;i<attrs.length;i++){
		table = table+"<th>"+attrs[i]+"</th>";
	}
//	for(var attr in attrs){
//		table = table+"<th>"+attr+"</th>";
//	}
	table += "</tr></thead><tbody>";
	for(var i = 0;i < data.length;i++){
		table += "<tr>";
		var res = data[i];
		for(var j = 0;j < attrs.length;j++){
			var value = typeof(res[attrs[j]]) == "undefined"?"":res[attrs[j]];
			table += "<td>"+value+"</td>";
		}
		table += "</tr>";
	}
	table += "</tbody></table>";
	return table;
}
function createTable(data,title,colnum){
	if(title.length != colnum.length){
		return "";
	}else{
		var table = "<table class='table table-bordered  table-striped table-hover' ><thead><tr>";
		for (var i=0;i<title.length;i++) {
			table = table+"<th>"+title[i]+"</th>";
		}
		table += "</tr></thead><tbody>";
		for(var i = 0;i < data.length;i++){
			table += "<tr>";
			var res = data[i];
			for (var j=0;j<colnum.length;j++) {
				var value = typeof(res[colnum[j]]) == "undefined"?"":res[colnum[j]];
				table += "<td>"+value+"</td>";
			}
			table += "</tr>";
		}
		
		if(data.length == 0) {
			table += "<tr> <td colspan=" + title.length + ">Empty !</td></tr>";
		}
		
		table += "</tbody></table>";
		return table;
	}
	
}

function createSelectTable(data,title,colnum){
	if(title.length != colnum.length){
		return "";
	}else{
		var table = "<table class='table table-bordered  table-striped table-hover'  style='word-wrap:break-word;word-break:break-all;white-space: pre-wrap;'><thead><tr>";
		for (var i=0;i<title.length;i++) {
			table = table+"<th>"+title[i]+"</th>";
		}
		table += "</tr></thead><tbody>";
		for(var i = 0;i < data.length;i++){
			table += "<tr>";
			var res = data[i];
			for (var j=0;j<colnum.length;j++) {
				var value = typeof(res[colnum[j]]) == "undefined"?"":res[colnum[j]];
				if(j==0){
					table += "<td style='width: 60px;'>"+value+"</td>";
				}else{
					table += "<td>"+value+"</td>";
				}
			}
			table += "</tr>";
		}
		
		if(data.length == 0) {
			table += "<tr> <td colspan=" + title.length + ">Empty !</td></tr>";
		}
		
		table += "</tbody></table>";
		return table;
	}
	
}



//页面组件、数据、x轴、y轴、标签
function createLineChart(ctx,data,x,y,label){
	$("#"+ctx).html("");
	Morris.Area({
		  element: ctx,
		  data: data,
		  xkey: x,
		  ykeys: [y],
		  labels: [label],
		  fillOpacity:0.5,
		  parseTime:false,
		  lineColors: ["#0B62A4"],
		  lineWidth:2
	});
}
//页面组件、数据、x轴、y轴、标签
function createBarChart(ctx,data,x,y,label){
	$("#"+ctx).html("");
	Morris.Bar({
		element: ctx,
		data: data,
		xkey: x,
		ykeys: [y],
		labels: [label]
	});
}
//页面组件、数据、x轴、y轴、标签
function createMuiltLineChart(ctx,cols, data,x,y){
	$("#"+ctx).html("");
	Morris.Area({
		element: ctx,
		data: data,
		xkey: x,
		ykeys: cols,
		labels: cols,
		fillOpacity:0.5,
		parseTime:false,
		lineWidth:2
	});
}

//页面组件、数据、x轴、y轴、标签
function createMuiltLineChartMore(ctx,cols, data,x){
	$("#"+ctx).html("");
	Morris.Area({
		element: ctx,
		data: data,
		xkey: x,
		ykeys: cols,
		labels: cols,
		fillOpacity:10,
		parseTime:false,
		lineWidth:2
	});
}


$(document).ajaxStart(function(){
	$("form").find("div.error").removeClass("error");
	$("form").find("span.s2_help_inline").remove();
	parent.loading();
});
$(document).ajaxComplete(function(){
	parent.loadingOut();
});

