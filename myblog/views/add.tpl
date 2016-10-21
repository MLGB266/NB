<div class="container">

	<h3>新增博客数据</h3>
	<form action="" method="post">
		<p>ID号:<input type="text" name="Id" value="" placeholder="请输入Id"></p>
		<p>姓名:<input type="text" name="Name" value="" placeholder="请输入姓名"></p>
		<p>性别:<input type="text" name="Sex" value="" placeholder="请输入性别"></p>
		<p>地址:<input type="text" name="Address" value="" placeholder="请输入地址"></p>

	

		<p>标题:<input type="text" class="form-control" name="Title" value="{{.Post.Title}}"  placeholder="请输入标题"></p>
		
		<p>内容:</p>
		<p><textarea name="Content" cols="100" rows="15" class="form-control" placeholder="请编辑文章">{{.Post.Content}}</textarea></p>
		<p><input type="hidden" name="Id" value="{{.Post.Id}}"></p>
		<p><input type="submit" value="提交"></p>
		<p><a href="/">取消</a></p>
	</form>

</div>