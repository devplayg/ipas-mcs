$(function() {

    /**
     * 1. 초기화
     *
     */
    var refreshInterval = 7 * 1000, // 새로고침 주기
        extraQuery = '';

    // 테이블 컬럼속성 복원
    $( ".table-data" ).each(function( idx, table ) {
        var tableKey =  getTableKey( $( this ), reqVars.ctrl );
        restoreTableColumns( $( this ), tableKey );
    });


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

    // 자산 (기관 / 그룹)
    var assets = { };
    $( "#select-groups" ).selectpicker( "hide" );
    initializeAssets();

    // 기관 선택
    $( "#select-orgs" ).change(function() {
        updateSelectGroups();
    });

    setInterval(function() {
        updateAllRanks();
    }, refreshInterval );


    // 기관/그룹 선택
    $( ".btn-apply" ).click(function( e ) {
        var org_id = $( "#select-orgs option:selected" ).map(function() {return $(this).val();}).get();
        var group_id = $( "#select-groups option:selected" ).map(function() {return $(this).val();}).get();

        extraQuery = $.param({
            org_id: org_id,
            group_id: group_id
        }, true );

        $( ".text-applied" ).removeClass( "hide" );
        setTimeout(function(){
            $( ".text-applied" ).addClass( "hide" );
        }, 1000);
    });


    /**
     * 3. 함수
     *
     */

    function updateAllRanks() {
        $( ".table-data" ).each(function( idx, table ) {
            $( this ).waitMe({
                effect: "win8",
                text: "Loading",
            });
            console.log(extraQuery);
            $( this ).bootstrapTable( "refresh", {
                url: $( this ).data( "url" ) + "&" + extraQuery,
                silent: true
            });
        });
    }


    // 자산 초기화
    function initializeAssets() {
        $.ajax({
            type  : "GET",
            async : true,
            url   : "/userassetclass/1/children"
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
                            value   :group.asset_id,
                            text    :group.name
                        })
                    );
                });
                assets[ org.type + "_" + org.asset_id ] = $optgroup;
            });


        }).always( function() {
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


});

