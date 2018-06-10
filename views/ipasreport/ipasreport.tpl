<div class="modal fade" id="modal-ipas-report" tabindex="-1" role="basic" aria-hidden="true">
    <div class="modal-dialog modal-lg">
        <div class="modal-content">
            <div class="modal-header">
                <h4 class="modal-title">
                    IPAS Report - <i><span class="rpt-data rpt-ipas-equipType bold font-blue-steel"></span></i>
                    <div class="pull-right">
                        <button class="btn default btn-sm btn-rpt-period" data-period="0">Today</button>
                        <button class="btn default btn-sm btn-rpt-period" data-period="1">1D</button>
                        <button class="btn default btn-sm btn-rpt-period" data-period="3">3D</button>
                        <button class="btn default btn-sm btn-rpt-period" data-period="7">1W</button>
                        <button class="btn default btn-sm btn-rpt-period" data-period="14">2W</button>
                        <button class="btn default btn-sm btn-rpt-period" data-period="30">1M</button>
                    </div>
                </h4>
            </div>
            <div class="modal-body">
                <div class="row">
                    <div class="col-sm-3 text-center">
                        <div class="rpt-equip-type">
                            <img id="rpt-img-equipType" src="" alt="" />
                        </div>
                        <div class="">
                            <button class="btn default btn-xs rpt-data rpt-ipas-orgName"></button>
                            <span class="rpt-data rpt-ipas-equipId"></span>
                        </div>
                    </div>
                    <div class="col-sm-9 " style="line-height: 200%; border-left: 1px dashed #acacac;">
                        <div>
                            <span class="rpt-startDate"></span> ~
                            <span class="rpt-endDate"></span>
                        </div>
                        <div class="row">
                            <div class="col-sm-5">
                                <div class="reportBox">
                                    <span class=""><i class="icon-power icon-width"></i> {{i18n .Lang "startup"}}</span>
                                    <span class="rpt-data rpt-counts-startup s24 font-grey-mint pull-right "></span>
                                </div>
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
                                <div class="clear mt10">
                                    <span><i class="fa fa-square-o"></i> {{i18n .Lang "tag"}}</span>
                                    <span class="rpt-data rpt-ipas-tag pull-right"></span>
                                </div>
                                <div class="clear mt10">
                                    <span><i class="fa fa-square-o"></i> USIM</span>
                                    <span class="rpt-data rpt-ipas-usim pull-right"></span>
                                </div>
                                <div class="clear mt10">
                                    <span><i class="fa fa-square-o"></i> {{i18n .Lang "last location"}}</span>
                                    <span class="rpt-data rpt-ipas-location pull-right"></span>
                                </div>
                                <div class="clear mt10">
                                <span><i class="fa fa-square-o"></i> {{i18n .Lang "ipas.first registered"}}</span>
                                <span class="rpt-data rpt-ipas-created pull-right"></span>
                            </div>
                                <div class="clear mt10">
                                    <span><i class="fa fa-square-o"></i> {{i18n .Lang "ipas.last used"}}</span>
                                    <span class="rpt-data rpt-ipas-updated pull-right"></span>
                                </div>
                            </div>
                        </div>

                        <div id="map-rpt-ipas" class="hide mt20" style="width:100%;height:350px;"></div>
                        <table  id="table-rpt-events"
                                class="table-condensed hide"
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
                <button type="submit" class="btn btn-primary">{{i18n .Lang "save"}}</button>
                <button type="button" class="btn btn-default" data-dismiss="modal">{{i18n .Lang "close"}}</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div>

<script type="text/javascript" src="//dapi.kakao.com/v2/maps/sdk.js?appkey={{.daumMapKey}}&&libraries=drawing"></script>
<script src="/static/modules/ipasreport/ipasreport.js" type="text/javascript"></script>