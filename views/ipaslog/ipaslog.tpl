{{template "base.tpl" .}}

{{define "contents"}}
{{.ctrl}}/{{.act}}
{{end}}

{{define "javascript"}}
<script src="/static/modules/{{.ctrl}}/ipaslog.js"></script>
<script>

</script>
{{end}}
