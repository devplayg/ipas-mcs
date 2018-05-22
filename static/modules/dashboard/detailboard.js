$(function() {

    /**
     * 1. 초기화
     *
     */
    // 자산 (기관 / 그룹)
    // var assets = { },
    //     interval = 211000,
    //     timer = null;
    //
    // // // 날짜
    // // $( ".datetime" ).datetimepicker({
    // //     format: "yyyy-mm-dd",
    // //     pickerPosition : "bottom-left",
    // //     todayHighlight : 1,
    // //     minView: 2,
    // //     maxView: 4,
    // //     autoclose: true
    // // });
    //
    // Chart - 이벤트 유형
    var eventTypeChart = Morris.Donut({
        element: 'chart-eventType',
        colors: [ ShockColor, SpeedingColor, ProximityColor ],

        resize: true,
        data: [
            { label: "N/A", value: 0 }
        ],
        formatter: function (x) { return x + ""}
    }).on('click', function(i, row){
        console.log(i, row);
    });

    initializeAssets();
    updateStats();
    // startTimer();
    //
    /**
     * 2. 이벤트
     *
     */

    // 자산 선택
    $( "#select-assets" ).on( "change", function() {
    //     stopTimer();
        updateStats();
    //     startTimer();
    });


    // $( ".btn-start" ).click(function(e) {
    //     // e.preventDefault();
    //     if ( timer === null ) {
    //         startTimer();
    //     } else {
    //         stopTimer();
    //     }
    // });
    //
    // $( ".activity" ).change(function() {
    //     var asset = $( "#select-assets :selected" ).val().split( "/", 2 ),
    //         orgId = asset[0],
    //         groupId = asset[1];
    //     updateLogs( orgId, groupId );
    // });
    //



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

    // function startTimer() {  // use a one-off timer
    //     $( ".btn-start" ).removeClass( "default" ).addClass( "blue" );
    //     $( ".btn-start .text" ).html( "<i class='fa fa-circle-o-notch fa-spin'></i>" );
    //     timer = setInterval(updateStats, interval);
    // }
    //
    // function stopTimer() {
    //     $( ".btn-start" ).removeClass( "blue" ).addClass( "default" );
    //     $( ".btn-start .text" ).html( "<i class='fa fa-play'></i>" );
    //     clearTimeout( timer );
    //     timer = null;
    // }
    //
    //
    function updateStats() {
        var asset = $( "#select-assets :selected" ).val().split( "/", 2 ),
            orgId = asset[0],
            groupId = asset[1];

        updateSummary( orgId, groupId );
        updateEventTags( orgId, groupId );
    //     updateRankings( orgId, groupId );
    //     updateLogs( orgId, groupId );
    //
    //     $( ".text-updated" ).removeClass( "hide" );
    //     setTimeout(function(){ $( ".text-updated" ).addClass( "hide" ); }, 500);
    }


    function updateSummary( orgId, groupId ) {
        var url = "/stats/summary/org/" + orgId + "/group/" + groupId;
        $.ajax({
            type  : "GET",
            async : true,
            url   : url
        }).done( function( r ) {
            $( ".count-startup" ).text( r.eventTypes[StartupEvent] );
            $( ".count-shock" ).text( r.eventTypes[ShockEvent] );
            $( ".count-speeding" ).text( r.eventTypes[SpeedingEvent] );
            $( ".count-proximity" ).text( r.eventTypes[ProximityEvent] );

            $( ".count-pt" ).text( r.equipCountByType[PT] );
            $( ".count-zt" ).text( r.equipCountByType[ZT] );
            $( ".count-vt" ).text( r.equipCountByType[VT] );
            // $( ".count-total-tags" ).text( r.equipCountByType[PT] + r.equipCountByType[ZT] + r.equipCountByType[VT] );
    //
            var  data = [
                {value:  r.eventTypes[ShockEvent], label: 'SHOCK'},
                {value:  r.eventTypes[SpeedingEvent], label: 'SPEEDING'},
                {value:  r.eventTypes[ProximityEvent], label: 'PROXIMITY'},
            ];

            var total = r.eventTypes[ShockEvent] + r.eventTypes[SpeedingEvent] + r.eventTypes[ProximityEvent];
            if ( total === 0 ) {
                total = 1;
            }
            $( "#pgb-shock" ).css( "width", (r.eventTypes[ShockEvent] / total * 100) + "%" );
            $( "#pgb-speeding" ).css( "width", (r.eventTypes[SpeedingEvent] / total * 100) + "%" );
            $( "#pgb-proximity" ).css( "width", (r.eventTypes[ProximityEvent] / total * 100) + "%" );
            // $('#progress-upload .progress-bar').css('width', progress + '%');
            if ( total > 0 ) {
                eventTypeChart.setData( data );
            } else {
                eventTypeChart.setData( [ { value: 0, label: 'N/A' } ] );
            }
        }).always( function() {
        });
    }
    //
    //
    // function updateRankings( orgId, groupId ) {
    //     $( ".table-ranking" ).each(function( idx, obj ) {
    //         var url = $( this ).data( "query" ) + "/org/" + orgId + "/group/" + groupId;
    //         $( this ).bootstrapTable( "refresh", {
    //             url: url,
    //             silent: true
    //         });
    //     });
    // }
    //
    //
    function updateEventTags( orgId, groupId ) {
        var url = '/getRealTimeLogs?',
            param = {
                limit: 30
            };
        if ( orgId > 1 ) {
            param.org_id = orgId;
            if ( groupId > 1 ) {
                param.group_id = groupId;
            }
        }
        url += $.param( param );
        console.log(url);

        var tags = '';
        $.ajax({
            type  : "GET",
            async : true,
            url   : url
        }).done( function( rows ) {
            $.each( rows, function( i, r ) {
                tags += '<div class="label label-default label-sm" style="display: block; display: inline-block; margin-right: 5px;">';
                tags += r.equip_id;
                tags += '</div>';
            });

        }).always( function() {
            $( "#event-tags" ).html( tags );
        });
    }

});

