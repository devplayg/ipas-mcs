{{template "base.tpl" .}}

{{define "contents"}}
<form id="form-config">
    {{ .xsrfdata }}

    <div class="table-custom">
        <ul class="nav nav-tabs">
            <li class="active">
                <a href="#tab-system" data-toggle="tab"> System </a>
            </li>
            <li>
                <a href="#tab-members" data-toggle="tab"> Members </a>
            </li>
        </ul>

        <div class="tab-content">
            <div class="tab-pane fade active in" id="tab-system">
                <div class="row">
                    <div class="col-md-4 col-xs-6">
                        <div class="form-group">
                            <label class="control-label">Data retention days</label>
                            <input type="text" class="form-control mask-0999" name="data_retention_days" />
                        </div>
                        <div class="form-group">
                            <label class="control-label">Show IP card</label>
                            <div>
                                <input type="checkbox" class="form-control make-switch" name="use_namecard" data-size="small" data-on-color="success" />
                            </div>
                        </div>
                        <div class="form-group">
                            <label class="control-label">Allow multiple login</label>
                            <div>
                                <input type="checkbox" class="form-control make-switch" name="allow_multiple_login" data-size="small" data-on-color="success" />
                            </div>
                        </div>
                    </div>
                </div>
            </div> <!-- #tab-system -->

            <div class="tab-pane fade" id="tab-members">
                <div class="row">
                    <div class="col-lg-3 col-md-3 col-sm-3 col-xs-6 form-group">
                        <label class="control-label">Max failed login attempts</label>
                        <input type="text" class="form-control mask-099" name="max_failed_login_attempts">
                    </div>
                </div>
            </div> <!-- #tab-members -->
        </div>
    </div>
    <button type="submit" class="btn btn-primary">{{i18n .Lang "save"}}</button>
</form>
{{end}}

{{define "javascript"}}
<!-- Module -->
<script>
    var config = {{.config}};
</script>
<script src="/static/modules/{{.ctrl}}/{{.ctrl}}.js"></script>
{{end}}
