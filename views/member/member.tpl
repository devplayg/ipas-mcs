{{template "base.tpl" .}}


{{define "contents"}}
<div id="toolbar-member">
    <button type="button" class="btn btn-primary btn-member-add" data-toggle="modal" data-target="#modal-member-add"><i class="fa fa-plus"></i> Add</button>
    <button type="button" class="btn btn-danger btn-member-remove">Delete</button>
</div>

<table  id="table-member"
        class="table-condensed"
        data-toggle="table"
        data-toolbar="#toolbar-member"
        data-show-refresh="true"
        data-show-columns="true"
        data-row-style="scoreRowStyle"
        {* 내보내기 *}
        data-show-export="true"
        data-export-types="['csv', 'excel']"
        {* 페이지 크기*}
        data-page-size="10"
        {* 정렬 *}
        data-sort-name="position"
        data-sort-order="desc"
        {* 일반 페이징 *}
        data-url="/members"
        data-pagination="true"
        data-side-pagination="server"
        data-pagination-loop="false"
>
    <thead>
    <tr>
        <th data-field="command" data-formatter="memberCommandFormatter" data-sortable="true">Command</th>
        <th data-field="member_id" data-visible="false" data-sortable="true">ID</th>
        <th data-field="username" data-sortable="true">Username</th>
        <th data-field="name" data-sortable="true">Name</th>
        <th data-field="position" data-sortable="true" data-formatter="memberPositionFormatter">Position</th>
        <th data-field="failed_login_count" data-sortable="true">Failed login</th>
        <th data-field="last_success_login" data-sortable="true">Last access</th>
    </tr>
    </thead>
</table>

<div class="modal fade" id="modal-member-add" tabindex="-1" role="basic" aria-hidden="true">
    <div class="modal-dialog modal-sm">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true"></button>
                <h4 class="modal-title"><i class="fa fa-plus"></i> {{i18n .Lang "registration"}}</h4>
            </div>
            <div class="modal-body">
                <div class="row">
                    <div class="col-lg-12 form-group">
                        <label class="control-label">{{i18n .Lang "username"}}</label>
                        <input name="username" type="text" class="form-control">
                    </div>
                </div>
                <div class="row">
                    <div class="col-lg-12 form-group">
                        <label class="control-label">{{i18n .Lang "member.name"}}</label>
                        <input name="name" type="text" class="form-control">
                    </div>
                </div>
                <div class="row">
                    <div class="col-lg-12 form-group">
                        <label class="control-label">{{i18n .Lang "password"}}</label>
                        <input name="password1" type="password" class="form-control">
                    </div>
                </div>
                <div class="row">
                    <div class="col-lg-12 form-group">
                        <label class="control-label">{{i18n .Lang "confirm_password"}}</label>
                        <input id="member_password2" name="password2" type="password" class="form-control">
                    </div>
                </div>
                <div class="row">
                    <div class="col-lg-12 form-group">
                        <label class="control-label">{{i18n .Lang "email"}}</label>
                        <input name="email" type="text" class="form-control">
                    </div>
                </div>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                <button type="button" class="btn btn-primary">{{i18n .Lang "save"}}</button>
            </div>
        </div>
        <!-- /.modal-content -->
    </div>
    <!-- /.modal-dialog -->
</div>
{{end}}

{{define "javascript"}}
<script>
    var positions = {{.positions}};
    console.log(positions);
</script>
<script src="/static/modules/{{.ctrl}}/{{.ctrl}}.js"></script>
<script src="/static/modules/{{.ctrl}}/formatter.js"></script>
{{end}}
