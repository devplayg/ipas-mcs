<link href="/static/modules/ipasreport/ipasreport.css" rel="stylesheet" type="text/css" />
<style>
    .customoverlay {position:relative;bottom:85px;border-radius:6px;border: 1px solid #ccc;border-bottom:2px solid #ddd;float:left;}
    .customoverlay:nth-of-type(n) {border:0; box-shadow:0px 1px 2px #888;}
    .customoverlay a {display:block;text-decoration:none;color:#000;text-align:center;border-radius:6px;font-size:14px;font-weight:bold;overflow:hidden;background: #d95050;background: #d95050 url(http://t1.daumcdn.net/localimg/localimages/07/mapapidoc/arrow_white.png) no-repeat right 14px center;}
    .customoverlay .title {display:block;text-align:center;background:#fff;margin-right:35px;padding:10px 15px;font-size:14px;font-weight:bold;}
    .customoverlay:after {content:'';position:absolute;margin-left:-12px;left:50%;bottom:-12px;width:22px;height:12px;background:url('http://t1.daumcdn.net/localimg/localimages/07/mapapidoc/vertex_white.png')}
</style>

<div class="modal fade" id="modal-ipas-report" tabindex="-1" role="basic" aria-hidden="true">
    <div class="modal-dialog modal-lg">
        <div class="modal-content">
            <div class="modal-header">
                <h4 class="modal-title">
                    <div class="caption pull-left">
                        <span class="caption-subject font-blue-sharp bold uppercas mr5">IPAS Report</span>
                    </div>
                    <div class="pull-right">
                        <button class="btn default btn-sm btn-rpt-date btn-rpt-theday hide"                 data-period="0"></button>
                        <button class="btn default btn-sm btn-rpt-date btn-rpt-period" data-period="0">{{i18n .Lang "report.day"}}</button>
                        <button class="btn default btn-sm btn-rpt-date btn-rpt-period" data-period="3">{{i18n .Lang "past 3 days"}}</button>
                        <button class="btn default btn-sm btn-rpt-date btn-rpt-period" data-period="7">{{i18n .Lang "past 1 week"}}</button>
                        <button class="btn default btn-sm btn-rpt-date btn-rpt-period" data-period="14">{{i18n .Lang "past 2 weeks"}}</button>
                        <button class="btn default btn-sm btn-rpt-date btn-rpt-period" data-period="30">{{i18n .Lang "past 1 month"}}</button>
                    </div>
                </h4>
            </div>
            <div class="modal-body">
                <div class="row">
                    <div class="col-sm-3">
                        <div class="rpt-equip-type text-center">
                            <img id="rpt-img-equipType" src="" alt="" />
                        </div>
                        <div class="text-center hide">
                            <button class="btn default btn-xs rpt-data rpt-ipas-orgName"></button>
                            <span class="rpt-data rpt-ipas-equipId"></span>
                        </div>
                        <div class="text-center">
                            <div class="rpt-data rpt-ipas-equipId s20 bold"></div>
                        </div>
                        <div class="mb10 mt20">
                            <div class="bold mb3">{{i18n .Lang "tag type"}}</div>
                            <pre class="mt-code"><span class="rpt-data rpt-ipas-equipType"></span></pre>
                        </div>
                        <div class="mb10">
                            <div class="bold mb3">USIM</div>
                            <pre class="mt-code"><span class="rpt-data rpt-ipas-usim"></span></pre>
                        </div>
                        <div class="mb10">
                            <div class="bold mb3">{{i18n .Lang "ipas.first registered"}}</div>
                            <pre class="mt-code"><span class="rpt-data rpt-ipas-created"></span></pre>
                        </div>
                        <div class="mb10">
                            <div class="bold mb3">{{i18n .Lang "ipas.last used"}}</div>
                            <pre class="mt-code"><span class="rpt-data rpt-ipas-updated"></span></pre>
                        </div>
                    </div>
                    <div class="col-sm-9 " style="line-height: 200%; border-left: 1px dashed #acacac;">
                        <div>
                            <label class="label label-success mr5">{{i18n .Lang "search period"}}</label>
                            <i>
                                <span class="rpt-startDate"></span> ~
                                <span class="rpt-endDate"></span>
                            </i>
                        </div>
                        <div class="row">
                            <div class="col-sm-5">
                                <div class="reportBox">
                                    <span class=""><i class="fa fa-bolt icon-width"></i> {{i18n .Lang "shock"}}</span>
                                    <span class="rpt-data rpt-counts-shock s24 font-grey-mint pull-right "></span>
                                </div>
                                <div class="reportBox">
                                    <span class=""><i class="icon-speedometer icon-width"></i> {{i18n .Lang "speeding"}}</span>
                                    <span class="rpt-data rpt-counts-speeding s24 font-grey-mint pull-right "></span>
                                </div>
                                <div class="reportBox">
                                    <span class=""><i class="icon-size-actual icon-width"></i> {{i18n .Lang "proximity"}}</span>
                                    <span class="rpt-data rpt-counts-proximity s24 font-grey-mint pull-right "></span>
                                </div>
                            </div>
                            <div class="col-sm-7">
                                <div class="rpt-log hide">
                                    <div class="">
                                        <span><i class="fa fa-square-o mr5"></i> {{i18n .Lang "occurrence date"}}</span>
                                        <span class="rpt-data rpt-log-eventDate pull-right"></span>
                                    </div>
                                    <div class="mt5">
                                        <span><i class="fa fa-square-o mr5"></i> {{i18n .Lang "tag"}}</span>
                                        <span class="rpt-data rpt-ipas-tag pull-right"></span>
                                    </div>
                                    <div class="mt5">
                                        <span><i class="fa fa-square-o mr5"></i> {{i18n .Lang "event type"}}</span>
                                        <span class="rpt-log-eventType pull-right"></span>
                                    </div>
                                    <div class="mt5">
                                        <span><i class="fa fa-square-o mr5"></i> {{i18n .Lang "location"}} (Lat./Lng.)</span>
                                        <span class="rpt-log-location pull-right"></span>
                                    </div>
                                    <div class="mt5">
                                        <span><i class="fa fa-square-o mr5"></i> SNR</span>
                                        <span class="pull-right">
                                            <span class="rpt-log-snr-value mr10"></span>
                                            <span class="rpt-log-snr"></span>
                                        </span>
                                    </div>
                                    <div class="mt5">
                                        <span><i class="fa fa-square-o mr5"></i> {{i18n .Lang "speed"}}</span>
                                        <span class="pull-right">
                                            <span class="rpt-log-speed"></span> km/h
                                        </span>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div id="map-rpt-ipas" class="table-rpt-events hide mt20 mb20" style="width:100%;height:350px;"></div>
                        <div id="toolbar-rpt-events" class="table-rpt-events s14"><i class="icon-clock"></i> {{i18n .Lang "recent events"}}</div>
                        <table  id="table-rpt-events"
                                class="table-condensed"
                                data-toolbar="#toolbar-rpt-events"
                                data-toggle="table"
                                data-show-columns="true"
                                {*Row강조*}
                                data-row-style="ipasEventRowStyle"
                                {* 정렬 *}
                                data-sort-name="date"
                                data-sort-order="desc"
                                {* 페이징 *}
                                data-page-size="15"
                                data-side-pagination="client"
                                data-pagination="true"
                                data-pagination-loop="true"
                        >
                            <thead>
                            <tr>
                                <th data-field="date" data-sortable="true" data-formatter="shortDateFormatter">{{i18n .Lang "occurrence date"}}</th>
                                <th data-field="org_id" data-sortable="true" data-visible="false" data-formatter="orgNameFormatter">{{i18n .Lang "org"}}</th>
                                <th data-field="group_id" data-sortable="true" data-visible="false" data-formatter="groupNameFormatter">{{i18n .Lang "group"}}</th>
                                <th data-field="event_type" data-sortable="true" data-formatter="ipaslogEventTypeFormatter">{{i18n .Lang "ipas.action"}}</th>
                                <th data-field="equip_id" data-formatter="ipasEquipIdFormatter" data-sortable="true" data-visible="false">{{i18n .Lang "tag"}}</th>
                                <th data-field="targets" data-formatter="ipaslogTargetsFormatter" data-sortable="true">{{i18n .Lang "ipas.target"}}</th>
                                <th data-field="location" data-sortable="true" data-formatter="ipaslogLocationFormatter" data-align="center">{{i18n .Lang "location"}}</th>
                                <th data-field="latitude" data-sortable="true" data-visible="false">{{i18n .Lang "latitude"}}</th>
                                <th data-field="longitude" data-sortable="true" data-visible="false">{{i18n .Lang "longitude"}}</th>
                                <th data-field="distance" data-sortable="true" data-formatter="ipaslogDistanceFormatter">{{i18n .Lang "distance"}} (m)</th>
                                <th data-field="speed" data-sortable="true" data-formatter="ipaslogSpeedingFormatter">{{i18n .Lang "speed"}} <small>(km/h)</small></th>
                                <th data-field="snr" data-sortable="true" data-formatter="snrFormatter">SNR&nbsp;&nbsp;</th>
                                <th data-field="usim" data-sortable="true" data-visible="false">USIM</th>
                                <th data-field="ip" data-sortable="true" data-formatter="int2ipFormatter" data-visible="false">IP</th>
                                <th data-field="recv_date" data-sortable="true" data-formatter="dateFormatter" data-visible="false">{{i18n .Lang "received date"}}</th>
                            </tr>
                            </thead>
                        </table>
                    </div>
                </div>


            </div><!-- modal-body -->
            <div class="modal-footer">
                <button type="button" class="btn default" data-dismiss="modal">{{i18n .Lang "close"}}</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div>

<div class="modal fade" id="modal-ipas-map" tabindex="-1" role="basic" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h4 class="modal-title">{{i18n .Lang "location"}}
                    <span class="pull-right">
                        <i class="fa fa-map-marker font-gallery"></i>
                        <span class="latitude"></span>,
                        <span class="longitude"></span>
                    </span>
                </h4>
            </div>
            <div class="modal-body">
                <div id="map-ipas" style="width:570px;height:400px;"></div>
            </div><!-- modal-body -->
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">{{i18n .Lang "close"}}</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div>

<script type="text/javascript" src="//dapi.kakao.com/v2/maps/sdk.js?appkey={{.daumMapKey}}&&libraries=drawing"></script>
<script src="/static/modules/ipasreport/ipasreport.js" type="text/javascript"></script>