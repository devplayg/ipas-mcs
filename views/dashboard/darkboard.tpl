{{template "base.tpl" .}}

{{define "css"}}
<link href="/static/modules/{{.ctrl}}/dashboard.css" rel="stylesheet" type="text/css" />
<link href="/static/modules/{{.ctrl}}/darkboard.css" rel="stylesheet" type="text/css" />
{{end}}

{{define "contents"}}
<div class="row">
    <div class="col-lg-12">
        <div class="portlet light bordered pt0 pb10 mb15">
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

    </div>
</div>

<div class="row">
    <div class="col-lg-4 col-md-5">
        <div class="mt-element-ribbon bg-white">
            <div class="ribbon ribbon-shadow ribbon-color-info uppercase">Tags</div>
            <div class="ribbon-content" style="height: 218px;">
                <div id="chart-tags" class="mh" style="poistion:absolute; top:-50px;"></div>
            </div>
        </div>
    </div>
    <div class="col-lg-8 col-md-7">
        <div class="portlet light bordered pt0 pb10 mb15">
            <div class="portlet-body">
                <div id="chart-trend" class="mh" style="border:1px dashed #ccc;"></div>
            </div>
        </div>
    </div>
</div>

<div class="row">
    <div class="col-md-4">
    </div>
    <div class="col-md-8">
            <div class="row">
                <div class="col-md-4">
                    <div class="mt-element-ribbon bg-grey-salt">
                        <div class="ribbon ribbon-shadow ribbon-color-danger uppercase">{{i18n .Lang "monitoring.shock event"}}
                        </div>
                        <p class="ribbon-content"><span class="count-startup"></span></p>
                    </div>
                </div>
                <div class="col-md-4">
                    <div class="mt-element-ribbon bg-grey-salt">
                        <div class="ribbon ribbon-shadow ribbon-color-info uppercase">{{i18n .Lang "monitoring.speeding event"}}</div>
                        <p class="ribbon-content">Duis mollis, est non commodo luctus, nisi erat porttitor ligula</p>
                    </div>
                </div>
                <div class="col-md-4">
                    <div class="mt-element-ribbon bg-grey-salt">
                        <div class="ribbon ribbon-shadow ribbon-color-success uppercase">{{i18n .Lang "monitoring.proximity event"}}</div>
                        <p class="ribbon-content">Duis mollis, est non commodo luctus, nisi erat porttitor ligula</p>
                    </div>
                </div>
            </div>

            <!--
            <div class="col-md-4 bg-red">
                <div class="step-value uppercase font-white"><span class="count-shock"></span></div>
                <div class="step-title font-white">Purchasing the item</div>
            </div>
            <div class="col-md-4 bg-blue">
                <div class="step-value uppercase font-white "><span class="count-speeding"></span></div>
                <div class="step-title font-white">Complete your payment</div>
            </div>
            <div class="col-md-4 bg-yellow-gold">
                <div class="step-value uppercase font-white"><span class="count-proximity"></span></div>
                <div class="step-title font-white">Receive item integration</div>
            </div>
            -->
        </div>
    </div>
    {{/*<div class="col-md-8">*/}}
        {{/*<div class="row">*/}}
            {{/*<div class="col-md-3">*/}}
                {{/*<div class="panel-box">*/}}
                    {{/*<div class="">*/}}
                        {{/*<span class="count-activated s28"></span>*/}}
                        {{/*<span class="activated s20 mlr5">/</span>*/}}
                        {{/*<span class="count-total-equips s20"></span>*/}}
                    {{/*</div>*/}}
                    {{/*<span class="">Active / Total</span>*/}}
                {{/*</div>*/}}
            {{/*</div>*/}}

            {{/*<div class="col-md-3">*/}}
                {{/*<div class="panel-box">*/}}
                    {{/*<div class="">*/}}
                        {{/*<span class="count-shock s28 bold"></span>*/}}
                    {{/*</div>*/}}
                    {{/*<span class="">{{i18n .Lang "monitoring.shock event"}}</span>*/}}
                {{/*</div>*/}}
            {{/*</div>*/}}
            {{/*<div class="col-md-3">*/}}
                {{/*<div class="panel-box">*/}}
                    {{/*<div class="box">*/}}
                        {{/*<span class="count-startup s28 bold"></span>*/}}
                    {{/*</div>*/}}
                    {{/*<span class="">{{i18n .Lang "monitoring.speeding event"}}</span>*/}}
                {{/*</div>*/}}
            {{/*</div>*/}}
            {{/*<div class="col-md-3">*/}}
                {{/*<div class="panel-box">*/}}
                    {{/*<div class="box">*/}}
                        {{/*<span class="count-proximity s28 bold"></span>*/}}
                    {{/*</div>*/}}
                    {{/*<span class="">{{i18n .Lang "monitoring.proximity event"}}</span>*/}}
                {{/*</div>*/}}
            {{/*</div>*/}}
            {{/*<div class="col-md-3">*/}}

            {{/*</div>*/}}
        {{/*</div>*/}}
    {{/*</div>*/}}
</div>

<div class="row ">
    <div class="col-md-4">
        <div class="">
        </div>
    </div>
    <div class="col-md-4">
    </div>
    <div class="col-md-4">
    </div>
</div>

{{/*
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
-->
*/}}
{{end}}



{{define "javascript"}}
    {{template "ipaslog/ipas_log_on_modal.tpl" .}}
    {{template "ipasreport/ipasreport.tpl" .}}
    <script src="/static/plugins/echarts/echarts.min.js"></script>
    <script src="/static/modules/{{.ctrl}}/darkboard.js"></script>
    <script src="/static/modules/{{.ctrl}}/formatter.js"></script>
{{end}}
