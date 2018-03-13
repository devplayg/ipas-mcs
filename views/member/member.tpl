{{template "base.tpl" .}}


{{define "contents"}}
        <div class="strong b">
            asdfljkashdfkljhasd kljfhasdjklfha lksdf
        </div>
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


<div class="modal modal-member fade" id="modal-member-add" tabindex="-1" role="basic" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <form role="form" id="form-member-add" class="form-member">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-hidden="true"></button>
                    <h4 class="modal-title"><i class="fa fa-plus"></i> {{i18n .Lang "registration"}}</h4>
                </div>
                <div class="modal-body">
                    <div class="row">
                        <div class="col-sm-6">
                            <div class="form-group">
                                <label class="control-label">{{i18n .Lang "username"}}</label>
                                <input name="username" type="text" class="form-control">
                            </div>
                            <div class="form-group">
                                <label class="control-label">{{i18n .Lang "member.name"}}</label>
                                <input name="name" type="text" class="form-control" autocomplete="name">
                            </div>
                            <div class="form-group">
                                <label class="control-label">{{i18n .Lang "password"}}</label>
                                <input name="password" type="password" class="form-control" autocomplete="password">
                            </div>
                            <div class="form-group">
                                <label class="control-label">{{i18n .Lang "confirm_password"}}</label>
                                <input id="password_confirm" name="password_confirm" type="password" class="form-control" autocomplete="password_confirm">
                            </div>
                            <div class="form-group">
                                <label class="control-label">{{i18n .Lang "email"}}</label>
                                <input name="email" type="text" class="form-control" autocomplete="email">
                            </div>
                        </div>
                        <div class="col-sm-6">
                            <div class="form-group">
                                <label class="control-label">Timezone</label>
                                <select name="timezone" class="selectpicker" data-width="100%">
                                    <option value="Asia/Seoul">(+09:00) Asia / Seoul</option>
                                    <option value="America/Los_Angeles">(−08:00) America / Los Angeles</option>
                                    <option value="America/Santiago">(−04:00) America / Santiago</option>
                                </select>
                            </div>
                            <div class="form-group">
                                <label for="allowed_ip" class="control-label">Allowed IP</label>
                                <textarea class="form-control" name="allowed_ip" rows="4" style="resize: none">0.0.0.0/0,
10.0.7.194, 255.255.255.256,
                                3.3.3.3,,,</textarea>
                            </div>
                            <div class="form-group">
                                <label>User Groups</label>
                                <div class="mt-checkbox-list">
                                    <label class="mt-checkbox mt-checkbox-outline"> Administrator
                                        <input type="checkbox" name="user_groups" value="10" />
                                        <span></span>
                                    </label>
                                    <label class="mt-checkbox mt-checkbox-outline"> Observer
                                        <input type="checkbox" name="user_groups" value="8" />
                                        <span></span>
                                    </label>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="alert alert-danger hidden">
                        <i class="fa fa-warning"></i> <span class="message"></span>
                    </div>
                </div><!-- modal-body -->
                <div class="modal-footer">
                    <button type="submit" class="btn btn-primary">{{i18n .Lang "save"}}</button>
                    <button type="button" class="btn btn-default" data-dismiss="modal">{{i18n .Lang "close"}}</button>
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
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
