<!DOCTYPE html>{{literal `
<!--[if IE 8]> <html lang="en" class="ie8 no-js"> <![endif]-->
<!--[if IE 9]> <html lang="en" class="ie9 no-js"> <![endif]-->
<!--[if !IE]><!-->
<html lang="en">
    <!--<![endif]-->`}}
    <head>
        <meta charset="utf-8" />
        <title>{{.Title}}</title>
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta content="width=device-width, initial-scale=1" name="viewport" />
        <meta content="" name="description" />
        <meta content="" name="author" />
        <meta name="_xsrf" content="{{.xsrf_token}}" />

        <!-- Theme -->
        <link href="/static/plugins/bootstrap/css/bootstrap.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/css/components-rounded.min.css" rel="stylesheet" id="style_components" type="text/css" />
        <link href="/static/assets/css/layout.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/css/themes/darkblue.min.css" rel="stylesheet" type="text/css" id="style_color" />
        <link href="/static/assets/css/plugins.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/css/custom.css" rel="stylesheet" type="text/css" />

        <!-- Fonts -->
        <link href="/static/assets/font/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/font/simple-line-icons/css/simple-line-icons.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/font/{{.Lang}}.css" rel="stylesheet" type="text/css" />

        <!-- Plugins -->
        {{ block "css" . }}{{ end }}
        <link rel="shortcut icon" href="favicon.ico" />
    </head>

    <body>
        {{ block "contents" . }}{{ end }}

        {{literal `
        <!--[if lt IE 9]>
<script src="/static/assets/js/respond.min.js"></script>
<script src="/static/assets/js/excanvas.min.js"></script>
<![endif]-->`}}

        <!-- Theme -->
        <script src="/static/assets/js/jquery.min.js" type="text/javascript"></script>
        <script src="/static/plugins/bootstrap/js/bootstrap.min.js" type="text/javascript"></script>
        <script src="/static/plugins/jquery-validation/jquery.validate.min.js" type="text/javascript"></script>
        {{ block "javascript" . }}
        {{ end }}
    </body>

</html>


