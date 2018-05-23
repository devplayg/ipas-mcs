{{template "base.tpl" .}}

{{define "css"}}
<link href="/static/modules/{{.ctrl}}/dashboard.css" rel="stylesheet" type="text/css" />
<link href="/static/plugins/morris.js/morris.css" rel="stylesheet" type="text/css" />
{{end}}

{{define "contents"}}
<div class="portlet light bordered pt0 pb10 ">
    <div class="portlet-body">
        <form id="form-filter" role="form" method="post">
            <div class="form-body">
                <div class="form-inline">
                    <div class="form-group">
                        <!-- 검색 시작날짜 -->
                        <div class="input-group date datetime hide" data-date="1979-09-16T05:25:07Z" data-date-format="yyyy-mm-dd HH:ii" data-link-field="">
                            <input id="start_date" class="form-control mask-yyyymmddhhii" size="16" type="text" name="start_date" value="{{.filter.StartDate}}" disabled>
                            <span class="input-group-addon"><i class="glyphicon glyphicon-th"></i></span>
                        </div>

                        <!-- 검색 끝날짜 -->
                        <div class="input-group date datetime hide">
                            <input id="end_date" class="form-control mask-yyyymmddhhii" size="16" type="text" name="end_date" value="{{.filter.EndDate}}" disabled>
                            <span class="input-group-addon"><i class="glyphicon glyphicon-th"></i></span>
                        </div>

                        <select id="select-assets" name="org_id" class="selectpicker" data-size="10" data-selected-text-format="count > 2">
                            <option value="-1/-1">{{i18n .Lang "select all"}}</option>
                        </select>

                        <button type="button" class="btn default btn-start"><span class="text"><i class="fa fa-play"></i></span> {{i18n .Lang "monitoring"}}</button>
                        <span class="text-updated font-red hide ml10">Updated</span>
                    </div>
                </div>
            </div>
        </form>
    </div>
</div>

<div class="row">
    <div class="col-md-3">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-grid"></i> {{i18n .Lang "summary"}}
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <div class="row">
                    <div class="col-xs-6 text-center">
                        <h1 class="count-total-tags mt0 mb5 font-grey-mint bold"></h1>
                        Total
                    </div>
                    <div class="col-xs-6 text-center" style="border-left: 1px dashed #acacac;">
                        <h1 class="count-pt mt0 mb5 grey-salsa"></h1>
                        Pedestrian Tag
                    </div>
                </div>
                <div class="row" style="border-top: 1px dashed #acacac; margin: 10px 0px 5px 0px;"></div>
                <div class="row" style="margin: 10px 0px 0px 0px;">
                    <div class="col-xs-6 text-center">
                        <h1 class="count-zt mt0 mb5"></h1>
                        Zone tag
                    </div>
                    <div class="col-xs-6 text-center" style="border-left: 1px dashed #acacac;">
                        <h1 class="count-vt mt0 mb5"></h1>
                        Vehicle tag
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-grid"></i> {{i18n .Lang "event type"}}
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <div id="chart-eventType" class="p0" style="height: 180px; margin: -100px 0px -50px 0px;"></div>
                {{/*<p>Startup: <span class="count-startup"></span></p>*/}}
                {{/*<p>Shock: <span class="count-shock"></span></p>*/}}
                {{/*<p>Speeding: <span class="count-speeding"></span></p>*/}}
                {{/*<p>Proximity: <span class="count-proximity"></span></p>*/}}
            </div>
        </div>
    </div>

    <div class="col-md-6">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-info"></i> {{i18n .Lang "activities"}}
                </div>
                <div class="pull-right">
                    <div class="mt-checkbox-inline p0">
                        <label class="mt-checkbox mt-checkbox-outline">
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
    <div class="col-md-3">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-power"></i> {{i18n .Lang "monitoring.startup event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "group"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table-condensed table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-query="/stats/evt1/by/group">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="15%" data-formatter="rankFormatter"></th>
                        <th data-field="org_name" data-formatter="dashboardOrgGroupNameOfStartupEventFormatter"></th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right"></th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="fa fa-bolt"></i> {{i18n .Lang "monitoring.shock event"}}
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

    <div class="col-md-3">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-speedometer"></i> {{i18n .Lang "monitoring.speeding event"}}
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

    <div class="col-md-3">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-size-actual"></i> {{i18n .Lang "monitoring.proximity event"}}
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


<div class="row">

    <div class="col-md-3">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-power"></i> {{i18n .Lang "monitoring.startup event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "tag"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table-condensed table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-query="/stats/evt1/by/equip">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="15%" data-formatter="rankFormatter"></th>
                        <th data-field="item" data-formatter="dashboardIpasEquipIdOfStartupEventFormatter"></th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right"></th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="fa fa-bolt"></i> {{i18n .Lang "monitoring.shock event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "tag"}} Top</button>
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

    <div class="col-md-3">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-speedometer"></i> {{i18n .Lang "monitoring.speeding event"}}
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

    <div class="col-md-3">
        <div class="portlet light bordered mh-rank">
            <div class="portlet-title">
                <div class="caption">
                    <i class="icon-size-actual"></i> {{i18n .Lang "monitoring.proximity event"}}
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
</div>


{{end}}



{{define "javascript"}}
    {{template "ipaslog/ipas_log_on_modal.tpl" .}}
    {{template "ipasreport/ipasreport.tpl" .}}
    <script src="/static/plugins/morris.js/morris.min.js"></script>
    <script src="/static/plugins/raphael/raphael.min.js"></script>
    <script src="/static/modules/{{.ctrl}}/dashboard.js"></script>
    <script src="/static/modules/{{.ctrl}}/formatter.js"></script>
{{end}}
