{{template "head" .}}

<div>
	{{if .IsLogin}}
		{{.LayoutContent}}
	{{else}}
		<script language='javascript'>document.location = '/login'</script>
	{{end}}
</div>

{{template "footer" .}}