{{template "base.tpl" .}}

{{define "css"}}
    <link href="/static/plugins/jstree/dist/themes/default/style.min.css" rel="stylesheet" type="text/css" />
{{end}}

{{define "contents"}}
<div class="row">
    <div class="col-lg-3">
        <div class="portlet light bordered">
            <div class="portlet-title hidden">
                <div class="caption font-green-sharp">
                    <i class="icon-speech font-green-sharp"></i>
                    <span class="caption-subject bold uppercase"> Assets</span>
                    <span class="caption-helper"></span>
                </div>
            </div> <!-- .portlet-title -->
            <div class="portlet-body mb5">
                <div class="mb10">
                    <button type="button" class="btn default btn-xs btn-tree-refresh"><i class="fa fa-refresh"></i> </button>
                    <button type="button" class="btn default btn-xs btn-tree-expand">Expand </button>
                    <button type="button" class="btn default btn-xs btn-tree-collapse">Collapse </button>
                    <button type="button" class="btn default btn-xs btn-asset-add" data-toggle="modal" data-target="#modal-asset-add"><i class="fa fa-plus"></i></button>
                </div>
                <div id="tree-assets"></div>
            </div> <!-- .portlet-body -->
        </div> <!-- .portlet -->
    </div>

    <div class="col-lg-9">
        <div class="portlet light bordered">
            <table class=""
                   id="table-assets"
                   data-toggle="table"
                   data-toolbar="#toolbar"
                   data-page-size="15"
                   data-pagination="true"
                   data-sort-order="desc"
                   data-show-refresh="true"
                   data-show-columns="true"
            >
                <thead>
                <tr>
                    <th data-field="AssetId" data-sortable="true">Code</th>
                    <th data-field="Name" data-sortable="true" data-formatter="assetsNameFormatter">Name</th>
                    <th data-field="Hostname" data-sortable="true" data-formatter="assetsHostnameFormatter">Host</th>
                </tr>
                </thead>
            </table>
        </div> <!-- .portlet -->

    </div>
</div>

<div id="modal-asset-add" class="modal fade" tabindex="-1" role="basic" aria-hidden="true">
    <div class="modal-dialog modal-sm">
        <div class="modal-content">
            <form role="form" id="form-asset-add">
                <input type="hidden" name="_xsrf" value="{{ .xsrf_token }}" />
                <input type="hidden" name="Type1" />

                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-hidden="true"></button>
                    <h4 class="modal-title">{{i18n .Lang "registration"}}</h4>
                </div>
                <div class="modal-body">
                    <div class="row">
                        <div class="col-lg-12 form-group">
                            <label for="Name" class="control-label">Name</label>
                            <input type="text" class="form-control" name="Name" />
                        </div>
                    </div>
                    <div class="note note-danger hidden"></div>
                </div>
                <div class="modal-footer">
                    <button type="submit" class="btn blue">{{i18n .Lang "registration"}}</button>
                    <button type="button" class="btn dark btn-outline" data-dismiss="modal">{{i18n .Lang "close"}}</button>
                </div>
            </form>
        </div> <!-- .modal-content -->
    </div> <!-- .modal-dialog -->
</div> <!-- .modal -->

{{end}}

{{define "javascript"}}
<!-- jstree -->
<script src="/static/plugins/jstree/dist/jstree.js" type="text/javascript"></script>
<!-- Module -->
<script src="/static/modules/{{.ctrl}}/formatter.js"></script>
<script src="/static/modules/{{.ctrl}}/{{.ctrl}}.js"></script>
{{end}}
