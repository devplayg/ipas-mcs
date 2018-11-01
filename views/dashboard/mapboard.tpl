{{template "base.tpl" .}}

{{define "css"}}
{{end}}

{{define "contents"}}
<div class="row">
    <div class="col-lg-8">
        <div class="portlet light bordered">
            <div class="portlet-body pt0 mh">
                <div id="map"></div>
            </div>
        </div>
    </div>
    <div class="col-lg-4">
        <div class="portlet light bordered">
            <div class="portlet-body form-group">
                <!-- 자산 정보 -->
                <select id="select-assets" name="org_id" class="selectpicker form-control" data-size="10" data-selected-text-format="count > 2">
                    <option value="-1/-1">{{i18n .Lang "total assets"}}</option>
                </select>
            </div>

            <div class="portlet-body form-group">
                <button class="btn btn-default btn-sm btn-period" data-period="0">Today</button>
                <button class="btn btn-default btn-sm btn-period" data-period="1">2d</button>
                <button class="btn btn-default btn-sm btn-period" data-period="2">3d</button>
                <button class="btn btn-default btn-sm btn-period" data-period="6">1w</button>
                <button class="btn btn-default btn-sm btn-period" data-period="13">2w</button>
                <button class="btn btn-default btn-sm btn-period" data-period="1m">1m</button>
            </div>
        </div>
        <div class="portlet light bordered">
            <div class="row">
                <div class="col-lg-6">
                    <i class="icon-grid"></i> {{i18n .Lang "event type"}}
                    <div id="chart-eventType" class="mt10" style="height: 180px; width: 180px;"></div>
                </div>
                <div class="col-lg-6" style="padding-top: 40px;">

                    <div class="">
                        <div class="progress mb5">
                            <span id="pgb-shock" class="progress-bar progress-bar-info" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100"></span>
                        </div>
                        <div class="s12 ">
                        {{i18n .Lang "shock"}} (<span class="count-shock"></span>)
                            <span class="rate-shock pull-right">%</span>
                        </div>
                    </div>

                    <div class=" mt10">
                        <div class="progress mb5">
                            <span id="pgb-speeding" class="progress-bar progress-bar-warning" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100"></span>
                        </div>
                        <div class="s12 ">
                        {{i18n .Lang "speeding"}} (<span class="count-speeding"></span>)
                            <span class="rate-speeding pull-right">%</span>
                        </div>
                    </div>

                    <div class=" mt10">
                        <div class="progress mb5">
                            <span id="pgb-proximity" class="progress-bar progress-bar-danger" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100"></span>
                        </div>
                        <div class="s12 ">
                        {{i18n .Lang "proximity"}} (<span class="count-proximity"></span>)
                            <span class="rate-proximity pull-right">%</span>
                        </div>
                    </div>

                </div>
            </div>
        </div>
    </div>
</div>
{{end}}

{{define "javascript"}}
    <script src="/static/plugins/morris.js/morris.min.js"></script>
    <script src="/static/plugins/raphael/raphael.min.js"></script>
    <script src="/static/modules/{{.ctrl}}/mapboard.js"></script>
    <script src="/static/modules/{{.ctrl}}/formatter.js"></script>
{{end}}
