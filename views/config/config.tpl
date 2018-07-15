{{template "base.tpl" .}}

{{define "contents"}}
<div class="portlet light bordered pt0">
    <div class="portlet-body">
        <form id="form-config">
            {{ .xsrfdata }}
            <div class="row">
                <div class="col-xs-4">
                    <h4 class="form-section">{{i18n .Lang "system"}}</h4>
                    <hr>
                    <div class="form-group">
                        <label class="control-label">{{i18n .Lang "data retention period"}}</label>
                            <input type="text" class="form-control mask-0999" name="data_retention_days" value="{{.system_data_retention_days.ValueN}}" />
                    </div>
                </div>
                <div class="col-xs-offset-2 col-xs-4">
                    <h4 class="form-section">{{i18n .Lang "signin"}}</h4>
                    <hr>

                    <div class="form-group">
                        <label class="control-label">{{i18n .Lang "max failed login attempts"}}</label>
                        <input type="text" class="form-control mask-099" name="max_failed_login_attempts" value="{{.login_max_failed_login_attempts.ValueN}}">
                    </div>
                    <div class="form-group">
                        <label class="control-label">{{i18n .Lang "login failure block time"}}</label>
                        <input type="text" class="form-control mask-099" name="login_failure_block_time" value="{{.login_failure_block_time.ValueN}}">
                    </div>
                </div>
            </div>
            <button type="submit" class="btn btn-primary">{{i18n .Lang "save"}}</button>
        </form>
    </div>
</div>
{{end}}

{{define "javascript"}}
    <!-- Module -->
    <script src="/static/modules/{{.ctrl}}/{{.ctrl}}.js"></script>
{{end}}
