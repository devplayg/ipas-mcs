$(function() {

    /**
     * 1. 초기화
     *
     */
        // 로그 테이블
    var logs        = [], // 로그 저장소(고속 페이징)
        $table      = $( "#table-event1" ),
        tableKey    = getTableKey( $table, reqVars.ctrl ); // 테이블 고유키


    /**
     * 2. 이벤트
     *
     */
    // 테이블 이벤트
    $table.on( "column-switch.bs.table", function( e, field, checked ) { // 테이블 컬럼 보기/숨기기 속성이 변경되는 경우
        captureTableColumns( $( this ), tableKey );

    // }).on( "page-change.bs.table", function( e, number, size ) { // 일반 페이징 모드에서 페이지 크기가 변경되는 경우
    //     $( "#form-filter input[name=limit]" ).val ( size );

    }).on( "load-success.bs.table", function() { // 테이블 새로고침 이벤트 발생 시(고속 페이징)
        $("body").tooltip({
            selector: '.tooltips'
        });
    });



    /**
     * 3. 함수
     *
     */


});

