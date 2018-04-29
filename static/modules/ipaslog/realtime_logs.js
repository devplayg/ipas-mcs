$(function() {

    /**
     * 1. 초기화
     *
     */
        // 로그 테이블
    var logs        = [],           // 로그 저장소(고속 페이징)
        refreshInterval = 7 * 1000; // 새로고침 주기

    //
    /**
     * 2. 이벤트
     *
     */
    // 테이블 이벤트
    $( ".table-data" ).on( "column-switch.bs.table", function( e, field, checked ) { // 테이블 컬럼 보기/숨기기 속성이 변경되는 경우
        var tableKey =  getTableKey( $( this ), reqVars.ctrl );
        captureTableColumns( $( this ), tableKey );

    }).on( "load-success.bs.table", function() { // 테이블 새로고침 이벤트 발생 시(고속 페이징)
        var id = $( this ).data( "url" );
        $( this ).waitMe( "hide" );

        $( "body" ).tooltip({
            selector: '.tooltips'
        });
    });




    var timer = setInterval(function() {
        updateAllRanks();
    }, refreshInterval );
    /**
     * 3. 함수
     *
     */

    function updateAllRanks() {
        $( ".table-data" ).each(function( idx, table ) {
            $( this ).waitMe({
                effect : 'bounce',
                text : 'Please wait..',
            });
            $( this ).bootstrapTable( "refresh", { silent: true });
        });
    }


});

