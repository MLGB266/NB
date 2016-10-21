{{define "head"}}
<!DOCTYPE html>
<html>
  	<head>
		<link rel="shortcut icon" href="/static/img/favicon.png" />
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<title>我的 beego 博客</title>
		<style>
			/*----------------------------------------------------------导航栏-----------------------------------*/
			.nav ul{
				width: auto;/*设置元素宽度为980px*/
				border: 1px solid #000;/*设置一个颜色为#000，宽度为1px的边框*/
				margin: 0px 100px 0px 100px;/*页面边距上、右、下、左*/
			}
			.nav ul li{
				float: left;/*悬浮框*/
				margin: 20px auto ;
			}
			.nav ul li a{
				width: 50px;/*设置元素宽为80px*/
				height: 30px;/*设置高度为28px*/
				line-height: 30px;/*设置行距为28px，让文字在每行的中间位置*/
				background: black;/*设置元素的背景为红色*/
				color: #FFF;/*文字颜色是白色*/
				margin:5px 10px;/*上下边距  左右边距*/
				font-size:12px;/*字体大小*/
				display:block;/*块级元素转变*/
				text-align: center;/*文本居中*/
				text-decoration: none;/*去掉下划线*/
			}
			.nav ul li a:hover{
				width: 80px;
				height: 30px;
				line-height: 30px;
				background: black;
				color: red;
				border: 0px solid red;	
			}

			/*----------------------------------------------------------登录-----------------------------------*/
			.pull-right ul {
				width: auto;
				border: 1px solid #000;
				margin: 0px 100px 0px 100px;
			}
			.pull-right ul li{
				float: right;
				margin: 20px auto ;
			}
			.pull-right ul li a{
				width: 100px;
				height: 30px;
				line-height: 30px;/*行距*/
				background: black;
				color: #FFF;/*文字颜色是白色*/
				margin:5px 10px;
				font-size:12px;/*字体大小*/
				display:block;/*块级元素转变*/
				text-align: center;/*文本居中*/
				text-decoration: none;/*去掉下划线*/
			}
			.pull-right ul li a:hover{
				width: 120px;
				height: 30px;
				line-height: 30px;
				background: black;
				color: red;
				border: 0px solid red;	
			}
			

			.container {
				width: auto;
				border: 0px solid #000;
				margin: 80px 100px auto 150px;/*页面边距上、右、下、左*/
			}

			.author {
				float: none;
				text-align: center;
				width:auto;
				height:auto;
				border: 0px solid #000;
				margin: 40px 0px 30px 0px;
				
			}

		</style>

		<style type="text/css">
	        a:link,a:visited{
	         text-decoration:none;  /*超链接无下划线*/
	        }
	        a:hover{
	         text-decoration:underline;  /*鼠标放上去有下划线*/
      		}
      	</style>

	</head>

	<body>
		
		<div class="nav">
			<ul>
				{{if .IsLogin}}
				<li><a href="/">首页</a></li>
				<li><a href="http://2316155791.qzone.qq.com" >空间</a></li>
				<li><a href="http://oa.yolo24.com/">国美</a></li>
				<li><a href="https://www.google.co.id/?gws_rd=cr&ei=xloDVtW_Go-yuASXt7DwDQ">谷歌</a></li>
				{{else}}
				<li><a href="/login">首页</a></li>
				<li><a href="/login">空间</a></li>
				<li><a href="/login">国美</a></li>
				<li><a href="/login">谷歌</a></li>
				{{end}}
			</ul>
		</div>

		<div class="pull-right">
			<ul>
				{{if .IsLogin}}
				<li><a href="/login?exit=true">123456退出</a></li>
				{{else}}
				<li><a href="/login">管理员登录</a></li>
				{{end}}
			</ul>
		</div>

{{end}}
