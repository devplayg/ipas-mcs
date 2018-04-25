{{template "base.tpl" .}}

{{define "contents"}}
<div class="row">
    <div class="col-md-6">
        <div id="toolbar-event1">
            <h4 class="bold">{{i18n .Lang "monitoring.startup event"}}</h4>
        </div>
        <table  id="table-event1"
                data-toolbar="#toolbar-event1"
                class="table-condensed"
                data-toggle="table"
                data-show-refresh="true"
                data-show-columns="true"
                {*Row강조*}
                data-row-style="rowStyle"
                {* 페이지 크기*}
                data-page-size="{{.filter.Limit}}"
                {* 정렬 *}
                data-sort-name="{{.filter.Sort}}"
                data-sort-order="{{.filter.Order}}"
                data-side-pagination="client"
                data-url="/getRealTimeLogs?event_type=1"
        >
            <thead>
            <tr>
                <th data-field="date" data-sortable="true" data-formatter="shortDateFormatter">{{i18n .Lang "occurrence date"}}</th>
                <th data-field="org_name" data-sortable="true">{{i18n .Lang "org"}}</th>
                <th data-field="group_name" data-sortable="true" data-visible="false">{{i18n .Lang "group"}}</th>
                <th data-field="event_type" data-sortable="true" data-formatter="ipaslogEventTypeFormatter">{{i18n .Lang "ipas.action"}}</th>
                <th data-field="equip_id" data-formatter="ipasEquipIdFormatter" data-sortable="true">{{i18n .Lang "tag"}}</th>
                <th data-field="location" data-sortable="true" data-formatter="ipaslogLocationFormatter" data-align="center">{{i18n .Lang "location"}}</th>
                <th data-field="latitude" data-sortable="true" data-visible="false">{{i18n .Lang "latitude"}}</th>
                <th data-field="longitude" data-sortable="true" data-visible="false">{{i18n .Lang "longitude"}}</th>
                <th data-field="snr" data-sortable="true" data-formatter="snrFormatter">SNR&nbsp;&nbsp;</th>
                <th data-field="usim" data-sortable="true" data-visible="false">USIM</th>
            </tr>
            </thead>
        </table>
    </div>
    <div class="col-md-6" style="border-left: 1px dashed #acacac;">
        <div id="toolbar-event2">
            <h4 class="bold">{{i18n .Lang "monitoring.shock event"}}</h4>
        </div>
        <table  id="table-event2"
                data-toolbar="#toolbar-event2"
                class="table-condensed"
                data-toggle="table"
                data-show-refresh="true"
                data-show-columns="true"
                {*Row강조*}
                data-row-style="rowStyle"
                {* 페이지 크기*}
                data-page-size="{{.filter.Limit}}"
                {* 정렬 *}
                data-sort-name="{{.filter.Sort}}"
                data-sort-order="{{.filter.Order}}"
                data-side-pagination="client"
                data-url="/getRealTimeLogs?event_type=2"
        >
            <thead>
            <tr>
                <th data-field="date" data-sortable="true" data-formatter="shortDateFormatter">{{i18n .Lang "occurrence date"}}</th>
                <th data-field="org_name" data-sortable="true">{{i18n .Lang "org"}}</th>
                <th data-field="group_name" data-sortable="true" data-visible="false">{{i18n .Lang "group"}}</th>
                <th data-field="event_type" data-sortable="true" data-formatter="ipaslogEventTypeFormatter">{{i18n .Lang "ipas.action"}}</th>
                <th data-field="equip_id" data-formatter="ipasEquipIdFormatter" data-sortable="true">{{i18n .Lang "tag"}}</th>
                <th data-field="location" data-sortable="true" data-formatter="ipaslogLocationFormatter" data-align="center">{{i18n .Lang "location"}}</th>
                <th data-field="latitude" data-sortable="true" data-visible="false">{{i18n .Lang "latitude"}}</th>
                <th data-field="longitude" data-sortable="true" data-visible="false">{{i18n .Lang "longitude"}}</th>
                <th data-field="snr" data-sortable="true" data-formatter="snrFormatter">SNR&nbsp;&nbsp;</th>
                <th data-field="usim" data-sortable="true" data-visible="false">USIM</th>
            </tr>
            </thead>
        </table>
    </div>
</div>
{{end}}

{{define "javascript"}}
{{template "ipasreport/ipasreport.tpl" .}}
<script src="/static/modules/{{.ctrl}}/realtime_logs.js"></script>
<script src="/static/modules/{{.ctrl}}/formatter.js"></script>
{{end}}