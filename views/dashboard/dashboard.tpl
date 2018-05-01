{{template "base.tpl" .}}

{{define "contents"}}
<div class="row">
    <div class="col-sm-4">
        <div class="portlet light bordered pt0">
            <div class="portlet-body pt0 mh">

            </div>
        </div>
    </div>
</div>

<div class="row">
    <div class="col-sm-3">
        <div class="portlet light bordered pt0">
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-url="/rank/startup"
                       data-classes="table table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="25%" data-formatter="comparedRankFormatter"></th>
                        <th data-field="tag" data-formatter="">bbb</th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right">ccc</th>
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
<script src="/static/modules/{{.ctrl}}/dashboard.js"></script>
{{end}}