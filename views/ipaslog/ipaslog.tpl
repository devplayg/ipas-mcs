{{template "base.tpl" .}}

{{define "contents"}}
<!--user icon in two different styles-->22
<div id="toolbar-log">
    <form id="form-filter" role="form" method="post">
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
                    <button type="submit" class="btn blue"/>{{i18n .Lang "search"}}</button>
                    <a class="btn btn-default" href="">{{i18n .Lang "cancel"}}</a>

                    {{if eq .filter.FastPaging "on"}} {{/* 고속 페이징 */}}
                    <div class="input-group btn-group btn-page-group">
                        <button type="button" class="btn btn-primary btn-move-page btn-prev" data-direction="-1" data-loading-text="&lt;">&lt;</button>
                        <button type="button" class="btn btn-primary btn-move-page btn-page-text" data-direction="0">1</button>
                        <button type="button" class="btn btn-primary btn-move-page btn-next" data-direction="1" data-loading-text="&gt;">&gt;</button>
                    </div>
                    {{end}}
                    <a href="#" data-toggle="modal" data-target="#modal-filter"><i class="fa fa-filter icon-filter hidden font-red"></i>{{i18n .Lang "detail_filter"}}</a>
                </div>
            </div>
        </div>

        <!-- 상세필터 -->
        <div id="modal-filter" class="modal fade" tabindex="-1" role="dialog">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title"><i class="fa fa-filter"></i> Filter</h4>
                    </div>

                    <div class="modal-body">
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label">GUID</label>
                                <input type="text" class="form-control" name="guid" value="{{.filter.Guid}}">
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label">Risk level</label>
                                <select name="risk_level[]" class="selectpicker" data-width="100%" data-size="5" multiple title="Risk level">
                                    <option value="1">Risk1</option>
                                    <option value="2">Risk2</option>
                                    <option value="3">Risk3</option>
                                    <option value="4">Risk4</option>
                                    <option value="5">Risk5</option>
                                </select>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label">Page size</label>
                                <input type="text" class="form-control mask-09999" name="limit" value="{{.filter.Limit}}">
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="mt-checkbox mt-checkbox-outline mt30">
                                    <input type="checkbox" name="fastPaging" {{if eq .filter.FastPaging "on"}}checked{{end}}> Page size
                                    <span></span>
                                </label>
                            </div>
                        </div>
                    </div><!-- modal-body-->
                    <div class="modal-footer">
                        <button type="submit" class="btn btn-primary">{{i18n .Lang "search"}}</button>
                        <button type="button" class="btn btn-default" data-dismiss="modal">{{i18n .Lang "close"}}</button>
                    </div>
                </div>
            </div>
        </div> <!-- #modal-filter -->
    </form>
</div>
<pre>
startDate={{.filter.StartDate}}
endDate={{.filter.EndDate}}
fastPaging={{.filter.FastPaging}}
guid={{.filter.Guid}}
limit={{.filter.Limit}}
sort={{.filter.Sort}}
order={{.filter.Order}}
{{range .filter.RiskLevel}}&risk_level[]={{.}}{{end}}
{{.Lang}}
</pre>

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
        data-sort-name="{{.filter.Sort}}"
        data-sort-order="{{.filter.Order}}"
        {* 페이징 *}
        {{if eq .filter.FastPaging "on"}} {* 고속 페이징 *}
            data-side-pagination="client"
        {{else}} {* 일반 페이징 *}
            data-url="/ipaslogs?startDate={{.filter.StartDate}}&endDate={{.filter.EndDate}}&fastPaging={{.filter.FastPaging}}&guid={{.filter.Guid}}{{range .filter.RiskLevel}}&risk_level[]={{.}}{{end}}"
            data-pagination="true"
            data-side-pagination="server"
            data-pagination-loop="false"
        {{end}}
>
    <thead>
    <tr>
        <th data-field="date" data-sortable="true" data-formatter="dateFormatter">Date</th>
        <th data-field="equip_id" data-formatter="equipIdFormatter" data-sortable="true">ID</th>
        <th data-field="target" data-formatter="targetEquipIdFormatter" data-sortable="true">Target</th>
        <th data-field="speeding_count" data-sortable="true">Speeding</th>
        <th data-field="shock_count" data-sortable="true">Shock</th>
        <th data-field="latitude" data-sortable="true">위도</th>
        <th data-field="longitude" data-sortable="true">경도</th>
        <th data-field="warning_dist" data-sortable="true">경보 설정(M)</th>
        <th data-field="caution_dist" data-sortable="true">주의 설정(M)</th>
        <th data-field="v2v_dist" data-sortable="true">V2V 거리</th>
        <th data-field="shock_threshold" data-sortable="true">충격 임계치</th>
        <th data-field="speed_threshold" data-sortable="true">속도 임계치</th>
    </tr>
    </thead>
</table>
{{end}}

{{define "javascript"}}
<script src="/static/modules/{{.ctrl}}/{{.ctrl}}.js"></script>
{{end}}
