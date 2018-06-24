<!-- 상세필터 -->
<div id="modal-ipasstatus" class="modal fade" tabindex="-1" role="dialog">
    <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4 class="modal-title"><i class="icon-list"></i> IPAS</h4>
            </div>

            <div class="modal-body">
                <div id="toolbar-ipasstatus-on-modal">
                    <button type="button" class="btn blue btn-move-page btn-prev" data-direction="-1" data-loading-text="&lt;">&lt;</button>
                    <button type="button" class="btn blue btn-move-page btn-page-text" data-direction="0">1</button>
                    <button type="button" class="btn blue btn-move-page btn-next" data-direction="1" data-loading-text="&gt;">&gt;</button>
                </div>
                <table  id="table-ipasstatus-on-modal"
                        class="modal-table table-condensed"
                        data-toggle="table"
                        data-toolbar="#toolbar-ipasstatus-on-modal"
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
                        <th data-field="org_name">{{i18n .Lang "org"}}</th>
                        <th data-field="group_name" data-formatter="groupNameFormatter">{{i18n .Lang "group"}}</th>
                        <th data-field="equip_type" data-formatter="ipasEquipTypeFormatter" data-sortable="true">{{i18n .Lang "equip type"}}</th>
                        <th data-field="equip_id" data-formatter="ipasEquipIdFormatter" data-sortable="true">{{i18n .Lang "tag"}}</th>
                        <th data-field="usim" data-visible="true">USIM</th>
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
{{/*<script src="/static/modules/ipasstatus/formatter.js"></script>*/}}

<script>
    var $ipasStatusModal  = $( "#modal-ipasstatus" ), // Modal
        $ipasStatusTable  = $ipasStatusModal.find( "table.modal-table" ),
        ipasStatusStorage = [], // 로그 저장소
        ipasTableKey   = getTableKey( $ipasStatusModal, $ipasStatusTable.attr( "id" ) ), // 테이블 고유키
        ipasStatusPaging  = {
            no:                     0,      // 페이지 번호
            blockIndex:             0,      // 블럭 인덱스 (현재)
            blockIndexJustBefore:   -1,     // 블럭 인덱스 (이전)
            urlPrefix:              "",     // 검색 URL
            size:                   15,     // 페이지 크기
            blockSize:              20,      // 블럭 크기 (값이 3이면, 서버로부터 ipasStatusPaging.size x 3 만큼 데이터를 미리 조회),
            sortBy:                 "equip_id", // 정렬 기준
            orderBy:                "asc",      // 정렬 순서
        };
    $ipasStatusTable.on( "column-switch.bs.table", function( e, field, checked ) { // 테이블 컬럼 보기/숨기기 속성이 변경되는 경우
        captureTableColumns( $( this ), ipasTableKey );
    });

    // 근거 로그 조회
    $( document ).on( "click", ".btn-show-ipasstatus-on-modal", function() {
        // var asset = $( "#select-assets" ).val();
        var query = $( this ).data( "query" ),
            asset = $( "#select-assets" ).val();

        if ( asset !== undefined ) {
            var arr = asset.split( "/" ).map( function(x) { return parseInt(x) } );

            if ( arr[0] > 0 ) { // 기관 조회
                if ( arr[1] > 0 ) { // 그룹 조회
                    ipasStatusPaging.urlPrefix = "/ipasgroup/" + arr[1];
                } else { // 기관 조회
                    ipasStatusPaging.urlPrefix = "/ipasorg/" + arr[0];
                }
            } else { // 전체 조회
                ipasStatusPaging.urlPrefix = "/ipasorg/0";
            }
            ipasStatusPaging.urlPrefix += "?fast_paging=on&" + query;

        //             param = {};
        //
        //     if ( arr[0] > 0 ) {
        //         param.org_id = arr[0];
        //     }
        //     if ( arr[1] > 0 ) {
        //         param.group_id = arr[1];
        //     }
        //     query += '&' + $.param( param );
            // console.log(query);
            // ipasStatusPaging.urlPrefix = "/getIpasLogs?" + query;
            movePage( +1, ipasStatusStorage, ipasStatusPaging, $ipasStatusModal, false );
            $ipasStatusModal.modal( "show" );
        }
    });

    // 페이지 이동 (고속페이징)
    $( ".btn-move-page", $ipasStatusModal ).click(function( e ) {
        e.preventDefault();
        var direction = $( this ).data( "direction" );
        movePage( direction, ipasStatusStorage, ipasStatusPaging, $ipasStatusModal, false );
    });

    $( "#modal-ipasstatus" )
            .on( "shown.bs.modal", function( e ) {
                restoreTableColumns( $ipasStatusTable, ipasTableKey );
            })
            .on( "hidden.bs.modal", function() {
                ipasStatusPaging.no = 0;
                ipasStatusPaging.blockIndex = 0;
                ipasStatusPaging.blockIndexJustBefore = -1;
                ipasStatusPaging.urlPrefix = "";

                // $( this ).bootstrapTable( "removeAll" );
                $( this ).bootstrapTable( "load", [] );
            });


    /**
     * 3. 함수
     *
     */

    // 페이지 이동(고속페이징)
    function movePage( direction, logs, paging, $modal, isRefresh ) {
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
            // console.log(url);
            // 데이터 조회
            $.ajax({
                type:  "GET",
                async: true,
                url:   url
            }).done( function( result ) {
                ipasStatusStorage = result || []; // 값이 null 이면 크기0의 배열을 할당
                showTableData( $table, paging, ipasStatusStorage );
                updateToolbarNav( $table, paging, ipasStatusStorage.length );
            });
        } else {
            showTableData( $table, paging, logs );
            updateToolbarNav( $table, paging, ipasStatusStorage.length );
        }

        paging.blockIndexJustBefore = paging.blockIndex;
    }
</script>