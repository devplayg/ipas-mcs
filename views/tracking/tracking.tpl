{{template "base.tpl" .}}

{{define "css"}}
{{end}}

{{define "contents"}}
<style>
    .pagination-info, .pagination-detail {
        display:none
    }
    .page-list {
        display:none
    }
</style>

<form id="form-filter" role="form" method="post">
{{ .xsrfdata }}

    <div class="form-body">
        <div class="form-inline">
            <div class="form-group">

                <!-- 검색 시작날짜 -->
                <div class="input-group date datetime" data-date="1979-09-16T05:25:07Z" data-date-format="yyyy-mm-dd HH:ii" data-link-field="">
                    <input class="form-control mask-yyyymmddhhii" size="16" type="text" name="start_date" value="{{.filter.StartDate}}">
                    <span class="input-group-addon"><i class="glyphicon glyphicon-th"></i></span>
                </div>

                <!-- 검색 끝날짜 -->
                <div class="input-group date datetime">
                    <input class="form-control mask-yyyymmddhhii" size="16" type="text" name="end_date" value="{{.filter.EndDate}}">
                    <span class="input-group-addon"><i class="glyphicon glyphicon-th"></i></span>
                </div>

                <select id="select-equipId" name="equip_id_with_org" class="selectpicker" data-live-search="true"></select>

                <button type="submit" class="btn blue">{{i18n .Lang "tracking"}}</button>

                <a href="#" data-toggle="modal" data-target="#modal-filter"><i class="fa fa-filter icon-filter hidden font-red"></i>{{i18n .Lang "detail_filter"}}</a>
            </div>
        </div>
    </div>
</form>


운행건수
충격/과속/근접 건수 - 추이 그래프

운행시간 이력(좌), 이력(우)


<div class="row hide">
    <div class="col-lg-4">
        <div class="portlet light bordered">
            <div class="portlet-body pt0">
                <table  id="table-log"
                        class="table-condensed"
                        data-toggle="table"
                        data-toolbar="#toolbar-ipas"
                        data-show-refresh="true"
                        data-search="true"
                        data-click-to-select="true"
                        {* 정렬 *}
                        data-sort-name=""
                        data-sort-order=""
                        {* 페이징 *}
                        data-pagination-v-align="bottom"
                        data-url="/ipaslist"
                        data-side-pagination="client"
                        data-pagination="true"
                        data-page-size="15"
                >
                    <thead>
                    <tr>
                        <th data-field="org_id" data-sortable="true" data-formatter="orgNameFormatter">{{i18n .Lang "org"}}</th>
                        <th data-field="group_id" data-sortable="true" data-formatter="groupNameFormatter">{{i18n .Lang "group"}}</th>
                        <th data-field="equip_id">{{i18n .Lang "tag"}}</th>

                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>
    <div class="col-lg-8">
        <div class="portlet light bordered">
            <div class="portlet-body pt0">
            </div>
        </div>
    </div>
</div>
{{end}}

{{define "javascript"}}
{{template "ipasreport/ipasreport.tpl" .}}
<script src="/static/modules/{{.ctrl}}/tracking.js"></script>
<script src="/static/modules/{{.ctrl}}/formatter.js"></script>
<script>
</script>
{{end}}