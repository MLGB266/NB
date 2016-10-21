
<!--进入HOME页-->

<div class="container"><br>
		<h2>博客用户基本信息</h2>
    <p>
      {{range .Blogs}}
      <br>
      <a href="/view/{{.Id}}">{{.}}</a>
      <br>
      {{end}}
    </p>
    <h3><a href="/add">新增博客用户</a></h3>
</div>








