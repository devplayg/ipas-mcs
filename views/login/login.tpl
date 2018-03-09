{{template "base-blank.tpl" .}}

{{define "css"}}
<link href="/static/modules/login/login.css_" rel="stylesheet" type="text/css" />
{{end}}

{{define "contents"}}
<div class="login">
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
                <label class="control-label visible-ie8 visible-ie9">Username</label>
                <input class="form-control form-control-solid placeholder-no-fix" type="text" autocomplete="off" placeholder="Username" name="username" /> </div>
            <div class="form-group">
                <label class="control-label visible-ie8 visible-ie9">Password</label>
                <input class="form-control form-control-solid placeholder-no-fix" type="password" autocomplete="off" placeholder="Password" name="password" /> </div>

            <div class="note note-danger hidden"></div>

            <div class="form-actions">
                <button type="submit" class="btn green uppercase">Login</button>
                <!--
                <label class="rememberme check mt-checkbox mt-checkbox-outline">
                    <input type="checkbox" name="remember" value="1" />Remember
                    <span></span>
                </label>
                <a href="javascript:;" id="forget-password" class="forget-password">Forgot Password?</a>
                -->
            </div>

            <div class="login-options hidden">
                <h4>Or login with</h4>
                <ul class="social-icons">
                    <li>
                        <a class="social-icon-color facebook" data-original-title="facebook" href="javascript:;"></a>
                    </li>
                    <li>
                        <a class="social-icon-color twitter" data-original-title="Twitter" href="javascript:;"></a>
                    </li>
                    <li>
                        <a class="social-icon-color googleplus" data-original-title="Goole Plus" href="javascript:;"></a>
                    </li>
                    <li>
                        <a class="social-icon-color linkedin" data-original-title="Linkedin" href="javascript:;"></a>
                    </li>
                </ul>
            </div>
            <div class="create-account hidden">
                <p>
                    <a href="javascript:;" id="register-btn" class="uppercase">Create an account</a>
                </p>
            </div>
        </form>
        <!-- END LOGIN FORM -->
    </div>

</div>
{{end}}

{{define "javascript"}}
<!-- CryptoJS -->
<script src="/static/plugins/crypto/rollups/sha256.js"></script>
<!-- Module -->
<script src="/static/modules/{{.ctrl}}/{{.ctrl}}.js"></script>
{{end}}

