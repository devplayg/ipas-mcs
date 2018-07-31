{{template "base-blank.tpl" .}}

{{define "css"}}
<!-- Flags -->
<link href="/static/plugins/flag-icon-css/css/flag-icon.min.css" rel="stylesheet">
<link href="/static/modules/login/login.css" rel="stylesheet" type="text/css" />
{{end}}

{{define "contents"}}
<div class="login pt20">
<!-- BEGIN LOGO -->
<!--
<div class="logo">
    <a href="index.html">
        <img src="/static/assets/pages/img/logo-big.png" alt="" /> </a>
</div>
-->
<!-- END LOGO -->


<!-- BEGIN LOGIN -->

    <div class="content">
        <form id="form-login" class="login-form" role="form" data-async autocomplete="off" >
            <input type="hidden" name="redirectUri" value="{{.redirectUri}}" />

            <h3 class="form-title font-green">Sign In</h3>
            <div class="form-group">
                <!--ie8, ie9 does not support html5 placeholder, so we just show field title for that-->
                <label class="control-label visible-ie8 visible-ie9">{{i18n .Lang "username"}}</label>
                <input class="form-control form-control-solid placeholder-no-fix" type="text" autocomplete="off" placeholder="{{i18n .Lang "username"}}" name="username" /> </div>
            <div class="form-group">
                <label class="control-label visible-ie8 visible-ie9">{{i18n .Lang "password"}}</label>
                <input class="form-control form-control-solid placeholder-no-fix" type="password" autocomplete="off" placeholder="{{i18n .Lang "password"}}" name="password" /> </div>

            <div class="note note-danger hidden"></div>

            <div class="form-actions">
                <button type="submit" class="btn green uppercase">{{i18n .Lang "signin"}}</button>
                <div class="pull-right mt10 font-lg">
                    <a href="javascript:;" data-lang="en-us" class="lang-changed"><span class="flag-icon flag-icon-us flag-border"></span></a>
                    <a href="javascript:;" data-lang="ja-jp" class="lang-changed hide"><span class="flag-icon flag-icon-jp flag-border"></span></a>
                    <a href="javascript:;" data-lang="ko-kr" class="lang-changed"><span class="flag-icon flag-icon-kr flag-border"></span></a>
                </div>
                <!--
                <label class="rememberme check mt-checkbox mt-checkbox-outline">
                    <input type="checkbox" name="remember" value="1" />Remember
                    <span></span>
                </label>
                <a href="javascript:;" id="forget-password" class="forget-password">Forgot Password?</a>
                -->
            </div>

            <div class="text-center">
                <h4 class="mt20">{{.title}}</h4>
                <small>IPAS Monitoring & Control System</small>
            </div>
        </form>
        <!-- END LOGIN FORM -->
    </div>
    <div class="copyright">
        <a href="http://www.kyungwoo.com/" class="font-grey-cararra" target="_blank">&copy; 2018 {{ .company_name }} Inc.</a>
    </div>

</div>
{{end}}

{{define "javascript"}}
<!-- CryptoJS -->
<script src="/static/plugins/crypto/rollups/sha256.js"></script>
<!-- Cookie -->
<script src="/static/assets/js/jquery.cookie.js" type="text/javascript"></script>
<!-- Module -->
<script src="/static/modules/{{.ctrl}}/{{.ctrl}}.js"></script>
{{end}}

