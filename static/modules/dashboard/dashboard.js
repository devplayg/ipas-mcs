$(function() {

    /**
     * 1. 초기화
     *
     */
    // 자산 (기관 / 그룹)
    var assets = { },
        interval = 60000000,
        timer = null;

    // 날짜
    $( ".datetime" ).datetimepicker({
        format: "yyyy-mm-dd",
        pickerPosition : "bottom-left",
        todayHighlight : 1,
        minView: 2,
        maxView: 4,
        autoclose: true
    });

    initializeAssets();
    updateStats();
    startTimer();
    // $( "#modal-ipaslog" ).modal( "show" );

    /**
     * 2. 이벤트
     *
     */

    // 자산 선택
    $( "#select-assets" ).on( "change", function() {
        stopTimer();
        updateStats();
        startTimer();
    });


    $( ".btn-start" ).click(function(e) {
        if ( timer === null ) {
            startTimer();
        } else {
            stopTimer();
        }
    });

    $( ".activity" ).change(function() {
        var asset = $( "#select-assets :selected" ).val().split( "/", 2 ),
            orgId = asset[0],
            groupId = asset[1];
        updateLogs( orgId, groupId );
    });




    /**
     * 3. 함수
     *
     */

    // 자산 초기화
    function initializeAssets() {
        $.ajax({
            type  : "GET",
            async : true,
            url   : "/userassetclass/1/children"
        }).done( function( result ) {
            // 기관
            $.each( result, function( idx, org ) {
                $( "#select-assets" ).append(
                    $( "<option>", {
                        value: org.asset_id + "/-1",
                        text: org.name
                    })
                );

                // 그룹
                $.each( org.children, function( i, group ) {
                    $( "#select-assets" ).append(
                        $( "<option>", {
                            value: org.asset_id + "/" + group.asset_id,
                            text: "- " + group.name,
                            class: "ml20"
                        })
                    );
                });
            });


        }).always( function() {
            $( "#select-assets" ).selectpicker( "refresh" );
        });
    }

    function startTimer() {  // use a one-off timer
        $( ".btn-start" ).removeClass( "default" ).addClass( "blue" );
        $( ".btn-start .text" ).html( "<i class='fa fa-circle-o-notch fa-spin'></i>" );
        timer = setInterval(updateStats, interval);
    }

    function stopTimer() {
        $( ".btn-start" ).removeClass( "blue" ).addClass( "default" );
        $( ".btn-start .text" ).html( "<i class='fa fa-play'></i>" );
        clearTimeout( timer );
        timer = null;
    }


    function updateStats() {
        var asset = $( "#select-assets :selected" ).val().split( "/", 2 ),
            orgId = asset[0],
            groupId = asset[1];

        updateSummary( orgId, groupId );
        updateRankings( orgId, groupId );
        updateLogs( orgId, groupId );

        $( ".text-updated" ).removeClass( "hide" );
        setTimeout(function(){ $( ".text-updated" ).addClass( "hide" ); }, 500);
    }


    function updateSummary( orgId, groupId ) {
        var url = "/stats/summary/org/" + orgId + "/group/" + groupId;
        $.ajax({
            type  : "GET",
            async : true,
            url   : url
        }).done( function( r ) {
            $( ".count-startup" ).text( r.eventTypes[1] );
            $( ".count-shock" ).text( r.eventTypes[2] );
            $( ".count-speeding" ).text( r.eventTypes[3] );
            $( ".count-proximity" ).text( r.eventTypes[4] );
            $( ".count-pt" ).text( r.equipCountByType[1] );
            $( ".count-zt" ).text( r.equipCountByType[2] );
            $( ".count-vt" ).text( r.equipCountByType[4] );
        }).always( function() {
        });
    }


    function updateRankings( orgId, groupId ) {
        $( ".table-ranking" ).each(function( idx, obj ) {
            var url = $( this ).data( "query" ) + "/org/" + orgId + "/group/" + groupId;
            $( this ).bootstrapTable( "refresh", {
                url: url,
                silent: true
            });
        });
    }


    function updateLogs( orgId, groupId ) {
        var url = '/getRealTimeLogs?limit=5';

        var activities = [];
        $( ".activity" ).each(function( idx, obj ) {
            if ( $( obj ).is(":checked") ) {
                activities.push( "event_type=" + $( obj ).val() );
            }
        });

        if ( activities.length < 1 ) {
            $( "#table-ipaslogs" ).bootstrapTable( "removeAll" );
        } else {
            var urlSuffix = "";
            if ( orgId > 0 ) {
                urlSuffix += "&org_id=" + orgId;
                if ( groupId > 0 ) {
                    urlSuffix += "&group_id=" + groupId;
                }
            }
            $( "#table-ipaslogs" ).bootstrapTable( "refresh", {
                url: url + "&" + activities.join( "&" ) + urlSuffix,
                silent: true
            });
        }
    }

});

