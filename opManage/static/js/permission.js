$(function () {
        function getTree() {
            var data = [
            	{
            		text:"所有权限",
            		nodeId:"1001",
            		icon:"glyphicon glyphicon-home",
            		nodes:[
            			{
            				text:"任务管理",
            				nodeId:"2001",          				
            				nodes:[
            					{
            						text:"sister",
            						nodeId:"3101",
            	
            					},
            					{
            						text:"body",
            						nodeId:"3102",
            						
            					}
            				]
            			},
            			{
            				text:"大资产管理",
            				nodeId:"2002",
            				
            				nodes:[
            					{
            						text:"www",
            						nodeId:"3201",
            						
            					},
            					{
            						text:"https",
            						nodeId:"3202",
            					}
            				]
            			},
            			{
            				text:"系统管理",
            				nodeId:"2003",
            				
            			},
            			{
            				text:"日志管理",
            				nodeId:"2004",
            				
            				nodes:[
            					{
            						text:"python",
            						nodeId:"3401",
            						
            					},
            					{
            						text:"golang",
            						nodeId:"3402",
            						
            					}
            				]
            			},
            			{
            				text:"运维工具",
            				nodeId:"2005",
            			},
            			{
            				text:"运维权限",
            				nodeId:"2006"
            			}
            		]

            	}
            ]
            return data;
        }


        var obj = {};
        obj.text = "123";
        $('#tree').treeview({
            data: getTree(),         // data is not optional
            levels: 5,
            nodeIcon:"glyphicon glyphicon-envelope",
        });


        $("#btn").click(function (e) {

            var arr = $('#tree').treeview('getSelected');

            alert(JSON.stringify(arr));
            for (var key in arr) {
                alert(arr[key].id);
            }

        })

})