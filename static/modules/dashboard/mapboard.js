$(function() {

    /**
     * 1. 초기화
     *
     */
    // 자산 (기관 / 그룹)
    var assets = { },
        interval = 10 * 1000,
        timer = null;

    // 날짜
    // $( ".datetime" ).datetimepicker({
    //     format:         "yyyy-mm-dd",
    //     pickerPosition: "bottom-left",
    //     todayHighlight: 1,
    //     minView:        2,
    //     maxView:        4,
    //     autoclose:      true
    // });

    // Chart - 이벤트 유형
    var eventTypeChart = Morris.Donut({
        element: 'chart-eventType',
        resize: true,
        data: [
            { label: "N/A", value: 0 }
        ],
        formatter: function (x) { return x + "" }
    }).on('click', function( i, row ){
        console.log(i, row);
    });

    initializeAssets();
    updateStats();
    startTimer();


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
        // updateRankings( orgId, groupId );
        // updateLogs( orgId, groupId );

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
            // console.log(r.operatingTime);
            // $( ".count-startup" ).text( r.eventTypes[StartupEvent] );
            $( ".count-shock" ).text( r.eventTypes[ShockEvent] );
            $( ".count-speeding" ).text( r.eventTypes[SpeedingEvent] );
            $( ".count-proximity" ).text( r.eventTypes[ProximityEvent] );

            // 비율
            var totalEvents = r.eventTypes[ShockEvent] + r.eventTypes[SpeedingEvent] + r.eventTypes[ProximityEvent];
            var total = totalEvents;
            if ( totalEvents === 0 ) {
                total = 1;
            }
            var shockRate = (r.eventTypes[ShockEvent] / total * 100).toFixed(1),
                speedingRate = (r.eventTypes[SpeedingEvent] / total * 100).toFixed(1),
                proximityRate = (r.eventTypes[ProximityEvent] / total * 100).toFixed(1);

            $( ".count-events" ).text( totalEvents );
            $( ".count-optime" ).text( sec2humanReadable( r.operatingTime ) );

            $( ".rate-shock" ).text( shockRate + "%" );
            $( ".rate-speeding" ).text( speedingRate + "%" );
            $( ".rate-proximity" ).text( proximityRate + "%" );

            $( "#pgb-shock" ).css( "width",shockRate + "%" );
            $( "#pgb-speeding" ).css( "width",speedingRate + "%" );
            $( "#pgb-proximity" ).css( "width",proximityRate + "%" );

            $( ".count-pt" ).text( r.equipCountByType[PedestrianTag] );
            $( ".count-zt" ).text( r.equipCountByType[ZoneTag] );
            $( ".count-vt" ).text( r.equipCountByType[VehicleTag] );
            $( ".count-total-tags" ).text( r.equipCountByType[PedestrianTag] + r.equipCountByType[ZoneTag] + r.equipCountByType[VehicleTag] );

            var  data = [
                // {value:  r.eventTypes[StartupEvent], label: 'STARTUP'},
                {value:  r.eventTypes[ShockEvent], label: 'SHOCK'},
                {value:  r.eventTypes[SpeedingEvent], label: 'SPEEDING'},
                {value:  r.eventTypes[ProximityEvent], label: 'PROXIMITY'},
            ];

            var total = r.eventTypes[ShockEvent] + r.eventTypes[SpeedingEvent] + r.eventTypes[ProximityEvent];
            if ( total > 0 ) {
                eventTypeChart.setData( data );
            } else {
                eventTypeChart.setData( [ { value: 0, label: 'N/A' } ] );
            }

            // time
            var time = moment(r.time);
            $( ".what-time" ).text( time.format("HH:mm") );
            $( ".what-date" ).text( time.format("ll") );

        }).always( function() {
        });
    }


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
    // function updateLogs( orgId, groupId ) {
    //     var url = '/getRealTimeLogs?limit=100';
    //     var activities = [];
    //     $( ".activity" ).each(function( idx, obj ) {
    //         if ( $( obj ).is(":checked") ) {
    //             activities.push( "event_type=" + $( obj ).val() );
    //         }
    //     });
    //
    //     if ( activities.length < 1 ) {
    //         $( "#table-ipaslogs" ).bootstrapTable( "removeAll" );
    //     } else {
    //         var urlSuffix = "";
    //         if ( orgId > 0 ) {
    //             urlSuffix += "&org_id=" + orgId;
    //             if ( groupId > 0 ) {
    //                 urlSuffix += "&group_id=" + groupId;
    //             }
    //         }
    //         $( "#table-ipaslogs" ).bootstrapTable( "refresh", {
    //             url: url + "&" + activities.join( "&" ) + urlSuffix,
    //             silent: true
    //         });
    //     }
    // }
    //
    //
    // function sec2humanReadable(duration) {
    //     var hour = 0;
    //     var min = 0;
    //     var sec = 0;
    //
    //     if (duration) {
    //         if (duration >= 60) {
    //             min = Math.floor(duration / 60);
    //             sec = duration % 60;
    //         }
    //         else {
    //             sec = duration;
    //         }
    //
    //         if (min >= 60) {
    //             hour = Math.floor(min / 60);
    //             min = min - hour * 60;
    //         }
    //
    //         if (hour < 10) {
    //             hour = '0' + hour;
    //         }
    //         if (min < 10) {
    //             min = '0' + min;
    //         }
    //         if (sec < 10) {
    //             sec = '0' + sec;
    //         }
    //     }
    //
    //     return hour +":"+ min +":"+ sec;
    // }

});

