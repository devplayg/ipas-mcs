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
                        <button type="button" class="btn default btn-xs btn-tree-expand"><i class="fa fa-expand"></i></button>
                        <button type="button" class="btn blue btn-xs btn-asset-manage btn-asset-add hide" data-perm="3"><i class="fa fa-plus"></i> {{i18n .Lang "registration"}}</button>
                        <button type="button" class="btn green btn-xs btn-asset-manage btn-asset-edit hide" data-perm="6"><i class="fa fa-edit"></i> {{i18n .Lang "edit"}}</button>
                        <button type="button" class="btn red btn-xs btn-asset-remove" data-perm="6"><i class="fa fa-trash"></i> {{i18n .Lang "remove"}}</button>
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
                {{.xsrfdata}}
                    <input type="hidden" name="class" value="1">
                    <input type="hidden" name="parent_id">
                    <input type="hidden" name="type1">

                    <div class="modal-header">
                        <h4 class="modal-title">{{i18n .Lang "registration"}} <span class="target pull-right"></span></h4>
                    </div>
                    <div class="modal-body">
                        <div class="form-body">
                            <div class="form-group">
                                <label for="name" class="control-label">{{i18n .Lang "name"}} <span class="name"></span></label>
                                <input type="text" class="form-control" name="name" />
                            </div>
                            <div class="form-group form-option form-option-org hide">
                                <label for="code" class="control-label">{{i18n .Lang "code"}}</label>
                                <input type="text" class="form-control" name="code" />
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

    <div id="modal-asset-edit" class="modal fade" tabindex="-1" role="basic" aria-hidden="true">
        <div class="modal-dialog modal-sm">
            <div class="modal-content">
                <form role="form" id="form-asset-edit">
                    {{.xsrfdata}}
                    <input type="hidden" name="asset_id">

                    <div class="modal-header">
                        <h4 class="modal-title">{{i18n .Lang "edit"}} <span class="target pull-right"></span></h4>
                    </div>
                    <div class="modal-body">
                        <div class="form-body">
                            <div class="form-group">
                                <label for="name" class="control-label">{{i18n .Lang "name"}} <span class="name"></span></label>
                                <input type="text" class="form-control" name="name" />
                            </div>
                            <div class="form-group form-option form-option-org hide">
                                <label for="name" class="control-label">{{i18n .Lang "code"}}</label>
                                <input type="text" class="form-control" name="code" />
                            </div>
                        </div>
                        <div class="note note-danger hidden"></div>
                    </div>
                    <div class="modal-footer">
                        <button type="submit" class="btn green">{{i18n .Lang "edit"}}</button>
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
<script>
    felang[ "org" ] = {{i18n .Lang "org" }};
    felang[ "group" ] = {{i18n .Lang "group" }};
</script>
{{end}}
