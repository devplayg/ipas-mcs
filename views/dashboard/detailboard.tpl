{{template "base.tpl" .}}

{{define "css"}}
<link href="/static/modules/{{.ctrl}}/detailboard.css" rel="stylesheet" type="text/css" />
<link href="/static/plugins/morris.js/morris.css" rel="stylesheet" type="text/css" />
{{end}}

{{define "contents"}}
<div class="row">
    <div class="col-md-4">
        <div class="panel panel-default mh">
            <div class="panel-body">
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

                <div class="row mt15">
                    <div class="col-xs-6 text-center">
                        <h1 class="count-total-tags mt0 font-grey-mint bold"></h1>
                        Total
                    </div>
                    <div class="col-xs-6 text-center" style="border-left: 1px dashed #acacac;">
                        <h1 class="count-pt mt0 grey-salsa"></h1>
                        Pedestrian Tag
                    </div>
                </div>
                <div class="row" style="border-top: 1px dashed #acacac; margin: 5px 0px 15px 0px;"></div>
                <div class="row" style="margin: 15px 0px 0px 0px;">
                    <div class="col-xs-6 text-center">
                        <h1 class="count-zt mt0"></h1>
                        Zone tag
                    </div>
                    <div class="col-xs-6 text-center" style="border-left: 1px dashed #acacac;">
                        <h1 class="count-vt mt0"></h1>
                        Vehicle tag
                    </div>
                </div>
            </div>
        </div>


        <div class="panel panel-default">
            <div class="panel-heading bg-white">
                <h3 class="panel-title s14 bold">{{i18n .Lang "event type"}}</h3>
            </div>
            <div class="panel-body" id="panel-shocklinks">
                <table id="table-activated" class="table"
                       data-classes="table-condensed table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-height="260"
                       data-sort-name="count"
                       data-sort-order="desc"
                       data-show-header="true">
                    <thead>
                    <tr>
                        <th data-field="org_name" data-sortable="true">{{i18n .Lang "org"}}</th>
                        <th data-field="group_name" data-sortable="true">{{i18n .Lang "group"}}</th>
                        <th data-field="count" data-align="right" data-sortable="true">{{i18n .Lang "count"}}</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-md-8">
        <div class="panel panel-default mh">
            <div class="panel-heading bg-white">
                <h3 class="panel-title s14 bold">{{i18n .Lang "event type"}}</h3>
            </div>
            <div class="panel-body">
                <div class="row">
                    <div class="col-lg-4">
                        <div id="chart-eventType" class="clear" style="height: 180px; width: 180px; margin-top: -10px;"></div>
                        <div class="progress-info mb15 mt15 clearfix">
                            <div class="progress mb0">
                                <span id="pgb-shock" style="width: 0%;" class="progress-bar progress-bar-info blue"></span>
                            </div>
                            <div class="status">
                                <div class="status-title pull-left">{{i18n .Lang "shock"}}</div>
                                <div class="status-number pull-right count-shock"></div>
                            </div>
                        </div>
                        <div class="progress-info mb15 clearfix">
                            <div class="progress mb0">
                                <span id="pgb-speeding" style="width: 0%;" class="progress-bar progress-bar-success green"></span>
                            </div>
                            <div class="status">
                                <div class="status-title pull-left">{{i18n .Lang "speeding"}}</div>
                                <div class="status-number pull-right count-speeding"></div>
                            </div>
                        </div>
                        <div class="progress-info">
                            <div class="progress mb0">
                                <span id="pgb-proximity" style="width: 0%;" class="progress-bar progress-bar-success red-haze"></span>
                            </div>
                            <div class="status">
                                <div class="status-title pull-left">{{i18n .Lang "proximity"}}</div>
                                <div class="status-number pull-right count-proximity"></div>
                            </div>
                        </div>
                    </div>
                    <div class="col-lg-8">
                        <table id="table-ipaslogs" class="table"
                               data-classes="table-condensed table-no-bordered"
                               data-toggle="table"
                               data-cache="false"
                               data-height="330"
                               data-show-header="false">
                            <thead>
                            <tr>
                                <th data-field="description" data-formatter="dashboardEventDescriptionFormatter" data-sortable="true">{{i18n .Lang "description"}}</th>
                                <th data-field="date_ago" data-formatter="dashboardDateAgoFormatter" data-align="right"></th>
                            </tr>
                            </thead>
                        </table>
                    </div>
                </div>
            </div>
        </div>

        <div class="panel panel-default">
            <div class="panel-heading bg-white">
                <h3 class="panel-title s14 bold">{{i18n .Lang "event type"}}</h3>
            </div>
            <div class="panel-body">
                <div id="chart-trend" style="height: 180px;"></div>
            </div>
        </div>
    </div>
</div>


<div class="row">
    <div class="col-md-4">

    </div>


</div>

<div class="row">
    <div class="col-md-6">



    </div>

    <div class="col-md-6">
        <div class="panel panel-default">
            <div class="panel-heading bg-white">
                <h3 class="panel-title s14 bold">{{i18n .Lang "event type"}}</h3>
            </div>
            <div class="panel-body text-center" id="panel-shocklinks" style="height: 500px;">
                <div id="chart-shocklinks" class="p0" style="margin: -50px 0px 0px -0px"></div>
            </div>
        </div>

    </div>
</div>

<div class="row">
    <div class="col-md-4">

    </div>
    <div class="col-md-8">
    </div>
</div>

{{end}}



{{define "javascript"}}
    {{template "ipaslog/ipas_log_on_modal.tpl" .}}
    {{template "ipasreport/ipasreport.tpl" .}}
    <script src="/static/plugins/morris.js/morris.min.js"></script>
    <script src="/static/plugins/raphael/raphael.min.js"></script>
    <script src="/static/plugins/echarts/echarts.min.js"></script>
    <script src="/static/plugins/d3/d3.v4.min.js"></script>
    <script src="/static/modules/{{.ctrl}}/detailboard.js"></script>
    <script src="/static/modules/{{.ctrl}}/formatter.js"></script>
{{end}}
