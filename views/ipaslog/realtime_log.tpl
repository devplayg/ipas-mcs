{{template "base.tpl" .}}
{{define "css"}}
<link href="/static/modules/{{.ctrl}}/realtime_log.css" rel="stylesheet" type="text/css" />
{{end}}

{{define "contents"}}
<!-- 자산 선택 -->
<div class="portlet light bordered mb10 pb8 pt8">
    <button class="btn default btn-start"><span class="text"><i class="fa fa-play"></i></span> {{i18n .Lang "monitoring"}}</button>
    <select id="select-orgs" name="org_id" class="selectpicker" multiple title="{{i18n .Lang "org"}}"  data-size="10" data-selected-text-format="count > 2"></select>
    <select id="select-groups" name="group_id" class="selectpicker" multiple title="{{i18n .Lang "group"}}"  data-size="10" data-selected-text-format="count > 2"></select>
    <button class="btn blue btn-apply">{{i18n .Lang "apply"}}</button>
    <a href="" class="btn default btn-action">{{i18n .Lang "cancel"}}</a>
    <span class="text-applied font-red hide ml10">{{i18n .Lang "applied"}}</span>
    <span class="text-updated font-red hide ml10">Updated</span>
</div>

<div class="row">
    <div class="col-md-6">
        <div class="portlet light bordered pt0">
            <div class="portlet-body pt0 mh">
                <div id="toolbar-event4">
                    <div class="caption s16 bold">
                        <i class="icon-size-actual"></i>
                        <span class="caption-subject uppercase"> {{i18n .Lang "monitoring.proximity event"}}</span>
                        <span class="caption-helper s13 font-grey-salsa">
                            {{if ne .Lang "en-us" }}Proximity events{{end}}
                        </span>
                    </div>
                </div>
                <table  id="table-event4"
                        data-toolbar="#toolbar-event4"
                        class="table-data table-condensed"
                        data-toggle="table"
                        data-show-refresh="true"
                        data-show-columns="true"
                        {* 페이지 크기*}
                        data-page-size="{{.filter.Limit}}"
                        {* 정렬 *}
                        data-sort-name="{{.filter.Sort}}"
                        data-sort-order="{{.filter.Order}}"
                        data-side-pagination="client"
                        data-url="/getRealTimeLogs?limit=9&event_type=4">
                    <thead>
                    <tr>
                        <th data-field="date" data-width="10%" data-sortable="true" data-formatter="shortDateFormatter">{{i18n .Lang "occurrence date"}}</th>
                        <th data-field="org_name" data-sortable="true">{{i18n .Lang "org"}}</th>
                        <th data-field="group_name" data-sortable="true" data-visible="false">{{i18n .Lang "group"}}</th>
                        <th data-field="event_type" data-sortable="true" data-formatter="ipaslogEventTypeFormatter">{{i18n .Lang "ipas.action"}}</th>
                        <th data-field="equip_id" data-formatter="ipasEquipIdFormatter" data-sortable="true">{{i18n .Lang "tag"}}</th>
                        <th data-field="distance" data-sortable="true" data-formatter="ipaslogDistanceFormatter">{{i18n .Lang "distance"}} (m)</th>
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
    </div>

    <div class="col-md-6">
        <div class="portlet light bordered pt0">
            <div class="portlet-body pt0 mh">
                <div id="toolbar-event2">
                    <div class="caption s16 bold">
                        <i class="fa fa-bolt"></i>
                        <span class="caption-subject uppercase"> {{i18n .Lang "monitoring.shock event"}}</span>
                        <span class="caption-helper s13 font-grey-salsa">
                        {{if ne .Lang "en-us" }}Shock events{{end}}
                        </span>
                    </div>
                </div>
                <table  id="table-event2"
                        data-toolbar="#toolbar-event2"
                        class="table-data table-condensed"
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
                        data-url="/getRealTimeLogs?limit=9&event_type=2">
                    <thead>
                    <tr>
                        <th data-field="date" data-width="10%" data-sortable="true" data-formatter="shortDateFormatter">{{i18n .Lang "occurrence date"}}</th>
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
    </div>
</div>

<div class="row">
    <div class="col-md-6">
        <div class="portlet light bordered pt0">
            <div class="portlet-body pt0 mh">
                <div id="toolbar-event3">
                    <div class="caption s16 bold">
                        <i class="icon-speedometer"></i>
                        <span class="caption-subject uppercase"> {{i18n .Lang "monitoring.speeding event"}}</span>
                        <span class="caption-helper s13 font-grey-salsa">
                        {{if ne .Lang "en-us" }}Shock events{{end}}
                        </span>
                    </div>
                </div>
                <table  id="table-event3"
                        class="table-data table-condensed"
                        data-toolbar="#toolbar-event3"
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
                        data-url="/getRealTimeLogs?limit=9&event_type=3">
                    <thead>
                    <tr>
                        <th data-field="date" data-width="10%" data-sortable="true" data-formatter="shortDateFormatter">{{i18n .Lang "occurrence date"}}</th>
                        <th data-field="org_name" data-sortable="true">{{i18n .Lang "org"}}</th>
                        <th data-field="group_name" data-sortable="true" data-visible="false">{{i18n .Lang "group"}}</th>
                        <th data-field="event_type" data-sortable="true" data-formatter="ipaslogEventTypeFormatter">{{i18n .Lang "ipas.action"}}</th>
                        <th data-field="equip_id" data-formatter="ipasEquipIdFormatter" data-sortable="true">{{i18n .Lang "tag"}}</th>
                        <th data-field="speed" data-sortable="true" data-formatter="ipaslogSpeedingFormatter">{{i18n .Lang "speed"}} <small>(km/h)</small></th>
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
    </div>

    <div class="col-md-6">
        <div class="portlet light bordered pt0">
            <div class="portlet-body pt0 mh">
                <div id="toolbar-event1">
                    <div class="caption s16 bold">
                        <i class="icon-power"></i>
                        <span class="caption-subject uppercase"> {{i18n .Lang "monitoring.startup event"}}</span>
                        <span class="caption-helper s13 font-grey-salsa">
                        {{if ne .Lang "en-us" }}Startup  events{{end}}
                        </span>
                    </div>
                </div>
                <table  id="table-event1"
                        class="table-data table-condensed"
                        data-toolbar="#toolbar-event1"
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
                        data-url="/getRealTimeLogs?limit=9&event_type=1">
                    <thead>
                    <tr>
                        <th data-field="date" data-width="10%" data-sortable="true" data-formatter="shortDateFormatter">{{i18n .Lang "occurrence date"}}</th>
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
    </div>
</div>
{{end}}

{{define "javascript"}}
{{template "ipasreport/ipasreport.tpl" .}}
<script src="/static/modules/{{.ctrl}}/realtime_log.js"></script>
<script src="/static/modules/{{.ctrl}}/formatter.js"></script>
{{end}}