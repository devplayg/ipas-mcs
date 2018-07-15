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

            var url = "/security/log?" + $( "#form-filter :input[name!=limit]" ).serialize() + "&" + $.param( param );

            // 데이터 조회
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

    // 선택박스 초기설정
    function resetMultiSelctedBoxesOfFilter() {
        var cols = "category";
        $.each(cols.split( "," ), function(i, c) {
            if ( reqVars[c] !== undefined ) {
                $( "#form-filter select[name='" + c + "']" ).selectpicker( "val", reqVars[c] ).selectpicker( "refresh" );
            }
        });
    }

});

