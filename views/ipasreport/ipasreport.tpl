<div class="modal fade" id="modal-ipas-report" tabindex="-1" role="basic" aria-hidden="true">
    <div class="modal-dialog modal-lg   ">
        <div class="modal-content">
            <div class="modal-header">
                <h4 class="modal-title">IPAS Report</h4>
            </div>
            <div class="modal-body">
                <div class="row">
                    <div class="col-sm-3 text-center">
                        <div class="ipasreport-data ipasreport-equip-id"></div>
                    </div>
                    <div class="col-sm-9 " style="line-height: 200%; border-left: 1px dashed #acacac;">
                        <div class="row ipasreport-log ipasreport-filetrans ipasreport-email">
                            <div class="col-sm-3 col-xs-4"><i class="fa fa-square-o mr5"></i>발생시간</div>
                            <div class="col-sm-9 col-xs-8"><span class="ipasreport-data ipasreport-date strong"></span></div>
                        </div>
                        <p>&nbsp;</p>
                        <p>&nbsp;</p>
                        <p>&nbsp;</p>
                        <p>&nbsp;</p>
                        <p>&nbsp;</p>
                    </div>
                </div>
            </div><!-- modal-body -->
            <div class="modal-footer">
                <button type="submit" class="btn btn-primary">{{i18n .Lang "save"}}</button>
                <button type="button" class="btn btn-default" data-dismiss="modal">{{i18n .Lang "close"}}</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div>

<div class="modal fade" id="modal-ipas-map" tabindex="-1" role="basic" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h4 class="modal-title">Map
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

<script type="text/javascript" src="//dapi.kakao.com/v2/maps/sdk.js?appkey={{.daumMapKey}}"></script>
<script src="/static/modules/ipasreport/ipasreport.js" type="text/javascript"></script>