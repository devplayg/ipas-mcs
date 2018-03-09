$(function() {

    /**
     * 1. 초기화
     *
     */

    // 로그 테이블
    var logs        = [], // 로그 저장소(고속 페이징)
        $table      = $( "#table-log" ),
        tableKey    = getTableKey( $table, jsonvars.ctrl, jsonvars.act ); // 테이블 고유키

    // 로그 페이징 변수
    var paging = {
        no                 : 1,  // 페이지 번호
        size               : $table.data( "page-size" ), // 페이지 크기
        blockIndex         : 0,  // 블럭 인덱스 (현재)
        blockIndex_before  : -1, // 블럭 인덱스 (이전)
        blockSize          : 3  // 블럭 크기 (값이 3이면, 서버로부터 paging.size x 3 만큼 데이터를 한 번에 조회함)
    };

    // 날짜
    $( ".datetime" ).datetimepicker({
        format: "yyyy-mm-dd hh:ii",
        pickerPosition : "bottom-left",
        todayHighlight : 1,
        minView: 2,
        maxView: 4,
        autoclose: true
    });

    // 필터 유효성 체크

    $( "#form-filter" ).validate({
        submitHandler: function( form, e ) {
            e.preventDefault();

            if ( ! $( "input[name=fast_paging]", form ).is( ":checked" ) ) {
                $( form ).addHidden( "fast_paging", "off" );
            }
            form.submit();
        },
        ignore: "input[type='hidden']",
        errorClass: "help-block",
        rules: {
            Md5: {
                minlength: 10,
                maxlength: 32
            },
            SrcIpCidr: {
                ipv4_cidr: true
            },
            SrcPortStart: {
                min: 0,
                max: 65535
            },
            SrcPortEnd: {
                min: 0,
                max: 65535
            }
        },
        messages: {
            SrcPortStart: "0 ~ 65535",
            SrcPortEnd  : "0 ~ 65535",
        },
        highlight: function( element ) {
            $( element ).closest( ".form-group" ).addClass( "has-error" );
        },
        unhighlight: function( element ) {
            $( element ).closest( ".form-group" ).removeClass( "has-error" );
        }
    });


    // 테이블 컬럼속성 복원
    restoreTableColumns( $table, tableKey );

    // 자산 (센서 / 그룹 / IP Pool)
    var assets = { };
    $( "#select-folders, #select-ippools" ).selectpicker( "hide" );
    initializeAssets();

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
    resetMultiSelctedBoxesOfFilter();

    // 필터상태 업데이트
    updateFilterStatus();



    /**
     * 2. 이벤트
     *
     */

    // 페이지 이동 (고속페이징)
    $( ".btn-move-page" ).click(function(e) {
        e.preventDefault();
        movePage( $( this ).data( "direction" ), false);
    });

    // 필터 창 닫힘
    $( "#modal-filter" )
        .on( "hidden.bs.modal", function() { // 창이 닫힐 때
            var $form = $( this ).closest( "form" );
            $form.validate().resetForm();
            $form.get( 0 ).reset();
            resetMultiSelctedBoxesOfFilter();
        })
        .on( "shown.bs.modal", function(e) {
            var $form = $( this ).closest( "form" );
            $( "input[name=md5]", $form).focus().select();
        });


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

    }).on( "page-change.bs.table", function( e, number, size ) { // 일반 페이징 모드에서 페이지 크기가 변경되는 경우
        $( "#form-filter input[name=limit]" ).val ( size );

    }).on( "refresh.bs.table", function() { // 테이블 새로고침 이벤트 발생 시(고속 페이징)
        if ( $( "#form-filter input[name=fast_paging]" ).is( ":checked" ) ) {
            movePage( 0, true );
        }
    });





    /**
     * 3. 함수
     *
     */

    // 페이지 이동(고속페이징)
    function movePage( direction, is_refresh ) {
        paging.no += direction; // 검색할 페이지
        if (paging.no < 1) {
            paging.no = 1;
            return;
        }
        $( ".btn-page-text" ).text( paging.no );


        // 페이징 컨트롤러
        paging.blockIndex = Math.floor( (paging.no - 1) / paging.blockSize );
        if ( paging.blockIndex != paging.blockIndex_before || is_refresh ) {
            var param = {
                offset: ( paging.size * paging.blockSize ) * paging.blockIndex,
                limit: paging.size * paging.blockSize,
            };

            var url = "/ipaslog/getlogs/?" + $( "#form-filter :input" ).serialize() + "&" + $.param( param );
            console.log(url);

            // 데이터 조회
            console.log( 'Fetching' );
            $.ajax({
                type  : "GET",
                async : true,
                url   : url
            }).done( function( result ) {
                logs = result || []; // 값이 null 이면 크기0의 배열을 할당
                showTableData( $table, logs, paging );
                updatePagingNavButtons();
            });
        } else {
            showTableData( $table, logs, paging );
            updatePagingNavButtons();
        }

        paging.blockIndex_before = paging.blockIndex;
    }


    // 네비게이션 버튼 상태변경(고속 페이징)
    function updatePagingNavButtons( offset ) {
        var offset = (( paging.no - 1 ) % paging.blockSize ) * paging.size;
        if ( logs.length - offset < paging.size ) {
            $( ".btn-next" ).prop( "disabled", true );
        } else {
            $( ".btn-next" ).prop( "disabled", false );
        }

        if ( paging.no == 1 ) {
            $( ".btn-prev" ).prop( "disabled", true );
        } else {
            $( ".btn-prev" ).prop( "disabled", false );
        }
    }

    // 필터 상태
    function updateFilterStatus() {

        var fields = $( "#form-filter :input" )
            .not( "input[type='hidden'], [name='StartDate'], [name='EndDate'], [name='FastPaging'], [name='limit'], [name='score']" ) // 제외할 항목
            .serializeArray();

        // 항목에 조건값이 한 개 이상 설정되어 있으면
        $.each( fields, function( i, field ) {
            if ( field.value.length > 0 ) {
                $( ".icon-filter" ).removeClass( "hidden" );
                return;
            }
        });

        // 악성가능성 조회 범위 값이 0 ~ 100이 아니면
        if ( $( "#form-filter input[name=score]" ).val() != "0,100" ) {
            $( ".icon-filter" ).removeClass( "hidden" );
            return;
        }
    }

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
    function initializeAssets() {
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
        //         if ( $( "#form-filter input[name='fast_paging']" ).is( ":checked" ) ) {
                    movePage( 0, false ); // 첫 페이지 디스플레이
        //         }
        //     });
    }

    // 선택박스 초기설정
    function resetMultiSelctedBoxesOfFilter() {
        //     var cols = "category,filetype,trans_type,ext3";
        //     $.each(cols.split( "," ), function(i, c) {
        //         if ( jsonvars[c] !== undefined ) {
        //             $( "#form-filter select[name='" + c + "[]']" ).selectpicker( "val", jsonvars[c] ).selectpicker( "refresh" );
        //         }
        //     });
    }

});


// $(function() {
// /**
//  * 0. Test
//  *
//  */
// //$( "#modal-filter" ).modal( "show" );
// //$( "#modal-ipcard" ).modal( "show" );
//
//
// /**
//  * 1. Initialize
//  *
//  */
// $( ".form_datetime" ).datetimepicker({
//     pickerPosition : "bottom-left",
//     autoclose      : 1,
//     todayHighlight : 1,
// });
//
// // Table
// $( "#table-log" ).bootstrapTable();
// var paging = {
//     page   : 1,
//     limit  : $( "#table-log" ).data( "page-size" ),
//     offset : 0
// };
//
//
// // Load first page
// if ( $( "#filter input[name='FastPaging']" ).is( ":checked" ) ) {
//     movePage( 0 );
// }
//
// // Set boxes
// $( "#select-groups" ).selectpicker( "hide" );
// $( "#select-networks" ).selectpicker( "hide" );
//
// // Initialize assets
// var assets = { 0: "" },
//     asset_count = {
//         "sensor"  : 0,
//         "group"   : 0,
//         "network" : 0
//     };
// $.ajax({
//     type  : "GET",
//     async : true,
//     url   : "/myassets/1"
// }).done( function( myAssets ) {
//
//     // Sensor
//     $.each( myAssets, function( idx, sensor ) { // Sensor
//         $( "#select-sensors" ).append(
//             $( "<option>", {
//                 value: sensor.AssetId,
//                 text: sensor.Name,
//             })
//         );
//         asset_count.sensor++;
//
//         var $optgroup_group = $( "<optgroup>", {
//             label: sensor.Name,
//         });
//         $.each( sensor.children, function( i, group ) { // Group
//             $optgroup_group.append(
//                 $( "<option>", {
//                     value: group.AssetId,
//                     text: group.Name,
//                 })
//             );
//             asset_count.group++;
//
//             var $optgroup_network = $( "<optgroup>", {
//                 label: group.Name,
//             });
//             $.each( group.children, function( i, network ) { // Network
//                 $optgroup_network.append(
//                     $( "<option>", {
//                         value: network.AssetId,
//                         text: network.Name,
//                     })
//                 );
//                 asset_count.network++;
//             });
//             assets[ group.AssetId ] = $optgroup_network;
//         });
//         assets[ sensor.AssetId ] = $optgroup_group;
//     });
// }).always( function() {
//     // Selected sensors
//     $( "#select-sensors" ).selectpicker( "val", selected_sensors ).selectpicker( "refresh" );
//     if ( selected_sensors !== null && selected_sensors.length > 0 ) {
//         updateSelectGroups();
//     }
//
//     // Selected groups
//     $( "#select-groups" ).selectpicker( "val", selected_groups ).selectpicker( "refresh" );
//     if ( selected_groups !== null && selected_groups.length > 0 ) {
//         updateSelectNetworks();
//     }
//
//     // Selected networks
//     $( "#select-networks" ).selectpicker( "val", selected_networks ).selectpicker( "refresh" );
// });
// updateFilterStatus();
//
//
// /**
//  * 2. Events
//  *
//  */
// // Validation
// $( "#filter" ).validate({
//     submitHandler: function( form, e ) {
//         e.preventDefault();
//         form.submit();
//     },
//     ignore: "input[type='hidden']",
//     errorClass: "help-block",
//     rules: {
//         Md5: {
//             minlength: 10,
//             maxlength: 32
//         },
//         SrcIpCidr: {
//             ipv4_cidr: true
//         },
//         SrcPortStart: {
//             min: 0,
//             max: 65535
//         },
//         SrcPortEnd: {
//             min: 0,
//             max: 65535
//         }
//     },
//     messages: {
//         SrcPortStart: "0 ~ 65535",
//         SrcPortEnd  : "0 ~ 65535",
//     },
//     highlight: function( element ) {
//         $( element ).closest( ".form-group" ).addClass( "has-error" );
//     },
//     unhighlight: function( element ) {
//         $( element ).closest( ".form-group" ).removeClass( "has-error" );
//     }
// });
//
//
// // Page buttons' events
// $( ".btn-move-page" ).click(function(e) {
//     e.preventDefault();
//     var direction = $( this ).data( "direction" );
//     $( this ).prop( "disabled", true );
//
//     movePage(direction);
// });
//
//
// // Table events
// $( "#table-log" ).on( "load-success.bs.table", function ( e, data ) {
//     $( ".btn-move-page" ).prop( "disabled", false );
//
//     if ( $( "#table-log" ).bootstrapTable( "getData" ).length == 0
//             || $( "#table-log" ).bootstrapTable( "getData" ).length < $( "#table-log" ).data( "page-size" ) ) {
//         $( ".btn-next" ).prop( "disabled", true );
//     }
// });
//
//
// // Modal filter
// $( "#modal-filter" ).on( "hidden.bs.modal", function () {
//     // Reset and clear form
//     $( "#filter" ).validate().resetForm();
//     $( "#filter" ).get( 0 ).reset();
// });
//
//
// // Select boxes
// $( "#select-sensors" ).change(function () {
//     updateSelectGroups();
//     updateSelectNetworks();
// });
// $( "#select-groups" ).change(function () {
//     updateSelectNetworks();
// });
//
//
// /**
//  * 3. Functions
//  *
//  */
// function movePage(direction) {
//     paging.offset = paging.offset + ( paging.limit * direction );
//     if ( paging.offset < 0 ) {
//         paging.offset = 0;
//     }
//     $( ".btn-page-text" ).text( paging.offset / paging.limit + 1 );
//
//     var param = {
//         "StartDate"    : $( '#filter input[name="StartDate"]' ).val(),
//         "EndDate"      : $( '#filter input[name="EndDate"]' ).val(),
//         "Msg"          : $( '#filter input[name="Msg"]' ).val(),
//         "SrcIpCidr"    : $( '#filter input[name="SrcIpCidr"]' ).val(),
//         "SrcPortStart" : $( '#filter input[name="SrcPortStart"]' ).val(),
//         "SrcPortEnd"   : $( '#filter input[name="SrcPortEnd"]' ).val(),
//         "FastPaging"   : $( '#filter input[name="FastPaging"]' ).is( ":checked" ) ? "on" : "off",
//         "offset"       : paging.offset,
//         "limit"        : paging.limit
//     };
//     var url = "/filetranslog?" + $.param(param);
//     $('#table-log').bootstrapTable( 'refresh', { url: url } );
// }
//
//
// function updateFilterStatus() {
//     if (   ( selected_sensors !== null && selected_sensors.length > 0 )
//         || ( selected_groups !== null && selected_groups.length > 0 )
//         || ( selected_networks !== null && selected_networks.length > 0 )
//         || $( "#filter input[name=SrcIpCidr]" ).val().length > 0
//         || $( "#filter input[name=SrcPortStart]" ).val().length > 0
//         || $( "#filter input[name=SrcPortEnd]" ).val().length > 0
//     ) {
//         $( ".icon-filter" ).removeClass( "hidden" );
//     }
// }
//
//
// function updateSelectGroups() {
//     if ( $( "#select-sensors :selected" ).length > 0) {
//         $( "#select-groups" ).empty();
//         $( "#select-sensors :selected" ).map(function() {
//             var assetId = $( this ).val();
//             $( "#select-groups" ).append( assets[ assetId ] );
//         });
//         $( "#select-groups" ).selectpicker( "refresh" ).selectpicker( "show" );
//     } else {
//         $( "#select-groups" ).empty().selectpicker( "hide" );
//         $( "#select-networks" ).empty().selectpicker( "hide" );
//     }
// }
//
//
// function updateSelectNetworks() {
//     if ( $( "#select-groups :selected" ).length > 0) {
//         $( "#select-networks" ).empty();
//
//         $( "#select-groups :selected" ).map(function() {
//             var assetId = $( this ).val();
//             $( "#select-networks" ).append( assets[ assetId ] );
//         });
//         $( "#select-networks" ).selectpicker( "refresh" ).selectpicker( "show" );
//     } else {
//         $( "#select-networks" ).empty().selectpicker( "hide" );
//     }
// }
// });
