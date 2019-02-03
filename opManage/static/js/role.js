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
            showCheckbox:true,
            //在参数中使用回调函数来绑定任何事件，或者使用标准的jQuery .on()方法来绑定事件
            onNodeChecked:nodeChecked ,
        	onNodeUnchecked:nodeUnchecked,
            
        });

        var nodeCheckedSilent = false;
		function nodeChecked (event, node){
    		if(nodeCheckedSilent){
        		return;
    		}
    		nodeCheckedSilent = true;
    		checkAllParent(node)
    		checkAllSon(node)
    		var parentNode = $("#tree").treeview("getNode", node.parentId);      
            setParentNodeCheck(node);
    		nodeCheckedSilent = false;
		}
 
		var nodeUncheckedSilent = false;
		function nodeUnchecked  (event, node){
		    if(nodeUncheckedSilent)
		        return;
		    nodeUncheckedSilent = true;
		    uncheckAllParent(node)
		    uncheckAllSon(node)
		    var parentNode = $("#tree").treeview("getNode", node.parentId);  //获取父节点
            //var selectNodes = getChildNodeIdArr(node);     
            setParentNodeCheck(node);
		    nodeUncheckedSilent = false;
		}

	 	// 选中父节点时，选中所有子节点
		function setParentNodeCheck(node) {   
 	       	var parentNode = $("#tree").treeview("getNode", node.parentId);   
 	        if (parentNode.nodes) {    
 		        var checkedCount = 0;    
	            for (x in parentNode.nodes) {     
	                if (parentNode.nodes[x].state.checked) {      
	                    checkedCount ++;     
	               } else {      
	                   break;     
	               }    
	           }    
	           if (checkedCount == parentNode.nodes.length) {  //如果子节点全部被选 父全选
	               $("#tree").treeview("checkNode", parentNode.nodeId);
	               setParentNodeCheck(parentNode);    
	           }else {   //如果子节点未全部被选 父未全选
	               $('#tree').treeview('uncheckNode', parentNode.nodeId); 
	               setParentNodeCheck(parentNode);        
	           }   
	       }  
	   } 



	   	// 取消父节点时 取消所有子节点
	   function setChildNodeUncheck(node) { 
	       if (node.nodes) {   
	           var ts = [];    //当前节点子集中未被选中的集合 
	           for (x in node.nodes) { 
	               if (!node.nodes[x].state.checked) {  
                   ts.push(node.nodes[x].nodeId);  
	               } 
	               if (node.nodes[x].nodes) {      
	                   var getNodeDieDai = node.nodes[x];      
	                   for (j in getNodeDieDai) {
	                       if (!getNodeDieDai.nodes[x].state.checked) {        
	                           ts.push(getNodeDieDai[j]); 
	                       }    
	                   }     
	               }    
           }   
	       }
	       return ts;  
	   }


 
		//选中全部父节点
		function checkAllParent(node){
		    $('#tree').treeview('checkNode',node.nodeId,{silent:true});
		    var parentNode = $('#tree').treeview('getParent',node.nodeId);
		    if(!("id" in parentNode)){
		        return;
		    }else{
		        checkAllParent(parentNode);
		    }
		}
		//取消全部父节点
		function uncheckAllParent(node){
		    $('#tree').treeview('uncheckNode',node.nodeId,{silent:true});
		    var siblings = $('#tree').treeview('getSiblings', node.nodeId);
		    var parentNode = $('#tree').treeview('getParent',node.nodeId);
		    if(!("id" in parentNode)) {
		        return;
		    }
		    var isAllUnchecked = true;  //是否全部没选中
		    for(var i in siblings){
		        if(siblings[i].state.checked){
		            isAllUnchecked=false;
		            break;
		        }
		    }
		    if(isAllUnchecked){
		        uncheckAllParent(parentNode);
		    }
		 
		}
 
		//级联选中所有子节点
		function checkAllSon(node){
		    $('#tree').treeview('checkNode',node.nodeId,{silent:true});
		    if(node.nodes!=null&&node.nodes.length>0){
		        for(var i in node.nodes){
		            checkAllSon(node.nodes[i]);
		        }
		    }
		}
		//级联取消所有子节点
		function uncheckAllSon(node){
		    $('#tree').treeview('uncheckNode',node.nodeId,{silent:true});
		    if(node.nodes!=null&&node.nodes.length>0){
		        for(var i in node.nodes){
		            uncheckAllSon(node.nodes[i]);
		        }
		    }
		}


        $("#btn").click(function (e) {

            var arr = $('#tree').treeview('getSelected');

            alert(JSON.stringify(arr));
            for (var key in arr) {
                alert(arr[key].id);
            }

        })

})