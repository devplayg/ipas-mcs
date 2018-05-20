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
        <link href="/static/assets/css/components-rounded.css" rel="stylesheet" id="style_components" type="text/css" />
        <link href="/static/assets/css/layout.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/css/themes/darkblue.min.css" rel="stylesheet" type="text/css" id="style_color" />
        <link href="/static/assets/css/plugins.min.css" rel="stylesheet" type="text/css" />

        <!-- Fonts -->
        <link href="/static/assets/font/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/font/simple-line-icons/css/simple-line-icons.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/font/{{.Lang}}.css" rel="stylesheet" type="text/css" />

        <!-- Plugins -->
        <link href="/static/plugins/bootstrap-switch/css/bootstrap-switch.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/plugins/bootstrap-select/css/bootstrap-select.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/plugins/bootstrap-table/bootstrap-table.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/plugins/bootstrap-datetimepicker/css/bootstrap-datetimepicker.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/plugins/bootstrap-select/css/bootstrap-select.min.css" rel="stylesheet" />
        {{/*<link href="/static/plugins/sweetalert/sweetalert2.min.css" rel="stylesheet" type="text/css" />*/}}
        <link href="/static/plugins/waitMe/waitMe.min.css" rel="stylesheet" type="text/css" />
        <link href="/static/plugins/flag-icon-css/css/flag-icon.css" rel="stylesheet" type="text/css" />
        <link href="/static/assets/css/custom.css" rel="stylesheet" type="text/css" />

        {{ block "css" . }}{{ end }}
        <link rel="shortcut icon" href="favicon.ico" />
    </head>

    <body class="page-header-fixed page-sidebar-closed-hide-logo page-content-white page-container-bg-solid">
        <div class="page-header navbar navbar-fixed-top">
            <div class="page-header-inner ">
                <div class="page-logo">
                    <a href="/">
                        <img src="/static/assets/img/logo.png" alt="logo" class="logo-default" style="margin-top: 15px; margin-left: -20px;" /> </a>
                    <div class="menu-toggler sidebar-toggler">
                        <span></span>
                    </div>
                </div>
                <form class="search-form" action="page_general_search_2.html" method="GET">
                    <div class="input-group">
                        <input type="text" class="form-control" placeholder="Search..." name="query">
                        <span class="input-group-btn">
                            <a href="javascript:;" class="btn submit">
                                <i class="icon-magnifier"></i>
                            </a>
                        </span>
                    </div>
                </form>
                <a href="javascript:;" class="menu-toggler responsive-toggler" data-toggle="collapse" data-target=".navbar-collapse">
                    <span></span>
                </a>
                <div class="top-menu">
                    <ul class="nav navbar-nav pull-right">
                        <li class="hidden dropdown dropdown-extended dropdown-notification" id="header_notification_bar">
                            <a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown" data-hover="dropdown" data-close-others="true">
                                <i class="icon-bell"></i>
                                <span class="badge badge-default"> 1 </span>
                            </a>
                            <ul class="dropdown-menu">
                                <li class="external">
                                    <h3>
                                        <span class="bold">12 pending</span> notifications</h3>
                                    <a href="page_user_profile_1.html">view all</a>
                                </li>
                                <li>
                                    <ul class="dropdown-menu-list scroller" style="height: 250px;" data-handle-color="#637283">
                                        <li>
                                            <a href="javascript:;">
                                                <span class="time">just now</span>
                                                <span class="details">
                                                    <span class="label label-sm label-icon label-success">
                                                        <i class="fa fa-plus"></i>
                                                    </span> New user registered. </span>
                                            </a>
                                        </li>
                                    </ul>
                                </li>
                            </ul>
                        </li>
                        <li class="hidden dropdown dropdown-extended dropdown-inbox" id="header_inbox_bar">
                            <a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown" data-hover="dropdown" data-close-others="true">
                                <i class="icon-envelope-open"></i>
                                <span class="badge badge-default"> 1 </span>
                            </a>
                            <ul class="dropdown-menu">
                                <li class="external">
                                    <h3>You have
                                        <span class="bold">7 New</span> Messages</h3>
                                    <a href="app_inbox.html">view all</a>
                                </li>
                                <li>
                                    <ul class="dropdown-menu-list scroller" style="height: 275px;" data-handle-color="#637283">
                                        <li>
                                            <a href="#">
                                                <span class="photo">
                                                    <img src="/static/assets/img/won.png" class="img-circle" alt=""> </span>
                                                <span class="subject">
                                                    <span class="from"> Nhakoe Snow </span>
                                                    <span class="time">Just Now </span>
                                                </span>
                                                <span class="message"> Vivamus sed auctor nibh congue nibh. auctor nibh auctor nibh... </span>
                                            </a>
                                        </li>
                                    </ul>
                                </li>
                            </ul>
                        </li>
                        <li class="hidden dropdown dropdown-extended dropdown-tasks" id="header_task_bar">
                            <a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown" data-hover="dropdown" data-close-others="true">
                                <i class="icon-calendar"></i>
                                <span class="badge badge-default"> 3 </span>
                            </a>
                            <ul class="dropdown-menu extended tasks">
                                <li class="external">
                                    <h3>You have
                                        <span class="bold">12 pending</span> tasks</h3>
                                    <a href="app_todo.html">view all</a>
                                </li>
                                <li>
                                    <ul class="dropdown-menu-list scroller" style="height: 275px;" data-handle-color="#637283">
                                        <li>
                                            <a href="javascript:;">
                                                <span class="task">
                                                    <span class="desc">New release v1.2 </span>
                                                    <span class="percent">30%</span>
                                                </span>
                                                <span class="progress">
                                                    <span style="width: 40%;" class="progress-bar progress-bar-success" aria-valuenow="40" aria-valuemin="0" aria-valuemax="100">
                                                        <span class="sr-only">40% Complete</span>
                                                    </span>
                                                </span>
                                            </a>
                                        </li>
                                        <li>
                                            <a href="javascript:;">
                                                <span class="task">
                                                    <span class="desc">New UI release</span>
                                                    <span class="percent">38%</span>
                                                </span>
                                                <span class="progress progress-striped">
                                                    <span style="width: 38%;" class="progress-bar progress-bar-important" aria-valuenow="18" aria-valuemin="0" aria-valuemax="100">
                                                        <span class="sr-only">38% Complete</span>
                                                    </span>
                                                </span>
                                            </a>
                                        </li>
                                    </ul>
                                </li>
                            </ul>
                        </li>
                        <li class="dropdown dropdown-user">
                            <a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown" data-hover="dropdown" data-close-others="true">
                                <img alt="" class="img-circle" src="/static/assets/img/won.png" />
                                <span class="username username-hide-on-mobile"> {{.member.Username}} {{.IsLogged}}</span>
                                <i class="fa fa-angle-down"></i>
                            </a>
                            <ul class="dropdown-menu dropdown-menu-default">
                                <li>
                                    <a href="/signout">
                                        <i class="icon-key"></i> {{i18n .Lang "signout"}}</a>
                                </li>
                            </ul>
                        </li>
                        <li class="dropdown dropdown-user">
                            <a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown" data-hover="dropdown" data-close-others="true">
                                <span class="username username-hide-on-mobile"> {{.CurLang}}</span>
                                <i class="fa fa-angle-down"></i>
                            </a>
                            <ul class="dropdown-menu dropdown-menu-default">
                                {{range .RestLangs}}
                                <li>
                                    <a href="javascript:;" data-lang="{{.Lang}}" class="lang-changed">
                                        <span class="flag-icon flag-icon-{{i18n $.Lang .Name | toLower}} flag-border mr10"></span>
                                        {{i18n $.Lang .Name}}
                                    </a>
                                </li>
                                {{end}}
                            </ul>
                        </li>
                        <li class="dropdown dropdown-quick-sidebar-toggler">
                            <a href="javascript:;" class="dropdown-toggle">
                                <i class="icon-logout"></i>
                            </a>
                        </li>
                    </ul>
                </div>
            </div>
        </div> <!-- // .page-header -->
        <div class="clearfix"> </div>
        <div class="page-container">
            <div class="page-sidebar-wrapper">
                <div class="page-sidebar navbar-collapse collapse">
                    <ul class="page-sidebar-menu page-header-fixed" data-keep-expanded="false" data-auto-scroll="true" data-slide-speed="200" style="padding-top: 20px">
                        <li class="sidebar-toggler-wrapper hide">
                            <div class="sidebar-toggler">
                                <span></span>
                            </div>
                        </li>
                        <li class="sidebar-search-wrapper"> </li>

                        <li class="heading">
                            <h3 class="uppercase">Features</h3>
                        </li>

                        <li class="nav-item">
                            <a href="javascript:;" class="nav-link nav-toggle">
                                <i class="icon-grid"></i>
                                <span class="title">{{i18n .Lang "dashboard"}}</span>
                                <span class="arrow open"></span>
                            </a>
                            <ul class="sub-menu">
                                <li class="nav-item">
                                    <a href="/dashboard" class="nav-link"><span class="title">{{i18n .Lang "dashboard"}} I</span></a>
                                </li>
                                <li class="nav-item">
                                    <a href="/darkboard" class="nav-link"><span class="title">{{i18n .Lang "dashboard"}} II</span></a>
                                </li>
                                <li class="nav-item">
                                    <a href="/realtimelogs" class="nav-link"><span class="title">{{i18n .Lang "menu.ipas event"}}</span></a>
                                </li>

                            {{/*<li class="nav-item ">*/}}
                            {{/*<a href="/samplelogs" class="nav-link"><span class="title">Sample logs</span></a>*/}}
                            {{/*</li>*/}}
                            </ul>
                        </li>
                        <li class="nav-item">
                            <a href="javascript:;" class="nav-link nav-toggle">
                                <i class="icon-list"></i>
                                <span class="title">{{i18n .Lang "log"}}</span>
                                <span class="arrow open"></span>
                            </a>
                            <ul class="sub-menu">
                                <li class="nav-item">
                                    <a href="/ipaslogs" class="nav-link"><span class="title">{{i18n .Lang "menu.ipas event"}}</span></a>
                                </li>
                                {{/*<li class="nav-item ">*/}}
                                    {{/*<a href="/samplelogs" class="nav-link"><span class="title">Sample logs</span></a>*/}}
                                {{/*</li>*/}}
                            </ul>
                        </li>
                        <li class="nav-item">
                            <a href="javascript:;" class="nav-link nav-toggle">
                                <i class="icon-settings"></i>
                                <span class="title">{{i18n .Lang "preferences"}}</span>
                                <span class="arrow open"></span>
                            </a>
                            <ul class="sub-menu">
                                <li class="nav-item">
                                    <a href="/assets" class="nav-link">
                                        <span class="title">{{i18n .Lang "menu.asset management"}}</span>
                                    </a>
                                </li>
                                <li class="nav-item">
                                    <a href="/members" class="nav-link">
                                        <span class="title">{{i18n .Lang "menu.user management"}}</span>
                                    </a>
                                </li>
                                <li class="nav-item">
                                    <a href="/config" class="nav-link">
                                        <span class="title">{{i18n .Lang "menu.system settings" }}</span>
                                    </a>
                                </li>
                            </ul>
                        </li>
                    </ul>
                </div>
            </div>
            <div class="page-content-wrapper">
                <div class="page-content">
                    <div class="page-bar mb20">
                        <ul class="page-breadcrumb">
                            <li>
                                <span class="menu-depth1-text">{{i18n .Lang "menu.log"}}</span>
                                <i class="fa fa-circle"></i>
                            </li>
                            <li>
                                <span class="menu-depth2-text">{{i18n .Lang "menu.ipas event"}}</span>
                            </li>
                        </ul>
                        <div class="page-toolbar">
                            <div class="btn-group pull-right">
                                <button type="button" class="btn green btn-sm btn-outline dropdown-toggle" data-toggle="dropdown"> Actions
                                    <i class="fa fa-angle-down"></i>
                                </button>
                                <ul class="dropdown-menu pull-right" role="menu">
                                    <li>
                                        <a href="#">
                                            <i class="icon-bell"></i> Action</a>
                                    </li>
                                    <li>
                                        <a href="#">
                                            <i class="icon-shield"></i> Another action</a>
                                    </li>
                                    <li>
                                        <a href="#">
                                            <i class="icon-user"></i> Something else here</a>
                                    </li>
                                    <li class="divider"> </li>
                                    <li>
                                        <a href="#">
                                            <i class="icon-bag"></i> Separated link</a>
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </div>
                    <h3 class="page-title hidden"> Search Box On Header 11
                        <small>search box on header</small>
                    </h3>
                    {{ block "contents" . }}{{ end }}
                </div>
            </div>
            <a href="javascript:;" class="page-quick-sidebar-toggler">
                <i class="icon-login"></i>
            </a>
            <div class="page-quick-sidebar-wrapper" data-close-on-body-click="false">
                <div class="page-quick-sidebar">
                    <ul class="nav nav-tabs">
                        <li class="active">
                            <a href="javascript:;" data-target="#quick_sidebar_tab_1" data-toggle="tab"> Users
                                <span class="badge badge-danger">2</span>
                            </a>
                        </li>
                        <li>
                            <a href="javascript:;" data-target="#quick_sidebar_tab_2" data-toggle="tab"> Alerts
                                <span class="badge badge-success">7</span>
                            </a>
                        </li>
                        <li class="dropdown">
                            <a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown"> More
                                <i class="fa fa-angle-down"></i>
                            </a>
                            <ul class="dropdown-menu pull-right">
                                <li>
                                    <a href="javascript:;" data-target="#quick_sidebar_tab_3" data-toggle="tab">
                                        <i class="icon-settings"></i> Settings </a>
                                </li>
                            </ul>
                        </li>
                    </ul>
                    <div class="tab-content">
                        <div class="tab-pane active page-quick-sidebar-chat" id="quick_sidebar_tab_1">
                            <div class="page-quick-sidebar-chat-users" data-rail-color="#ddd" data-wrapper-class="page-quick-sidebar-list">
                                <h3 class="list-heading">Staff</h3>
                                <ul class="media-list list-items">
                                    <li class="media">
                                        <div class="media-status">
                                            <span class="badge badge-success">8</span>
                                        </div>
                                        <img class="media-object" src="/static/assets/img/won.png" alt="...">
                                        <div class="media-body">
                                            <h4 class="media-heading">Bob Nilson</h4>
                                            <div class="media-heading-sub"> Project Manager </div>
                                        </div>
                                    </li>

                                </ul>
                                <h3 class="list-heading">Customers</h3>
                                <ul class="media-list list-items">
                                    <li class="media">
                                        <div class="media-status">
                                            <span class="badge badge-warning">2</span>
                                        </div>
                                        <img class="media-object" src="/static/assets/img/won.png" alt="...">
                                        <div class="media-body">
                                            <h4 class="media-heading">Lara Kunis</h4>
                                            <div class="media-heading-sub"> CEO, Loop Inc </div>
                                            <div class="media-heading-small"> Last seen 03:10 AM </div>
                                        </div>
                                    </li>
                                </ul>
                            </div>
                            <div class="page-quick-sidebar-item">
                                <div class="page-quick-sidebar-chat-user">
                                    <div class="page-quick-sidebar-nav">
                                        <a href="javascript:;" class="page-quick-sidebar-back-to-list">
                                            <i class="icon-arrow-left"></i>Back</a>
                                    </div>
                                    <div class="page-quick-sidebar-chat-user-messages">
                                        <div class="post out">
                                            <img class="avatar" alt="" src="/static/assets/img/won.png" />
                                            <div class="message">
                                                <span class="arrow"></span>
                                                <a href="javascript:;" class="name">Bob Nilson</a>
                                                <span class="datetime">20:15</span>
                                                <span class="body"> When could you send me the report ? </span>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="page-quick-sidebar-chat-user-form">
                                        <div class="input-group">
                                            <input type="text" class="form-control" placeholder="Type a message here...">
                                            <div class="input-group-btn">
                                                <button type="button" class="btn green">
                                                    <i class="icon-paper-clip"></i>
                                                </button>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="tab-pane page-quick-sidebar-alerts" id="quick_sidebar_tab_2">
                            <div class="page-quick-sidebar-alerts-list">
                                <h3 class="list-heading">General</h3>
                                <ul class="feeds list-items">
                                    <li>
                                        <div class="col1">
                                            <div class="cont">
                                                <div class="cont-col1">
                                                    <div class="label label-sm label-info">
                                                        <i class="fa fa-check"></i>
                                                    </div>
                                                </div>
                                                <div class="cont-col2">
                                                    <div class="desc"> You have 4 pending tasks.
                                                        <span class="label label-sm label-warning "> Take action
                                                            <i class="fa fa-share"></i>
                                                        </span>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="col2">
                                            <div class="date"> Just now </div>
                                        </div>
                                    </li>
                                </ul>
                                <h3 class="list-heading">System</h3>
                                <ul class="feeds list-items">
                                    <li>
                                        <div class="col1">
                                            <div class="cont">
                                                <div class="cont-col1">
                                                    <div class="label label-sm label-info">
                                                        <i class="fa fa-check"></i>
                                                    </div>
                                                </div>
                                                <div class="cont-col2">
                                                    <div class="desc"> You have 4 pending tasks.
                                                        <span class="label label-sm label-warning "> Take action
                                                            <i class="fa fa-share"></i>
                                                        </span>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="col2">
                                            <div class="date"> Just now </div>
                                        </div>
                                    </li>
                                </ul>
                            </div>
                        </div>
                        <div class="tab-pane page-quick-sidebar-settings" id="quick_sidebar_tab_3">
                            <div class="page-quick-sidebar-settings-list">
                                <h3 class="list-heading">General Settings</h3>
                                <ul class="list-items borderless">
                                    <li> Enable Notifications
                                        <input type="checkbox" class="make-switch" checked data-size="small" data-on-color="success" data-on-text="ON" data-off-color="default" data-off-text="OFF"> </li>
                                    <li> Allow Tracking
                                        <input type="checkbox" class="make-switch" data-size="small" data-on-color="info" data-on-text="ON" data-off-color="default" data-off-text="OFF"> </li>
                                    <li> Log Errors
                                        <input type="checkbox" class="make-switch" checked data-size="small" data-on-color="danger" data-on-text="ON" data-off-color="default" data-off-text="OFF"> </li>
                                    <li> Auto Sumbit Issues
                                        <input type="checkbox" class="make-switch" data-size="small" data-on-color="warning" data-on-text="ON" data-off-color="default" data-off-text="OFF"> </li>
                                    <li> Enable SMS Alerts
                                        <input type="checkbox" class="make-switch" checked data-size="small" data-on-color="success" data-on-text="ON" data-off-color="default" data-off-text="OFF"> </li>
                                </ul>
                                <h3 class="list-heading">System Settings</h3>
                                <ul class="list-items borderless">
                                    <li> Security Level
                                        <select class="form-control input-inline input-sm input-small">
                                            <option value="1">Normal</option>
                                            <option value="2" selected>Medium</option>
                                            <option value="e">High</option>
                                        </select>
                                    </li>
                                    <li> Failed Email Attempts
                                        <input class="form-control input-inline input-sm input-small" value="5" /> </li>
                                    <li> Secondary SMTP Port
                                        <input class="form-control input-inline input-sm input-small" value="3560" /> </li>
                                    <li> Notify On System Error
                                        <input type="checkbox" class="make-switch" checked data-size="small" data-on-color="danger" data-on-text="ON" data-off-color="default" data-off-text="OFF"> </li>
                                    <li> Notify On SMTP Error
                                        <input type="checkbox" class="make-switch" checked data-size="small" data-on-color="warning" data-on-text="ON" data-off-color="default" data-off-text="OFF"> </li>
                                </ul>
                                <div class="inner-content">
                                    <button class="btn btn-success">
                                        <i class="icon-settings"></i> Save Changes</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="page-footer">
            <div class="page-footer-inner"> 2018 &copy;
                <a href="#" title="" target="_blank">Kyungwoo Inc.</a>
            </div>
            <div class="scroll-to-top">
                <i class="icon-arrow-up"></i>
            </div>
        </div>
        {{literal `
        <!--[if lt IE 9]>
<script src="/static/assets/js/respond.min.js"></script>
<script src="/static/assets/js/excanvas.min.js"></script>
<![endif]-->`}}

        <!-- Theme -->
        <script src="/static/assets/js/jquery.min.js" type="text/javascript"></script>
        <script src="/static/plugins/bootstrap/js/bootstrap.min.js" type="text/javascript"></script>
        <script src="/static/assets/js/app.min.js" type="text/javascript"></script>
        <script src="/static/assets/js/components-bootstrap-select.min.js" type="text/javascript"></script>
        <script src="/static/assets/js/layout.js" type="text/javascript"></script>
        <script src="/static/assets/js/quick-sidebar.min.js" type="text/javascript"></script>

        <!-- Plugins -->
        <script src="/static/plugins/bootstrap-switch/js/bootstrap-switch.min.js" type="text/javascript"></script>
        <script src="/static/plugins/bootstrap-select/js/bootstrap-select.min.js" type="text/javascript"></script>
        <script src="/static/plugins/jquery-slimscroll/jquery.slimscroll.min.js" type="text/javascript"></script>
        <!-- Bootstrap-table -->
        <script src="/static/plugins/bootstrap-table/bootstrap-table.min.js" type="text/javascript"></script>
        <script src="/static/plugins/bootstrap-table/locale/bootstrap-table-{{.Lang}}.min.js" type="text/javascript"></script>

        <script src="/static/plugins/jquery-validation/jquery.validate.min.js" type="text/javascript"></script>
        <script src="/static/plugins/jquery-validation/localization/messages-{{.Lang}}.min.js" type="text/javascript"></script>
        <script src="/static/plugins/jquery-mask/jquery.mask.min.js" type="text/javascript"></script>
        <script src="/static/plugins/bootstrap-datetimepicker/js/bootstrap-datetimepicker.min.js" type="text/javascript"></script>
        <script src="/static/plugins/sweetalert2/sweetalert2.all.min.js"></script>
        <script src="/static/plugins/jquery-base64/jquery.base64.min.js"></script>
        <script src="/static/plugins/moment/moment.min.js" type="text/javascript"></script>
        <script src="/static/plugins/waitMe/waitMe.min.js" type="text/javascript"></script>
        <script src="/static/assets/js/jquery.cookie.js" type="text/javascript"></script>
        <script src="/static/assets/js/common.js" type="text/javascript"></script>
        <script src="/static/assets/js/formatter.js" type="text/javascript"></script>
        <script>
            // Timezone
            var member = {
                "timezone": "{{.member.Timezone}}" || "Asia/Seoul",
                "position": {{.member.Position}},
            };
            var reqVars = {{.reqVars}};
            reqVars["ctrl"] = {{.ctrl}};
            reqVars["act"] = {{.act}};

            var daumMapKey = {{.daumMapKey}},
                felang = {{.frontLang}}, // front-end languages
                lang = {{.Lang}};
        </script>

        <!-- Javascript -->
        {{ block "javascript" . }}{{ end }}
    </body>
</html>

