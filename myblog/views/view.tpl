

  <div class="container">
    <div class="row">
        
        <table class="table" border="0" width="20%">

          <tr>
            <th>ID号</th>
            <td width="50%">{{.Blogs.Id}}</td>          
          </tr>

          <tr>
            <th>姓名</th>
            <td width="50%">{{.Blogs.Name}}</td>          
          </tr>

          <tr>
            <th>性别</th>
            <td width="50%">{{.Blogs.Sex}}</td>          
          </tr>

          <tr>
            <th>地址</th>
            <td width="50%">{{.Blogs.Address}}</td>          
          </tr>

          <tr>
            <th>标题</th>
            <td width="50%">{{.Blogs.Title}}</td>          
          </tr>

          <tr>
            <th>操作</th>
            </th>
            <td width="50%">
              <a href="/edit/{{.Blogs.Id}}">编辑</a>
              <a href="/delete/{{.Blogs.Id}}">删除</a>
              <a href="/add">增加</a>
            </td>          
          </tr>

        </table>

        <h2>文章：{{.Blogs.Title}}</h2>
        {{.Blogs.Content}}    


    </div>
  </div>
