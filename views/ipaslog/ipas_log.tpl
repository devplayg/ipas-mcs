{{template "base.tpl" .}}

{{define "contents"}}
        <style>
            #map {
                width: 100%;
                height: 400px;
                border: 1px dashed #acacac;
            }

            .dashboard-stat .visual {
                height: 40px;
            }

            .dashboard-stat .details .number {
                height: 40px;
                padding-top: 0px;
                font-size: 25px;
            }
        </style>
     <pre class="hide">
        start_date={{.filter.StartDate}}
        end_date={{.filter.EndDate}}
        equip_id={{.filter.EquipId}}
        tag_pattern={{.filter.TagPattern}}
        fast_paging={{.filter.FastPaging}}
        limit={{.filter.Limit}}
        sort={{.filter.Sort}}
        order={{.filter.Order}}
        event_type={{range .filter.EventType}}{{.}}{{end}}
        org_id={{range .filter.OrgId}}{{.}},{{end}}
        group_id={{range .filter.GroupId}}{{.}}{{end}}
        {{.Lang}}
    </pre>
    <div class="portlet light bordered">
        <div class="portlet-body pt0">
            <div id="toolbar-log">
                <form id="form-filter" role="form" method="post">
                {{ .xsrfdata }}

                    <div class="form-body">
                        <div class="form-inline">
                            <div class="form-group">

                                <!-- 검색 시작날짜 -->
                                <div class="input-group date datetime" data-date="1979-09-16T05:25:07Z" data-date-format="yyyy-mm-dd HH:ii" data-link-field="">
                                    <input class="form-control mask-yyyymmddhhii" size="17" type="text" name="start_date" value="{{.filter.StartDate}}">
                                    <span class="input-group-addon"><i class="glyphicon glyphicon-th"></i></span>
                                </div>

                                <!-- 검색 끝날짜 -->
                                <div class="input-group date datetime">
                                    <input class="form-control mask-yyyymmddhhii" size="17" type="text" name="end_date" value="{{.filter.EndDate}}">
                                    <span class="input-group-addon"><i class="glyphicon glyphicon-th"></i></span>
                                </div>

                                <!-- 자산 선택 -->
                                <select id="select-orgs" name="org_id" class="selectpicker" multiple title="{{i18n .Lang "org"}}"  data-size="10" data-selected-text-format="count > 2"></select>
                                <select id="select-groups" name="group_id" class="selectpicker" multiple title="{{i18n .Lang "group"}}"  data-size="10" data-selected-text-format="count > 2"></select>

                                <!-- Buttons -->
                                <button type="submit" class="btn blue">{{i18n .Lang "search"}}</button>
                                <a class="btn btn-default" href="">{{i18n .Lang "cancel"}}</a>

                            {{if eq .filter.FastPaging "on"}} {{/* 고속 페이징 */}}
                                <div class="input-group btn-group btn-page-group">
                                    <button type="button" class="btn blue btn-move-page btn-prev" data-direction="-1" data-loading-text="&lt;">&lt;</button>
                                    <button type="button" class="btn blue btn-move-page btn-page-text" data-direction="0">1</button>
                                    <button type="button" class="btn blue btn-move-page btn-next" data-direction="1" data-loading-text="&gt;">&gt;</button>
                                </div>
                            {{end}}
                                <a href="#" data-toggle="modal" data-target="#modal-filter"><i class="fa fa-filter icon-filter hidden font-red"></i>{{i18n .Lang "detail_filter"}}</a>
                            </div>
                        </div>
                    </div>

                    <!-- 상세필터 -->
                    <div id="modal-filter" class="modal fade" tabindex="-1" role="dialog">
                        <div class="modal-dialog" role="document">
                            <div class="modal-content">
                                <div class="modal-header">
                                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                                    <h4 class="modal-title"><i class="fa fa-filter"></i> Filter</h4>
                                </div>

                                <div class="modal-body">
                                    <div class="row">
                                        <div class="col-sm-6 form-group">
                                            <label class="control-label">{{i18n .Lang "tag"}}</label>
                                            <input type="text" class="form-control" name="tag_pattern" value="{{.filter.TagPattern}}">
                                        </div>
                                        <div class="col-sm-6 form-group ">
                                            <label class="control-label">{{i18n .Lang "event type"}}</label>
                                            <select name="event_type" class="selectpicker" data-width="100%" data-size="5" multiple title="{{i18n .Lang "event type"}}">
                                                <option value="1">{{i18n .Lang "startup"}}</option>
                                                <option value="2">{{i18n .Lang "shock"}}</option>
                                                <option value="3">{{i18n .Lang "speeding"}}</option>
                                                <option value="4">{{i18n .Lang "proximity"}}</option>
                                            </select>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col-sm-4 form-group">
                                            <label class="control-label">{{i18n .Lang "page size"}}</label>
                                            <input type="text" class="form-control mask-09999" name="limit" value="{{.filter.Limit}}">
                                        </div>
                                        <div class="col-sm-4 form-group">
                                            <label class="mt-checkbox mt-checkbox-outline mt30">
                                                <input type="checkbox" name="fast_paging" {{if eq .filter.FastPaging "on"}}checked{{end}}> {{i18n .Lang "fast_paging"}}
                                                <span></span>
                                            </label>
                                        </div>
                                        <div class="col-sm-4 form-group">
                                            <label class="mt-checkbox mt-checkbox-outline mt30">
                                                <input type="checkbox" name="event_map" {{if eq .filter.EventMap "on"}}checked{{end}}> {{i18n .Lang "event map"}}
                                                <span></span>
                                            </label>
                                        </div>
                                    </div>
                                </div><!-- modal-body-->
                                <div class="modal-footer">
                                    <button type="submit" class="btn btn-primary">{{i18n .Lang "search"}}</button>
                                    <button type="button" class="btn btn-default" data-dismiss="modal">{{i18n .Lang "close"}}</button>
                                </div>
                            </div>
                        </div>
                    </div> <!-- #modal-filter -->
                </form>
            </div>

            {{if eq .filter.EventMap "on"}}
            <div class="row mt20 mb20">
                <div class="col-lg-12">
                    <div id="map"></div>
                </div>
            </div>
            {{end}}

            <table  id="table-log"
                    class="table-condensed"
                    data-toggle="table"
                    {{if eq .filter.EventMap "off"}}
                    data-toolbar="#toolbar-log"
                    data-show-refresh="true"
                    data-show-columns="true"
                    {{end}}
                    {* 내보내기 *}
                    data-show-export="true"
                    data-export-types="['csv', 'excel']"
                    {*Row강조*}
                    data-row-style="ipasEventRowStyle"
                    {* 정렬 *}
                    data-sort-name="{{.filter.Sort}}"
                    data-sort-order="{{.filter.Order}}"
                    {* 페이징 *}
                    data-page-size="{{.filter.Limit}}"
                    {{/*data-pagination-v-align="both"*/}}
                    {{if eq .filter.FastPaging "on"}} {* 고속 페이징 *}
                        data-side-pagination="client"
                    {{else}} {* 일반 페이징 *}
                        data-url="/getIpasLogs?start_date={{.filter.StartDate}}&end_date={{.filter.EndDate}}&fast_paging={{.filter.FastPaging}}&equip_id={{.filter.EquipId}}{{range .filter.EventType}}&event_type={{.}}{{end}}{{range .filter.OrgId}}&org_id={{.}}{{end}}{{range .filter.GroupId}}&group_id={{.}}{{end}}"
                        data-pagination="true"
                        data-side-pagination="server"
                        data-pagination-loop="false"
                    {{end}}
            >
                <thead>
                <tr>
                    <th data-field="date" data-sortable="true" data-formatter="dateFormatter">{{i18n .Lang "occurrence date"}}</th>
                    <th data-field="org_id" data-sortable="true" data-formatter="orgNameFormatter">{{i18n .Lang "org"}}</th>
                    <th data-field="group_id" data-sortable="true" data-formatter="groupNameFormatter">{{i18n .Lang "group"}}</th>
                    <th data-field="event_type" data-sortable="true" data-formatter="ipaslogEventTypeFormatter">{{i18n .Lang "ipas.action"}}</th>
                    <th data-field="equip_id" data-formatter="ipasEquipIdFormatter" data-sortable="true">{{i18n .Lang "tag"}}</th>
                    <th data-field="targets" data-formatter="ipaslogTargetsFormatter" data-sortable="true">{{i18n .Lang "ipas.target"}}</th>
                    <th data-field="location" data-sortable="true" data-formatter="ipaslogLocationFormatter" data-align="center">{{i18n .Lang "location"}}</th>
                    <th data-field="latitude" data-sortable="true" data-visible="false">{{i18n .Lang "latitude"}}</th>
                    <th data-field="longitude" data-sortable="true" data-visible="false">{{i18n .Lang "longitude"}}</th>
                    <th data-field="distance" data-sortable="true" data-formatter="ipaslogDistanceFormatter">{{i18n .Lang "distance"}} (m)</th>
                    <th data-field="speed" data-sortable="true" data-formatter="ipaslogSpeedingFormatter">{{i18n .Lang "speed"}} <small>(km/h)</small></th>
                    <th data-field="snr" data-sortable="true" data-formatter="snrFormatter">SNR</th>
                    <th data-field="usim" data-sortable="true" data-visible="false">USIM</th>
                    <th data-field="ip" data-sortable="true" data-formatter="int2ipFormatter" data-visible="false">IP</th>
                    <th data-field="recv_date" data-sortable="true" data-formatter="dateFormatter" data-visible="false">{{i18n .Lang "received date"}}</th>
                </tr>
                </thead>
            </table>

        </div>
    </div>
{{end}}

{{define "javascript"}}
    {{template "ipasreport/ipasreport.tpl" .}}
    {{if eq .filter.EventMap "on"}}
        <script>
            var map = null;
            // infoWIndow = null;
            var customLabel = {
                startup: {
                    label: 'ST'
                },
                shock: {
                    label: 'SH'
                },
                speeding: {
                    label: 'SP'
                },
                proximity: {
                    label: 'PX'
                }
            };

            function initMap() {
                map = new google.maps.Map( document.getElementById( "map" ), {
                    zoom: 8,
                    center: new google.maps.LatLng( 37.532600, 127.024612 ), // Seoul

                    // Basic Map Types
                    //      https://developers.google.com/maps/documentation/javascript/maptypes?hl=ko
                    mapTypeId: 'roadmap' // roadmap, satellite, hybrid, terrain
                });


                var script = document.createElement('script');
                script.src = '/getMapLogs?start_date={{.filter.StartDate}}&end_date={{.filter.EndDate}}&fast_paging={{.filter.FastPaging}}&equip_id={{.filter.EquipId}}{{range .filter.EventType}}&event_type={{.}}{{end}}{{range .filter.OrgId}}&org_id={{.}}{{end}}{{range .filter.GroupId}}&group_id={{.}}{{end}}&mode=gmimport';
                // script.src = 'https://developers.google.com/maps/documentation/javascript/examples/json/earthquake_GeoJSONP.js';
                document.getElementsByTagName('head')[0].appendChild(script);
            }

            // Loop through the results array and place a marker for each
            // set of coordinates.
            window.mapfeed_callback = function( events ) {
                var infoWindow = new google.maps.InfoWindow;
                var image = 'https://developers.google.com/maps/documentation/javascript/examples/full/images/beachflag.png';


                for (var i = 0; i < events.length; i++) {
                    var eventType;
                    if ( events[i].event_type === 1 ) {
                        eventType = "startup";
                    } else if ( events[i].event_type === 2 ) {
                        eventType = "shock";
                    } else if ( events[i].event_type === 3 ) {
                        eventType = "speeding";
                    } else if ( events[i].event_type === 4 ) {
                        eventType = "proximity";
                    } else {
                        eventType = "unknown";
                    }
                    var infowincontent = '<div class="bold  s16">' + events[i].org_name + ' / ' + events[i].equip_id + '</div>';
                    infowincontent += '<div class="bold s12 mt5 mb10 font-red">' + eventType.toUpperCase() + '</div>';
                    infowincontent += '<div>';
                    infowincontent += '- Event type: ' + eventType + '<br>';
                    infowincontent += '- Time: ' + events[i].date + '<br>';
                    infowincontent += '- Latitude: ' + events[i].latitude + '<br>';
                    infowincontent += '- Longitude: ' + events[i].longitude + '<br>';
                    infowincontent += '</div>';

                    var icon = customLabel[ eventType ] || {};
                    console.log(eventType);

                    var latLng = new google.maps.LatLng( events[i].latitude, events[i].longitude );
                    var marker = new google.maps.Marker({
                        position: latLng,
                        map: map,
                        title: events[i].org_name + " / " + events[i].equip_id,
                        label:  {
                            text: ""+icon.label,
                            color: "#ffffff",
                            fontSize: "12px",
                        },
                        // label: icon.label,
                        // icon: {
                        //     // path: google.maps.SymbolPath.CIRCLE,
                        //     scale: 8.5,
                        //     fillColor: "#F00",
                        //     fillOpacity: 0.4,
                        //     strokeWeight: 0.4
                        // },
                        // icon:  image,
                        // animation: google.maps.Animation.,
                        infowincontent: infowincontent
                    });

                    marker.addListener('click', function() {
                        infoWindow.setContent(this.infowincontent);
                        infoWindow.open(map, this);
                    });
                }
            }
        </script>
        <script async defer src="https://maps.googleapis.com/maps/api/js?key=AIzaSyCQHrPThkndn-kfySgCUgkZxEbdTU8rrNg&callback=initMap"></script>
    {{end}}

    <script src="/static/modules/{{.ctrl}}/ipas_log.js"></script>
    <script src="/static/modules/{{.ctrl}}/formatter.js"></script>
    <script>
        //var filterUrl = "start_date={{.filter.StartDate}}&end_date={{.filter.EndDate}}&fast_paging=on&equip_id={{.filter.EquipId}}{{range .filter.EventType}}&event_type={{.}}{{end}}{{range .filter.OrgId}}&org_id={{.}}{{end}}{{range .filter.GroupId}}&group_id={{.}}{{end}}";
    </script>
{{end}}