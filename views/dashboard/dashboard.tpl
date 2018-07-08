{{template "base.tpl" .}}

{{define "css"}}
<link href="/static/modules/{{.ctrl}}/dashboard.css" rel="stylesheet" type="text/css" />
<link href="/static/plugins/morris.js/morris.css" rel="stylesheet" type="text/css" />
{{end}}

{{define "contents"}}
<div class="row">
    <div class="col-lg-4 col-md-6">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title hide">
                <div class="caption">
                    <i class="icon-grid"></i> {{i18n .Lang "summary"}}
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <form id="form-filter" role="form" method="post">
                    <div class="form-body">
                        <div class="form-inline">
                            <div class="form-group">
                                <!-- 검색 시작날짜 -->
                                <div class="input-group date datetime hide" data-date="1979-09-16T05:25:07Z" data-date-format="yyyy-mm-dd HH:ii" data-link-field="">
                                    <input class="form-control mask-yyyymmddhhii" size="16" type="text" name="start_date" value="{{.filter.StartDate}}" disabled>
                                    <span class="input-group-addon"><i class="glyphicon glyphicon-th"></i></span>
                                </div>

                                <!-- 검색 끝날짜 -->
                                <div class="input-group date datetime hide">
                                    <input class="form-control mask-yyyymmddhhii" size="16" type="text" name="end_date" value="{{.filter.EndDate}}" disabled>
                                    <span class="input-group-addon"><i class="glyphicon glyphicon-th"></i></span>
                                </div>

                                <select id="select-assets" name="org_id" class="selectpicker" data-size="10" data-selected-text-format="count > 2">
                                    <option value="-1/-1">{{i18n .Lang "total assets"}}</option>
                                </select>

                                <button type="button" class="btn default btn-start"><span class="text"><i class="fa fa-play"></i></span> {{i18n .Lang "monitoring"}}</button>
                                <span class="text-updated font-red hide ml10">Updated</span>
                            </div>
                        </div>
                    </div>
                </form>
                <div class="row mt15">
                    <div class="col-xs-6 text-center">
                        <h1 class="mt0 mb5 font-grey-mint bold">
                            <a href="#" class="btn-show-ipasstatus-on-modal font-grey-mint" data-query="    ">
                                <span class="count-total-tags"></span>
                            </a>
                        </h1>
                        Total tags
                    </div>
                    <div class="col-xs-6 text-center" style="border-left: 1px dashed #acacac;">
                        <h1 class="mt0 mb5 grey-salsa">
                            <a href="#" class="btn-show-ipaslog-on-modal font-grey-mint" data-query="stats_mode=1">
                                <span class="count-events"></span>
                            </a>
                        </h1>
                        Total events
                    </div>
                </div>
                <div class="row" style="border-top: 1px dashed #acacac; margin: 10px 0px 5px 0px;"></div>
                <div class="row" style="margin: 10px 0px 0px 0px;">
                    <div class="col-xs-12 text-center">
                        <h1 class="mt0 mb5">
                            <span class="count-optime"></span>
                        </h1>
                        <a href="#" class="btn-show-ipasstatus-on-modal font-grey-mint" data-query="equip_type={{.ZoneTag}}">
                            Operating time
                        </a>
                    </div>
                    <div class="col-xs-6 hide text-center" style="border-left: 1px dashed #acacac;">
                        <h1 class="mt0 mb5">
                            <div class="">
                                <span class="what-time"></span>
                            </div>
                        </h1>
                        <span class="what-date"></span>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="col-lg-4 col-md-6">
        <div class="portlet light bordered mh-rank pb0">
            <div class="portlet-body pt0 mh">
                <table width="100%">
                    <tr>
                        <td width="40%" align="center">
                            <div id="chart-eventType" class="mt10" style="height: 180px; width: 180px;"></div>
                        </td>
                        <td width="60%">
                            <div class="portlet-title mt5">
                                <div class="caption mb20 s16 pl10">
                                    <i class="icon-grid"></i> {{i18n .Lang "event type"}}
                                </div>
                            </div>
                            <div class="plr10">
                                <div class="progress mb5">
                                    <span id="pgb-shock" class="progress-bar progress-bar-info" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100"></span>
                                </div>
                                <div class="s12 ">
                                    {{i18n .Lang "shock"}} (<span class="count-shock"></span>)
                                    <span class="rate-shock pull-right">%</span>
                                </div>
                            </div>

                            <div class="plr10 mt10">
                                <div class="progress mb5">
                                    <span id="pgb-speeding" class="progress-bar progress-bar-warning" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100"></span>
                                </div>
                                <div class="s12 ">
                                    {{i18n .Lang "speeding"}} (<span class="count-speeding"></span>)
                                    <span class="rate-speeding pull-right">%</span>
                                </div>
                            </div>

                            <div class="plr10 mt10">
                                <div class="progress mb5">
                                    <span id="pgb-proximity" class="progress-bar progress-bar-danger" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100"></span>
                                </div>
                                <div class="s12 ">
                                    {{i18n .Lang "proximity"}} (<span class="count-proximity"></span>)
                                    <span class="rate-proximity pull-right">%</span>
                                </div>
                            </div>
                        </td>
                    </tr>
                </table>
            </div>
        </div>
    </div>

    <div class="col-lg-4 col-md-12">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-body pt0 mh">
                <div class="scroller" style="height: 200px;" data-always-visible="1" data-rail-visible="0">
                    <table class="table table-ranking"
                           data-classes="table-condensed table-no-bordered"
                           data-toggle="table"
                           data-cache="false"
                           data-sort-name="count"
                           data-sort-order="desc"
                           data-show-header="true"
                           data-query="/stats/activatedGroup">
                        <thead>
                        <tr>
                            <th data-field="org_name" data-formatter="orgGroupNameFormatter">{{i18n .Lang "operating info"}}</th>
                            <th data-field="count" data-sortable="true" data-formatter="numberFormatter" data-align="right">사용 횟수</th>
                            <th data-field="optime" data-sortable="true" data-formatter="numberFormatter" data-align="right">{{i18n .Lang "operating time"}}</th>
                        </tr>
                        </thead>
                    </table>
                </div>
            </div>
        </div>
    </div>
</div>

<div class="row">
    <div class="col-lg-6 col-md-6">
        <div class="portlet light bordered">
            <div class="portlet-body pt0 mh">
                <div class="row">
                    <div class="col-sm-4">
                        <div class="nbox grey">
                            <div class="number">
                                <a href="#" class="btn-show-ipasstatus-on-modal font-grey-mint" data-query="equip_type={{.PedestrianTag}}">
                                    <span class="count-pt"></span>
                                </a>
                            </div>
                            <div class="desc">
                                <div>PT</div>
                                <div><span class="uppercase s10">Pedestrian Tags</span></div>
                            </div>
                        </div>
                    </div>
                    <div class="col-sm-4">
                        <div class="nbox grey">
                            <div class="number">
                                <a href="#" class="btn-show-ipasstatus-on-modal font-grey-mint" data-query="equip_type={{.VehicleTag}}">
                                    <span class="count-vt"></span>
                                </a>
                            </div>
                            <div class="desc">
                                <div>VT</div>
                                <div><span class="uppercase s10">Vehicle Tags</span></div>
                            </div>
                        </div>
                    </div>
                    <div class="col-sm-4">
                        <div class="nbox grey">
                            <div class="number">
                                <a href="#" class="btn-show-ipasstatus-on-modal font-grey-mint" data-query="equip_type={{.ZoneTag}}">
                                    <span class="count-zt"></span>
                                </a>
                            </div>
                            <div class="desc">
                                <div>ZT</div>
                                <div><span class="uppercase s10">Zone Tags</span></div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="col-lg-6 col-md-6">
        <div class="portlet light bordered">
            <div class="portlet-body pt0">
                <div class="row">
                    <div class="col-sm-4">
                        <div class="nbox blue">
                            <div class="number">
                                <a href="#" class="btn-show-ipaslog-on-modal font-grey-mint" data-query="stats_mode=1&event_type={{.ShockEvent}}">
                                    <span class="count-shock"></span>
                                </a>
                            </div>
                            <div class="desc">
                                <div>{{i18n .Lang "monitoring.shock event"}}</div>
                                {{if eq .Lang  "ko-kr"}}<div><span class="uppercase s10">Shock events</span></div>{{end}}
                            </div>
                        </div>
                    </div>
                    <div class="col-sm-4">
                        <div class="nbox yellow">
                            <div class="number">
                                <a href="#" class="btn-show-ipaslog-on-modal font-grey-mint" data-query="stats_mode=1&event_type={{.SpeedingEvent}}">
                                    <span class="count-speeding"></span>
                                </a>
                            </div>
                            <div class="desc">
                                <div>{{i18n .Lang "monitoring.speeding event"}}</div>
                                {{if eq .Lang  "ko-kr"}}<div><span class="uppercase s10">Speeding events</span></div>{{end}}
                            </div>
                        </div>
                    </div>
                    <div class="col-sm-4">
                        <div class="nbox red">
                            <div class="number">
                                <a href="#" class="btn-show-ipaslog-on-modal font-grey-mint" data-query="stats_mode=1&event_type={{.ProximityEvent}}">
                                    <span class="count-proximity"></span>
                                </a>
                            </div>
                            <div class="desc">
                                <div>{{i18n .Lang "monitoring.proximity event"}}</div>
                                {{if eq .Lang  "ko-kr"}}<div><span class="uppercase s10">Proximity events</span></div>{{end}}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<div class="row">
    <div class="col-lg-3 col-md-6">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="fa fa-bolt"></i> {{i18n .Lang "monitoring.shock event"}}
                </div>
                <div class="pull-right">
                    <button class="btn default btn-xs mt5">{{i18n .Lang "tag"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table-condensed table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-query="/stats/evt2/by/equip">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="15%" data-formatter="rankFormatter"></th>
                        <th data-field="item" data-formatter="dashboardIpasEquipIdOfShockEventFormatter"></th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right"></th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-lg-3 col-md-6">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="fa fa-bolt"></i> {{i18n .Lang "monitoring.shock event"}}
                </div>
                <div class="pull-right">
                    <button class="btn default btn-xs">{{i18n .Lang "group"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table-condensed table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-query="/stats/evt2/by/group">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="15%" data-formatter="rankFormatter"></th>
                        <th data-field="org_name" data-formatter="dashboardOrgGroupNameOfShockEventFormatter"></th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right"></th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-lg-6 col-md-12">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-info"></i> {{i18n .Lang "real-time activities"}}
                </div>
                <div class="pull-right">
                    <div class="mt-checkbox-inline p0">
                        <label class="mt-checkbox mt-checkbox-outline hide">
                            <input type="checkbox" class="activity" name="startupEvent" value="1"> {{i18n .Lang "startup"}}
                            <span></span>
                        </label>
                        <label class="mt-checkbox mt-checkbox-outline">
                            <input type="checkbox" class="activity" name="shockEvent" value="2" checked="checked"> {{i18n .Lang "shock"}}
                            <span></span>
                        </label>
                        <label class="mt-checkbox mt-checkbox-outline">
                            <input type="checkbox" class="activity" name="speedingEvent" value="3" checked="checked"> {{i18n .Lang "speeding"}}
                            <span></span>
                        </label>
                        <label class="mt-checkbox mt-checkbox-outline">
                            <input type="checkbox" class="activity" name="proximityEvent " value="4" checked="checked"> {{i18n .Lang "proximity"}}
                            <span></span>
                        </label>
                    </div>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <div class="scroller" style="height: 150px;" data-always-visible="1" data-rail-visible="0">
                    <ul class="feeds">
                        <table id="table-ipaslogs" class="table"
                               data-classes="table-condensed table-no-bordered"
                               data-toggle="table"
                               data-cache="false"
                               data-show-header="false">
                            <thead>
                            <tr>
                                <th data-field="description" data-formatter="dashboardEventDescriptionFormatter" data-sortable="true">{{i18n .Lang "description"}}</th>
                                <th data-field="date_ago" data-formatter="dashboardDateAgoFormatter" data-align="right"></th>
                            </tr>
                            </thead>
                        </table>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</div>


<div class="row">
    <div class="col-lg-3 col-md-6">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-speedometer"></i> {{i18n .Lang "monitoring.speeding event"}}
                </div>
                <div class="pull-right">
                    <button class="btn default btn-xs">{{i18n .Lang "tag"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table-condensed table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-query="/stats/evt3/by/equip">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="15%" data-formatter="rankFormatter"></th>
                        <th data-field="item" data-formatter="dashboardIpasEquipIdOfSpeedingEventFormatter"></th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right"></th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>
    <div class="col-lg-3 col-md-6">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-speedometer"></i> {{i18n .Lang "monitoring.speeding event"}}
                </div>
                <div class="pull-right">
                    <button class="btn default btn-xs">{{i18n .Lang "group"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table-condensed table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-query="/stats/evt3/by/group">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="15%" data-formatter="rankFormatter"></th>
                        <th data-field="org_name" data-formatter="dashboardOrgGroupNameOfSpeedingEventFormatter"></th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right"></th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-lg-3 col-md-6">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-size-actual"></i> {{i18n .Lang "monitoring.proximity event"}}
                </div>
                <div class="pull-right">
                    <button class="btn default btn-xs">{{i18n .Lang "tag"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table-condensed table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-query="/stats/evt4/by/equip">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="15%" data-formatter="rankFormatter"></th>
                        <th data-field="item" data-formatter="dashboardIpasEquipIdOfProximityEventFormatter" data-event-type="xx"></th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right"></th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-lg-3 col-md-6">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-size-actual"></i> {{i18n .Lang "monitoring.proximity event"}}
                </div>
                <div class="pull-right">
                    <button class="btn default btn-xs">{{i18n .Lang "group"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table-condensed table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-query="/stats/evt4/by/group">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="15%" data-formatter="rankFormatter"></th>
                        <th data-field="org_name" data-formatter="dashboardOrgGroupNameOfProximityEventFormatter"></th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right"></th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>
</div>
{{end}}

{{define "javascript"}}
    {{template "ipaslog/ipas_log_on_modal.tpl" .}}
    {{template "ipasstatus/ipas_status_on_modal.tpl" .}}
    {{template "ipasreport/ipasreport.tpl" .}}
    <script src="/static/plugins/morris.js/morris.min.js"></script>
    <script src="/static/plugins/raphael/raphael.min.js"></script>
    <script src="/static/modules/{{.ctrl}}/dashboard.js"></script>
    <script src="/static/modules/{{.ctrl}}/formatter.js"></script>
{{end}}
