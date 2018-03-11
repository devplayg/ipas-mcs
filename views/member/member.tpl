{{template "base.tpl" .}}


{{define "contents"}}
<table  id="table-member"
        class="table-condensed"
        data-toggle="table"
        data-toolbar="#toolbar-log"
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
        <th data-field="member_id" data-sortable="true">ID</th>
        <th data-field="username" data-sortable="true">Username</th>
        <th data-field="name" data-sortable="true">Name</th>
        <th data-field="position" data-sortable="true">Position</th>
        <th data-field="failed_login_count" data-sortable="true">Failed login</th>
        <th data-field="last_success_login" data-sortable="true">Last access</th>
    </tr>
    </thead>
</table>
{{end}}

{{define "javascript"}}
<script src="/static/modules/member/member.js"></script>
{{end}}
