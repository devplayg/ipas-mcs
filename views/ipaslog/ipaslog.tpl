{{template "base.tpl" .}}

{{define "contents"}}
<div id="toolbar-log">
    <form id="form-filter" role="form" method="post" action="/ipaslog">
        {{ .xsrfdata }}

        <div class="form-body">
            <div class="form-inline">
                <div class="form-group">

                    <!-- Start date -->
                    <div class="input-group date datetime" data-date="1979-09-16T05:25:07Z" data-date-format="yyyy-mm-dd HH:ii" data-link-field="">
                        <input class="form-control mask-yyyymmddhhii" size="16" type="text" name="startDate" value="{{.filter.StartDate}}">
                        <span class="input-group-addon"><i class="glyphicon glyphicon-th"></i></span>
                    </div>

                    <!-- End date -->
                    <div class="input-group date datetime">
                        <input class="form-control mask-yyyymmddhhii" size="16" type="text" name="endDate" value="{{.filter.EndDate}}">
                        <span class="input-group-addon"><i class="glyphicon glyphicon-th"></i></span>
                    </div>

                    <!-- Buttons -->
                    <button type="submit" class="btn btn-primary" value="Search"/>Search</button>
                    <a class="btn btn-default" href=".">Cancel</a>

                    <div class="input-group btn-group btn-page-group">
                        <button type="button" class="btn btn-primary btn-move-page btn-prev" data-direction="-1" data-loading-text="&lt;">&lt;</button>
                        <button type="button" class="btn btn-primary btn-move-page btn-page-text" data-direction="0">1</button>
                        <button type="button" class="btn btn-primary btn-move-page btn-next" data-direction="1" data-loading-text="&gt;">&gt;</button>
                    </div>
                    <a href="#" data-toggle="modal" data-target="#modal-filter"><i class="fa fa-filter icon-filter hidden font-red"></i> Filter</a>
                </div>
            </div>
        </div>

        <div id="modal-filter" class="modal fade" tabindex="-1" role="dialog">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title"><i class="fa fa-filter"></i> Filter</h4>
                    </div>

                    <div class="modal-body">
                        <div class="row lh18">
                            <div class="col-lg-12 form-group">
                                <div class="md-checkbox">
                                    <input type="checkbox" name="fastPaging" class="md-check" {{if eq .filter.FastPaging "on"}}checked{{end}}>
                                    <label for="FastPaging">
                                        <span></span>
                                        <span class="check"></span>
                                        <span class="box"></span>
                                        <span style="margin: -10px 0px 0px 30px;">Fast Paging</span>
                                    </label>
                                </div>
                            </div>
                        </div>

                    </div>
                    <div class="modal-footer">
                        <button type="submit" class="btn btn-primary">Search</button>
                        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    </div>
                </div>
            </div>
        </div> <!-- #modal-filter -->
    </form>
</div>
data-url="/ipaslog/getlogs/?StartDate={{.filter.StartDate}}&EndDate={{.filter.EndDate}}&FastPaging={{.filter.FastPaging}}"
<table  id="table-log"
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
        data-page-size="{{.filter.Limit}}"
        {* 정렬 *}
        data-sort-name="date"
        data-sort-order="desc"
        {* 페이징 *}
    {{if eq .filter.FastPaging "on"}} {* 고속 페이징 *}
            data-side-pagination="client"
    {{else}} {* 일반 페이징 *}
        data-url="/ipaslog/getlogs/?StartDate={{.filter.StartDate}}&EndDate={{.filter.EndDate}}&FastPaging={{.filter.FastPaging}}"
            data-pagination="true"
            data-side-pagination="server"
            data-pagination-loop="false"
            data-page-list="[10, 25, 50, 100, 1000, 2000, 5000]"
            data-pagination-v-align="both"
    {{end}}

>
    <thead>
    <tr>
        <th data-field="no" data-sortable="false">No</th>
        <th data-field="date" data-sortable="true" data-formatter="dateFormatter">Date</th>
        <th data-field="asset_level1" data-sortable="true">Group1</th>
        <th data-field="asset_level2" data-sortable="true">Group2</th>
        <th data-field="risk_level" data-sortable="true">Group2</th>
        <th data-field="guid" data-sortable="true">GUID</th>
        <th data-field="contents" data-sortable="true">Contents</th>
    </tr>
    </thead>
</table>
{{end}}

{{define "javascript"}}
<script src="/static/modules/{{.ctrl}}/ipaslog.js"></script>
{{end}}
