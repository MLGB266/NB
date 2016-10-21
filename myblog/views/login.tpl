
<html>
  	<head>
		<link rel="shortcut icon" href="/static/img/favicon.png" />
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

		 <!-- Stylesheets -->
		<link href="/static/css/bootstrap.min.css" rel="stylesheet" />
		<title>登录 - 我的 beego 博客</title>
  	</head>
  	
  	<body>

    
		<div class="container" style="width:500px;">
			<form class="form-horizontal" method="POST" action="/login">
				<div class="form-group">
					<label class="col-lg-4 control-lable">请输入管理员账号：</label>
					<div class="col-lg-6">
						<input id="username" class="form-control" name="username" placeholder="Account">	
					</div>
				</div>

				<div class="form-group">
				    <label class="col-lg-4 control-label">请输入管理员密码：</label>
				    <div class="col-lg-6">
				      <input id="userpwd" type="password" class="form-control" name="userpwd" placeholder="Password">
				    </div>
				</div>

				<div class="form-group">
				    <div class="col-lg-offset-4 col-lg-10">
				      <div class="checkbox">
				        <label><input name="autoLogin" type="checkbox">下次自动登录</label>
				      </div>
				    </div>
				</div>



				<div class="form-group">
				    <div class="col-lg-offset-2 col-lg-10">
				    	<a href="/login"><button type="submit" class="btn btn-default" onclick="return checkInput();">登录</button></a>
				      	<button class="btn btn-default" onclick="return backToHome();">返回</button>
				      	<script type="text/javascript">
					      	function backToHome() {
					      		window.location.href = "/";
					      		return false;
					      	}

					      	function checkInput() {
					      		var username = document.getElementById("username");
					      		if (username.value.length == 0) {
					      			alert("请输入管理员帐号");
					      			return false;
					      		}

					      		var userpwd = document.getElementById("userpwd");
					      		if (userpwd.value.length == 0) {
					      			alert("请输入管理员密码");
					      			return false;
					      		}
					      	}
				      	</script>
				    </div>
				 </div>


			</form>
		</div>
	</body>
</html>

