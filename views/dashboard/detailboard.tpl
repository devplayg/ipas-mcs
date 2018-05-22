{{template "base.tpl" .}}

{{define "css"}}
<link href="/static/modules/{{.ctrl}}/detailboard.css" rel="stylesheet" type="text/css" />
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
    <div class="col-md-4">
        <div class="panel panel-default mh">
            <div class="panel-heading">
                <h3 class="panel-title s14 bold">{{i18n .Lang "event type"}}</h3>
            </div>
            <div class="panel-body">
                <div class="row">
                    <div class="col-lg-6">
                        <div id="chart-eventType" style="height: 180px; width: 180px; margin-top: -10px;"></div>
                    </div>
                    <div class="col-lg-6">
                        <div class="progress-info mb15 mt15 clearfix">
                            <div class="progress mb0">
                                <span id="pgb-shock" style="width: 0%;" class="progress-bar progress-bar-info blue"></span>
                            </div>
                            <div class="status">
                                <div class="status-title pull-left"> Shock </div>
                                <div class="status-number pull-right count-shock"></div>
                            </div>
                        </div>
                        <div class="progress-info mb15 clearfix">
                            <div class="progress mb0">
                                <span id="pgb-speeding" style="width: 0%;" class="progress-bar progress-bar-success green"></span>
                            </div>
                            <div class="status">
                                <div class="status-title pull-left"> Speeding </div>
                                <div class="status-number pull-right count-speeding"></div>
                            </div>
                        </div>
                        <div class="progress-info">
                            <div class="progress mb0">
                                <span id="pgb-proximity" style="width: 0%;" class="progress-bar progress-bar-success red-haze"></span>
                            </div>
                            <div class="status">
                                <div class="status-title pull-left"> Proximity </div>
                                <div class="status-number pull-right count-proximity"></div>
                            </div>
                        </div>
                    </div>

                </div>
            </div>
        </div>
    </div>
    <div class="col-md-4">
        <div class="panel panel-default mh">
            <div class="panel-heading">
                <h3 class="panel-title s14 bold">{{i18n .Lang "event type"}}</h3>
            </div>
            <div class="panel-body">

            </div>
        </div>
    </div>

    <div class="col-md-4">
        <div class="panel panel-default mh">
            <div class="panel-heading">
                <h3 class="panel-title s14 bold">{{i18n .Lang "activities"}}</h3>
            </div>
            <div class="panel-body">
                <div id="event-tags" style="line-height: 180%"></div>
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
    <script src="/static/modules/{{.ctrl}}/detailboard.js"></script>
    <script src="/static/modules/{{.ctrl}}/formatter.js"></script>
{{end}}
