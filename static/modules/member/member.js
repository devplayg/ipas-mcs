$(function() {

    /**
     * 1. 초기화
     *
     */
    //
    //     // 로그 테이블
    // var logs        = [], // 로그 저장소(고속 페이징)
    //     $table      = $( "#table-log" ),
    //     tableKey    = getTableKey( $table, reqVars.ctrl, reqVars.act ); // 테이블 고유키
    //
    // // 로그 페이징 변수
    // var paging = {
    //     no                 : 1,                          // 페이지 번호
    //     size               : $table.data( "page-size" ), // 페이지 크기
    //     blockIndex         : 0,                          // 블럭 인덱스 (현재)
    //     blockIndex_before  : -1,                         // 블럭 인덱스 (이전)
    //     blockSize          : 3                           // 블럭 크기 (값이 3이면, 서버로부터 paging.size x 3 만큼 데이터를 미리 조회)
    // };

    // 테이블 컬럼속성 복원
    restoreTableColumns( $table, tableKey );

    // 자산 (센서 / 그룹 / IP Pool)
    // var assets = { };
    // $( "#select-folders, #select-ippools" ).selectpicker( "hide" );
    // initializeAssets();

    // // 악성가능성
    // var score = {
    //     min: ( jsonvars.score_min === undefined ) ?   0 : parseInt( jsonvars.score_min ),
    //     max: ( jsonvars.score_max === undefined ) ? 100 : parseInt( jsonvars.score_max ),
    // };
    // $( "#form-filter input[name=score]" ).bootstrapSlider({
    //     min     : 0,
    //     max     : 100,
    //     range   : true,
    //     step    : 10,
    //     value   : [ score.min, score.max ]
    // });
    // updateScore( score.min, score.max );

    // 선택박스 초기화
    // resetMultiSelctedBoxesOfFilter();

    // 필터상태 업데이트
    updateFilterStatus();



    /**
     * 2. 이벤트
     *
     */
    //
    // // 페이지 이동 (고속페이징)
    // $( ".btn-move-page" ).click(function(e) {
    //     e.preventDefault();
    //     movePage( $( this ).data( "direction" ), false);
    // });

    // // 필터 창 닫힘
    // $( "#modal-filter" )
    //     .on( "hidden.bs.modal", function() { // 창이 닫힐 때
    //         var $form = $( this ).closest( "form" );
    //         $form.validate().resetForm();
    //         $form.get( 0 ).reset();
    //         resetMultiSelctedBoxesOfFilter();
    //     })
    //     .on( "shown.bs.modal", function(e) {
    //         var $form = $( this ).closest( "form" );
    //         $( "input[name=md5]", $form).focus().select();
    //     });
    //

    // // 센서 선택
    // $( "#select-sensors" ).change(function() {
    //     updateSelectFolders();
    //     updateSelectIppools();
    // });
    //
    // // 그룹 선택
    // $( "#select-folders" ).change(function() {
    //     updateSelectIppools();
    // });
    //
    // // 악성가능성 변경
    // $( "#form-filter input[name=score]" ).on( "change", function( slideEvt ) {
    //     updateScore( slideEvt.value.newValue[0], slideEvt.value.newValue[1] );
    // });
    //
    // 테이블 이벤트
    $table.on( "column-switch.bs.table", function( e, field, checked ) { // 테이블 컬럼 보기/숨기기 속성이 변경되는 경우
        captureTableColumns( $( this ), tableKey );

    // }).on( "page-change.bs.table", function( e, number, size ) { // 일반 페이징 모드에서 페이지 크기가 변경되는 경우
    //     $( "#form-filter input[name=limit]" ).val ( size );

    // }).on( "refresh.bs.table", function() { // 테이블 새로고침 이벤트 발생 시(고속 페이징)
    //     if ( $( "#form-filter input[name=fastPaging]" ).is( ":checked" ) ) {
    //         movePage( 0, true );
    //     }
    // }).on( "sort.bs.table", function ( e, name, order ) {
    //     //console.log(name + " / " + order)
    //     $( "#form-filter input[name=sort]" ).val ( name );
    //     $( "#form-filter input[name=order]" ).val ( order );
    });





    /**
     * 3. 함수
     *
     */



    // // 그룹 업데이트
    // function updateSelectFolders() {
    //     if ( $( "#select-sensors :selected" ).length > 0) {
    //         $( "#select-folders" ).empty();
    //         $( "#select-sensors :selected" ).map(function() {
    //             var asset_id = $( this ).val();
    //             $( "#select-folders" ).append( assets[ "1_" + asset_id ] );
    //         });
    //         $( "#select-folders" ).selectpicker( "refresh" ).selectpicker( "show" );
    //     } else {
    //         $( "#select-folders" ).empty().selectpicker( "hide" );
    //         $( "#select-ippools" ).empty().selectpicker( "hide" );
    //     }
    // }
    //
    // // IP Pool 업데이트
    // function updateSelectIppools() {
    //     if ( $( "#select-folders :selected" ).length > 0) {
    //         $( "#select-ippools" ).empty();
    //         $( "#select-folders :selected" ).map(function() {
    //             var asset_id = $( this ).val();
    //             $( "#select-ippools" ).append( assets[ "2_" + asset_id ] );
    //         });
    //         $( "#select-ippools" ).selectpicker( "refresh" ).selectpicker( "show" );
    //     } else {
    //         $( "#select-ippools" ).empty().selectpicker( "hide" );
    //     }
    // }
    //
    // // 악성가능성 업데이트
    // function updateScore( min, max ) {
    //     $( "#form-filter input[name=score_min]" ).val( min );
    //     $( "#form-filter input[name=score_max]" ).val( max );
    //     $( ".score_min" ).text( min );
    //     $( ".score_max" ).text( max );
    // }
    //
    // 자산 초기화
    // function initializeAssets() {
        //     $.ajax({
        //         type  : "GET",
        //         async : true,
        //         url   : "?mod=asset&act=procGetAssets"
        //     }).done( function( result ) {
        //         if ( result.state ) {
        //             // 센서
        //             $.each( result.data, function( idx, sensor ) {
        //                 $( "#select-sensors" ).append(
        //                     $( "<option>", {
        //                         value: sensor.asset_id,
        //                         text: sensor.name
        //                     })
        //                 );
        //
        //                 var $optgroup_folder = $( "<optgroup>", {
        //                     label: sensor.name
        //                 });
        //
        //                 // 그룹
        //                 $.each( sensor.children, function( i, folder ) {
        //                     $optgroup_folder.append(
        //                         $( "<option>", {
        //                             value   :folder.asset_id,
        //                             text    :folder.name
        //                         })
        //                     );
        //
        //                     var $optgroup_ippool = $( "<optgroup>", {
        //                         label: folder.name
        //                     });
        //
        //                     // IP Pool
        //                     $.each( folder.children, function( i, ippool ) {
        //                         $optgroup_ippool.append(
        //                             $( "<option>", {
        //                                 value   :ippool.asset_id,
        //                                 text    :ippool.name
        //                             })
        //                         );
        //                     });
        //                     assets[ folder.asset_type + "_" + folder.asset_id ] = $optgroup_ippool;
        //                 });
        //                 assets[ sensor.asset_type + "_" + sensor.asset_id ] = $optgroup_folder;
        //             });
        //         } else {
        //             var msg = 'Error';
        //             if (result.state == __FAIL__) {
        //                 if (result.db !== undefined) {
        //                     msg = result.db.error;
        //                 } else if (result.message !== undefined) {
        //                     msg = result.message;
        //                 }
        //             }
        //
        //             swal({ title : msg, type  : "warning" });
        //         }
        //
        //     }).always( function() {
        //         // Selected sensors
        //         if ( jsonvars.sensors !== undefined && jsonvars.sensors.length > 0 ) {
        //             $( "#select-sensors" ).selectpicker( "val", jsonvars.sensors ).selectpicker( "refresh" );
        //             updateSelectFolders();
        //         } else {
        //             $( "#select-sensors" ).selectpicker( "refresh" );
        //         }
        //
        //         // Selected folders
        //         if ( jsonvars.folders !== undefined && jsonvars.folders.length > 0 ) {
        //             $( "#select-folders" ).selectpicker( "val", jsonvars.folders ).selectpicker( "refresh" );
        //             updateSelectIppools();
        //         }
        //
        //         // Selected ippools
        //         if ( jsonvars.ippools !== undefined && jsonvars.ippools.length > 0 ) {
        //             $( "#select-ippools" ).selectpicker( "val", jsonvars.ippools ).selectpicker( "refresh" );
        //         }
        //
        //         // 빠른 페이징일 때는
        // if ( $( "#form-filter input[name='fastPaging']" ).is( ":checked" ) ) {
        //     movePage( 0, false ); // 첫 페이지 디스플레이
        // }
        //     });
    // }
    //
    // // 선택박스 초기설정
    // function resetMultiSelctedBoxesOfFilter() {
    //     var cols = "risk_level[]";
    //     $.each(cols.split( "," ), function(i, c) {
    //         if ( reqVars[c] !== undefined ) {
    //             $( "#form-filter select[name='" + c + "']" ).selectpicker( "val", reqVars[c] ).selectpicker( "refresh" );
    //         }
    //     });
    // }

});

