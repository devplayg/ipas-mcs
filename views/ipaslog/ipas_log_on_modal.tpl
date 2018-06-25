<!-- 상세필터 -->
<div id="modal-ipaslog" class="modal fade" tabindex="-1" role="dialog">
    <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4 class="modal-title"><i class="icon-list"></i> {{i18n .Lang "log"}}</h4>
            </div>

            <div class="modal-body">
                <div id="toolbar-ipaslog-on-modal">
                    <button type="button" class="btn blue btn-move-page btn-prev" data-direction="-1" data-loading-text="&lt;">&lt;</button>
                    <button type="button" class="btn blue btn-move-page btn-page-text" data-direction="0">1</button>
                    <button type="button" class="btn blue btn-move-page btn-next" data-direction="1" data-loading-text="&gt;">&gt;</button>
                </div>
                <table  id="table-ipaslog-on-modal"
                        class="modal-table table-condensed"
                        data-toggle="table"
                        data-toolbar="#toolbar-ipaslog-on-modal"
                        data-show-columns="true"
                        {* 내보내기 *}
                        data-show-export="true"
                        data-export-types="['csv', 'excel']"
                        {*Row강조*}
                        data-row-style="ipasEventRowStyle"
                        {* 정렬 *}
                        data-sort-name="date"
                        data-sort-order="desc"
                        {* 페이징 *}
                        data-side-pagination="client"
                >
                    <thead>
                    <tr>
                        <th data-field="no">No</th>
                        <th data-field="date" data-formatter="dateFormatter">{{i18n .Lang "occurrence date"}}</th>
                        <th data-field="org_name">{{i18n .Lang "org"}}</th>
                        <th data-field="group_name" data-formatter="groupNameFormatter">{{i18n .Lang "group"}}</th>
                        <th data-field="event_type" data-formatter="ipaslogEventTypeFormatter">{{i18n .Lang "ipas.action"}}</th>
                        <th data-field="equip_id" data-formatter="ipasEquipIdFormatter">{{i18n .Lang "tag"}}</th>
                        <th data-field="targets" data-formatter="ipaslogTargetsFormatter">{{i18n .Lang "ipas.target"}}</th>
                        <th data-field="location" data-formatter="ipaslogLocationFormatter" data-align="center">{{i18n .Lang "location"}}</th>
                        <th data-field="latitude">{{i18n .Lang "latitude"}}</th>
                        <th data-field="longitude">{{i18n .Lang "longitude"}}</th>
                        <th data-field="distance" data-formatter="ipaslogDistanceFormatter">{{i18n .Lang "distance"}} (m)</th>
                        <th data-field="speed" data-formatter="ipaslogSpeedingFormatter">{{i18n .Lang "speed"}} <small>(km/h)</small></th>
                        <th data-field="snr" data-formatter="snrFormatter">SNR&nbsp;&nbsp;</th>
                        <th data-field="usim" data-visible="false">USIM</th>
                        <th data-field="ip" data-formatter="int2ipFormatter" data-visible="false">IP</th>
                        <th data-field="recv_date" data-formatter="dateFormatter" data-visible="false">{{i18n .Lang "receiv  ed date"}}</th>
                    </tr>
                    </thead>
                </table>
            </div><!-- modal-body-->
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">{{i18n .Lang "close"}}</button>
            </div>
        </div>
    </div>
</div> <!-- #modal-filter -->

<script>
    var $ipasLogModal  = $( "#modal-ipaslog" ), // Modal
        $ipasLogTable  = $ipasLogModal.find( "table.modal-table" ),
        ipasLogStorage = [], // 로그 저장소
        ipasTableKey   = getTableKey( $ipasLogModal, $ipasLogTable.attr( "id" ) ), // 테이블 고유키
        ipasLogPaging  = {
            no:                     0,      // 페이지 번호
            blockIndex:             0,      // 블럭 인덱스 (현재)
            blockIndexJustBefore:   -1,     // 블럭 인덱스 (이전)
            urlPrefix:              "",     // 검색 URL
            size:                   15,     // 페이지 크기
            blockSize:              20,      // 블럭 크기 (값이 3이면, 서버로부터 ipasLogPaging.size x 3 만큼 데이터를 미리 조회),
            sortBy:                 "date", // 정렬 기준
            orderBy:                "desc", // 정렬 순서
        };
    // restoreTableColumns( $ipasLogTable, ipasTableKey );

    $ipasLogTable.on( "column-switch.bs.table", function( e, field, checked ) { // 테이블 컬럼 보기/숨기기 속성이 변경되는 경우
        captureTableColumns( $( this ), ipasTableKey );
    });

    // 근거 로그 조회
    $( document ).on( "click", ".btn-show-ipaslog-on-modal", function() {
        var query = $( this ).data( "query" ),
            asset = $( "#select-assets" ).val();

        if ( asset !== undefined ) {
            var arr = asset.split( "/" ),
                param = {};

            if ( arr[0] > 0 ) {
                param.org_id = arr[0];
            }
            if ( arr[1] > 0 ) {
                param.group_id = arr[1];
            }
            query += '&' + $.param( param );
        }
        ipasLogPaging.urlPrefix = "/getIpasLogs?" + query + "&fast_paging=on";
        moveIpasLogPage( +1, ipasLogStorage, ipasLogPaging, $ipasLogModal, false );
        $ipasLogModal.modal( "show" );
    });

    // 페이지 이동 (고속페이징)
    $( ".btn-move-page", $ipasLogModal ).click(function( e ) {
        e.preventDefault();
        var direction = $( this ).data( "direction" );
        moveIpasLogPage( direction, ipasLogStorage, ipasLogPaging, $ipasLogModal, false );
    });

    $( "#modal-ipaslog" )
        .on( "shown.bs.modal", function( e ) {
            restoreTableColumns( $ipasLogTable, ipasTableKey );
        })
        .on( "hidden.bs.modal", function() {
            ipasLogPaging.no = 0;
            ipasLogPaging.blockIndex = 0;
            ipasLogPaging.blockIndexJustBefore = -1;
            ipasLogPaging.urlPrefix = "";

            // $( this ).bootstrapTable( "removeAll" );
            $( this ).bootstrapTable( "load", [] );
        });


    /**
     * 3. 함수
     *
     */

    // 페이지 이동(고속페이징)
    function moveIpasLogPage( direction, logs, paging, $modal, isRefresh ) {
        // 검색할 페이지 설정
        paging.no += direction;
        if (paging.no < 1) {
            paging.no = 1;
            return;
        }
        $( ".btn-page-text", $modal ).text( paging.no );

        // 테이블 찾기
        var $table = $modal.find( "table.modal-table" );

        // 페이징 컨트롤러
        paging.blockIndex = Math.floor( ( paging.no - 1 ) / paging.blockSize );
        if ( paging.blockIndex != paging.blockIndexJustBefore || isRefresh ) {
            var param = {
                offset: ( paging.size * paging.blockSize ) * paging.blockIndex,
                limit : paging.size * paging.blockSize,
                sort  : paging.sortBy,
                order : paging.orderBy
            };
            var url = paging.urlPrefix + "&" + $.param( param );

            // 데이터 조회
            $.ajax({
                type:  "GET",
                async: true,
                url:   url
            }).done( function( result ) {
                ipasLogStorage = result || []; // 값이 null 이면 크기0의 배열을 할당
                showTableData( $table, paging, ipasLogStorage );
                updateToolbarNav( $table, paging, ipasLogStorage.length );
            });
        } else {
            showTableData( $table, paging, logs );
            updateToolbarNav( $table, paging, ipasLogStorage.length );
        }

        paging.blockIndexJustBefore = paging.blockIndex;
    }

</script>