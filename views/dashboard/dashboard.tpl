{{template "base.tpl" .}}

{{define "contents"}}
<div class="row">
    <div class="col-md-6">
    </div>
</div>

{{end}}

{{define "javascript"}}
{{template "ipasreport/ipasreport.tpl" .}}
<script src="/static/modules/{{.ctrl}}/dashboard.js"></script>
{{end}}