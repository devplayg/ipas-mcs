$(function() {

    /**
     * 1. 초기화
     *
     */
    // 로그 테이블
    var logs        = [], // 로그 저장소(고속 페이징)
        $table      = $( "#table-log" ),
        tableKey    = getTableKey( $table, reqVars.ctrl ); // 테이블 고유키

    // 로그 페이징 변수
    var paging = {
        no:                    1,                          // 페이지 번호
        size:                  $table.data( "page-size" ), // 페이지 크기
        blockIndex:            0,                          // 블럭 인덱스 (현재)
        blockIndexJustBefore: -1,                          // 블럭 인덱스 (이전)
        blockSize:             20                          // 블럭 크기 (값이 3이면, 서버로부터 paging.size x 3 만큼 데이터를 미리 조회)
    };

    // 날짜
    $( ".datetime" ).datetimepicker({
        format:         "yyyy-mm-dd hh:ii",
        pickerPosition: "bottom-left",
        todayHighlight: 1,
        minView:        2,
        maxView:        4,
        autoclose:      true
    });

    // 필터 유효성 체크
    $( "#form-filter" ).validate({
        submitHandler: function( form, e ) {
            e.preventDefault();

            if ( ! $( "input[name=fast_paging]", form ).is( ":checked" ) ) {
                $( form ).addHidden( "fast_paging", "off" );
            }
            $( form ).addHidden( "sort", $table.bootstrapTable("getOptions").sortName );
            $( form ).addHidden( "order", $table.bootstrapTable("getOptions").sortOrder );

            form.submit();
        },
        ignore: "input[type='hidden']",
        errorClass: "help-block",
        rules: {
            guid: {
                minlength: 2,
                maxlength: 5
            },
        },
        messages: {
            SrcPortStart: "0 ~ 65535",
            SrcPortEnd:   "0 ~ 65535",
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

    // 자산 (기관 / 그룹)
    var assets = { };
    $( "#select-groups" ).selectpicker( "hide" );
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
    $( ".btn-move-page" ).click(function( e ) {
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
        .on( "shown.bs.modal", function( e ) {
            var $form = $( this ).closest( "form" );
            $( "input[name=md5]", $form).focus().select();
        });



    // 기관 선택
    $( "#select-orgs" ).change(function() {
        updateSelectGroups();
    });
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
    function movePage( direction, isRefresh ) {
        paging.no += direction; // 검색할 페이지
        if (paging.no < 1) {
            paging.no = 1;
            return;
        }
        $( ".btn-page-text" ).text( paging.no );

        // 페이징 컨트롤러
        paging.blockIndex = Math.floor( ( paging.no - 1 ) / paging.blockSize );
        if ( paging.blockIndex != paging.blockIndexJustBefore || isRefresh ) {
            var param = {
                offset: ( paging.size * paging.blockSize ) * paging.blockIndex,
                limit : paging.size * paging.blockSize,
                sort  : $table.bootstrapTable( "getOptions" ).sortName,
                order : $table.bootstrapTable( "getOptions" ).sortOrder
            };

            var url = "/getIpasLogs?" + $( "#form-filter :input[name!=limit]" ).serialize() + "&" + $.param( param );

            // 데이터 조회
            // console.log(url);
            // console.log( $( "#form-filter" ).serializeArray());
            // console.log( $( "#form-filter :input[name!='limit']" ).serializeArray()    );
            // console.log( $( "#form-filter" ).not("#form-filter input[name=limit]").serializeArray()    );
            console.log( 'Fetching' );
            $.ajax({
                type:  "GET",
                async: true,
                url:   url
            }).done( function( result ) {
                logs = result || []; // 값이 null 이면 크기0의 배열을 할당
                // console.log(logs);
                showTableData( $table, paging, logs );
                updateToolbarNav( $table, paging, logs.length );
            });
        } else {
            showTableData( $table, paging, logs );
            updateToolbarNav( $table, paging, logs.length );
        }

        paging.blockIndexJustBefore = paging.blockIndex;
    }


    // 네비게이션 버튼 상태변경(고속 페이징)
    // function updatePagingNavButtons() {
    //     var offset = (( paging.no - 1 ) % paging.blockSize ) * paging.size;
    //     if ( logs.length - offset < paging.size ) {
    //         $( ".btn-next" ).prop( "disabled", true );
    //     } else {
    //         $( ".btn-next" ).prop( "disabled", false );
    //     }
    //
    //     if ( paging.no == 1 ) {
    //         $( ".btn-prev" ).prop( "disabled", true );
    //     } else {
    //         $( ".btn-prev" ).prop( "disabled", false );
    //     }
    // }

    // 필터 상태
    function updateFilterStatus() {
        var fields = $( "#form-filter :input" )
            .not( "input[type='hidden'], [name='start_date'], [name='end_date'], [name='fast_paging'], [name='limit']" ) // 제외할 항목
            .serializeArray();

        // 항목에 조건값이 한 개 이상 설정되어 있으면
        $.each( fields, function( i, field ) {
            if ( field.value.length > 0 ) {
                $( ".icon-filter" ).removeClass( "hidden" );
                return;
            }
        });
    }

    // 그룹 업데이트
    function updateSelectGroups() {
        if ( $( "#select-orgs :selected" ).length > 0) {
            $( "#select-groups" ).empty();
            $( "#select-orgs :selected" ).map(function() {
                var asset_id = $( this ).val();
                $( "#select-groups" ).append( assets[ "1_" + asset_id ] );
            });
            $( "#select-groups" ).selectpicker( "refresh" ).selectpicker( "show" );
        } else {
            $( "#select-groups" ).empty().selectpicker( "hide" );
        }
    }

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

    // 자산 초기화
    function initializeAssets() {
        $.ajax({
            type:  "GET",
            async: true,
            url:   "/userassetclass/1/children"
        }).done( function( result ) {
            // 기관
            $.each( result, function( idx, org ) {
                $( "#select-orgs" ).append(
                    $( "<option>", {
                        value: org.asset_id,
                        text: org.name
                    })
                );

                var $optgroup = $( "<optgroup>", {
                    label: org.name
                });

                // 그룹
                $.each( org.children, function( i, group ) {
                    $optgroup.append(
                        $( "<option>", {
                            value: group.asset_id,
                            text:  group.name
                        })
                    );
                });
                assets[ org.type + "_" + org.asset_id ] = $optgroup;
            });


        }).always( function() {
            // 기관 선택
            if ( reqVars.org_id !== undefined && reqVars.org_id.length > 0 ) {
                $( "#select-orgs" ).selectpicker( "val", reqVars.org_id ).selectpicker( "refresh" );
                updateSelectGroups();

            } else {
                $( "#select-orgs" ).selectpicker( "refresh" );
            }

            // 그룹 선택
            if ( reqVars.group_id !== undefined && reqVars.group_id.length > 0 ) {
                $( "#select-groups" ).selectpicker( "val", reqVars.group_id ).selectpicker( "refresh" );
            }

            // 빠른 페이징일 때는
            if ( $( "#form-filter input[name='fast_paging']" ).is( ":checked" ) ) {
                movePage( 0, false ); // 첫 페이지 디스플레이
            }
        });
    }

    // 선택박스 초기설정
    function resetMultiSelctedBoxesOfFilter() {
        var cols = "event_type";
        $.each(cols.split( "," ), function(i, c) {
            if ( reqVars[c] !== undefined ) {
                $( "#form-filter select[name='" + c + "']" ).selectpicker( "val", reqVars[c] ).selectpicker( "refresh" );
            }
        });
    }

});

