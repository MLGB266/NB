

<div class="container">

	<h3>编辑基本信息</h3>
	<form action="" method="post">
		<p>ID号:{{.Post.Id}}</p>
		<p>姓名:<input type="text" name="Name" value="{{.Post.Name}}"></p>
		<p>性别:<input type="text" name="Sex" value="{{.Post.Sex}}"></p>
		<p>地址:<input type="text" name="Address" value="{{.Post.Address}}"></p>
		<p>标题:<input type="text" class="form-control" name="Title" value="{{.Post.Title}}" placeholder="请输入标题"></p>
		<p>内容:</p>
		<textarea name="Content" cols="100" rows="18" class="form-control" placeholder="请编辑文章">{{.Post.Content}}</textarea>
		<p><input type="hidden" name="Id" value="{{.Post.Id}}"></p>
		<p><input type="submit"></p>
		<p><a href="/view/{{.Post.Id}}">取消</a></p>
	</form>
	
</div>


  
  