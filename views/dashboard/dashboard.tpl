{{template "base.tpl" .}}

{{define "css"}}
<link href="/static/modules/{{.ctrl}}/dashboard.css" rel="stylesheet" type="text/css" />
{{end}}

{{define "contents"}}
<div class="portlet light bordered pt0">
    <div class="portlet-body pt0 mh">
        <select id="select-orgs" name="org_id" class="selectpicker" multiple title="{{i18n .Lang "org"}}"  data-size="10" data-selected-text-format="count > 2"></select>
    </div>
</div>

<div class="row">

    <div class="col-md-3">
        <div class="portlet light bordered">
            <div class="portlet-title">
                <div class="caption">
                    {{i18n .Lang "monitoring.startup event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "group"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-url="/stats/evt1/by/group">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="25%" data-formatter="rankFormatter"></th>
                        <th data-field="item_text">Item</th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right">ccc</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="portlet light bordered">
            <div class="portlet-title">
                <div class="caption">
                {{i18n .Lang "monitoring.shock event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "group"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-url="/stats/evt2/by/group">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="25%" data-formatter="rankFormatter"></th>
                        <th data-field="item_text">Item</th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right">ccc</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="portlet light bordered">
            <div class="portlet-title">
                <div class="caption">
                {{i18n .Lang "monitoring.speeding event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "group"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-url="/stats/evt3/by/group">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="25%" data-formatter="rankFormatter"></th>
                        <th data-field="item_text">Item</th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right">ccc</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="portlet light bordered">
            <div class="portlet-title">
                <div class="caption">
                {{i18n .Lang "monitoring.proximity event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "group"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-url="/stats/evt4/by/group">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="25%" data-formatter="rankFormatter"></th>
                        <th data-field="item_text">Item</th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right">ccc</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

</div>


<div class="row">

    <div class="col-md-3">
        <div class="portlet light bordered">
            <div class="portlet-title">
                <div class="caption">
                {{i18n .Lang "monitoring.startup event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "tag"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-url="/stats/evt1/by/equip">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="25%" data-formatter="rankFormatter"></th>
                        <th data-field="item">Item</th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right">ccc</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="portlet light bordered">
            <div class="portlet-title">
                <div class="caption">
                {{i18n .Lang "monitoring.shock event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "tag"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-url="/stats/evt2/by/equip">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="25%" data-formatter="rankFormatter"></th>
                        <th data-field="item">Item</th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right">ccc</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="portlet light bordered">
            <div class="portlet-title">
                <div class="caption">
                {{i18n .Lang "monitoring.speeding event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "tag"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-url="/stats/evt3/by/equip">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="25%" data-formatter="rankFormatter"></th>
                        <th data-field="item">Item</th>
                        <th data-field="count" data-formatter="numberFormatter" data-align="right">ccc</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="portlet light bordered">
            <div class="portlet-title">
                <div class="caption">
                {{i18n .Lang "monitoring.proximity event"}}
                    <button class="btn default btn-xs">{{i18n .Lang "tag"}} Top</button>
                </div>
            </div>
            <div class="portlet-body pt0 mh">
                <table class="table table-ranking"
                       data-classes="table table-no-bordered"
                       data-toggle="table"
                       data-cache="false"
                       data-show-header="false"
                       data-url="/stats/evt4/by/equip">
                    <thead>
                    <tr>
                        <th data-field="rank" data-width="25%" data-formatter="rankFormatter"></th>
                        <th data-field="item">Item</th>
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
<script src="/static/modules/{{.ctrl}}/formatter.js"></script>
{{end}}