{{template "base.tpl" .}}

{{define "contents"}}
<div class="portlet light bordered">
    <div class="portlet-body pt0">
        <div id="toolbar-log">

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

                            <!-- 자산 선택 -->
                            <select id="select-orgs" name="org_id" class="selectpicker" multiple title="{{i18n .Lang "org"}}"  data-size="10" data-selected-text-format="count > 2"></select>
                            <select id="select-groups" name="group_id" class="selectpicker" multiple title="{{i18n .Lang "group"}}"  data-size="10" data-selected-text-format="count > 2"></select>

                            <!-- Buttons -->
                            <button type="submit" class="btn blue"/>{{i18n .Lang "search"}}</button>
                            <a class="btn btn-default" href="">{{i18n .Lang "cancel"}}</a>

                        {{if eq .filter.FastPaging "on"}} {{/* 고속 페이징 */}}
                            <div class="input-group btn-group btn-page-group">
                                <button type="button" class="btn blue btn-move-page btn-prev" data-direction="-1" data-loading-text="&lt;">&lt;</button>
                                <button type="button" class="btn blue btn-move-page btn-page-text" data-direction="0">1</button>
                                <button type="button" class="btn blue btn-move-page btn-next" data-direction="1" data-loading-text="&gt;">&gt;</button>
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
                                    <div class="col-sm-6 form-group">
                                        <label class="control-label">{{i18n .Lang "tag"}}</label>
                                        <input type="text" class="form-control" name="tag_pattern" value="{{.filter.TagPattern}}">
                                    </div>
                                    <div class="col-sm-6 form-group ">
                                    </div>
                                </div>
                                <div class="row">
                                    <div class="col-sm-4 form-group">
                                        <label class="control-label">Page size</label>
                                        <input type="text" class="form-control mask-09999" name="limit" value="{{.filter.Limit}}">
                                    </div>
                                    <div class="col-sm-4 form-group">
                                        <label class="mt-checkbox mt-checkbox-outline mt30">
                                            <input type="checkbox" name="fast_paging" {{if eq .filter.FastPaging "on"}}checked{{end}}> {{i18n .Lang "fast_paging"}}
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

        <table  id="table-log"
                class="table-condensed"
                data-toggle="table"
                data-toolbar="#toolbar-log"
                data-show-refresh="true"
                data-show-columns="true"
                {* 내보내기 *}
                data-show-export="true"
                data-export-types="['csv', 'excel']"
                {*Row강조*}
                data-row-style="ipasEventRowStyle"
                {* 정렬 *}
                data-sort-name="{{.filter.Sort}}"
                data-sort-order="{{.filter.Order}}"
                {* 페이징 *}
                data-page-size="{{.filter.Limit}}"
                data-pagination-v-align="both"
        {{if eq .filter.FastPaging "on"}} {* 고속 페이징 *}
                data-side-pagination="client"
        {{else}} {* 일반 페이징 *}
                data-url="/getStatusLogs?start_date={{.filter.StartDate}}&end_date={{.filter.EndDate}}&fast_paging={{.filter.FastPaging}}&equip_id={{.filter.EquipId}}{{range .filter.OrgId}}&org_id={{.}}{{end}}{{range .filter.GroupId}}&group_id={{.}}{{end}}"
                data-pagination="true"
                data-side-pagination="server"
                data-pagination-loop="false"
        {{end}}
        >
            <thead>
            <tr>
                <th data-field="date" data-sortable="true" data-formatter="dateFormatter">{{i18n .Lang "occurrence date"}}</th>
                <th data-field="org_id" data-sortable="true" data-formatter="orgNameFormatter">{{i18n .Lang "org"}}</th>
                <th data-field="group_id" data-sortable="true" data-formatter="groupNameFormatter">{{i18n .Lang "group"}}</th>
                <th data-field="equip_id" data-formatter="ipasEquipIdFormatter" data-sortable="true">{{i18n .Lang "tag"}}</th>
                <th data-field="location" data-sortable="true" data-formatter="ipaslogLocationFormatter" data-align="center">{{i18n .Lang "location"}}</th>
                <th data-field="latitude" data-sortable="true" data-visible="false">{{i18n .Lang "latitude"}}</th>
                <th data-field="longitude" data-sortable="true" data-visible="false">{{i18n .Lang "longitude"}}</th>
                <th data-field="speed" data-sortable="true" data-formatter="ipaslogSpeedingFormatter">{{i18n .Lang "speed"}} <small>(km/h)</small></th>
                <th data-field="snr" data-sortable="true" data-formatter="snrFormatter">SNR</th>
                <th data-field="usim" data-sortable="true" data-visible="false">USIM</th>
                <th data-field="ip" data-sortable="true" data-formatter="int2ipFormatter" data-visible="false">IP</th>
                <th data-field="recv_date" data-sortable="true" data-formatter="dateFormatter" data-visible="false">{{i18n .Lang "received date"}}</th>
            </tr>
            </thead>
        </table>

    </div>
</div>
{{end}}

{{define "javascript"}}
{{template "ipasreport/ipasreport.tpl" .}}
<script src="/static/modules/{{.ctrl}}/status_log.js"></script>
<script src="/static/modules/{{.ctrl}}/formatter.js"></script>
<script>
    var filterUrl = "start_date={{.filter.StartDate}}&end_date={{.filter.EndDate}}&fast_paging=on&equip_id={{.filter.EquipId}}{{range .filter.EventType}}&event_type={{.}}{{end}}{{range .filter.OrgId}}&org_id={{.}}{{end}}{{range .filter.GroupId}}&group_id={{.}}{{end}}";
    // console.log(url);
</script>
{{end}}