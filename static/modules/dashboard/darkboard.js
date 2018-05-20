$(function() {
    var dom = document.getElementById( "chart-tags" );
    var myChart = echarts.init( dom );
    // var scaleData = [
    //     {
    //         'name': '工程建设',
    //         'value': 10
    //     },
    //     {
    //         'name': '土地交易',
    //         'value': 20
    //     },
    //     {
    //         'name': '其他交易',
    //         'value': 27
    //     },
    // ];
    // console.log(scaleData);
    var rich = {
        white: {
            color: '#ddd',
            align: 'center',
            padding: [5, 0]
        }
    };
    var placeHolderStyle = {
        normal: {
            label: {
                show: false
            },
            labelLine: {
                show: false
            },
            color: 'rgba(0, 0, 0, 0)',
            borderColor: 'rgba(0, 0, 0, 0)',
            borderWidth: 0
        }
    };

    window.onresize = function() {
        myChart.resize();
    };



    // updateTagsChart( scaleData );

    /**
     * 1. 초기화
     *
     */
    // 자산 (기관 / 그룹)
    var assets = { },
        interval = 211000,
        timer = null;

    // // 날짜
    // $( ".datetime" ).datetimepicker({
    //     format: "yyyy-mm-dd",
    //     pickerPosition : "bottom-left",
    //     todayHighlight : 1,
    //     minView: 2,
    //     maxView: 4,
    //     autoclose: true
    // });
    //
    initializeAssets();
    updateStats();
    // startTimer();
    // // $( "#modal-ipaslog" ).modal( "show" );

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

        // $( ".btn-test" ).click(function(e) {
        //     updateStats();
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
    //     updateRankings( orgId, groupId );
    //     updateLogs( orgId, groupId );
    //
    //     $( ".text-updated" ).removeClass( "hide" );
    //     setTimeout(function(){ $( ".text-updated" ).addClass( "hide" ); }, 500);
    }


    function updateSummary( orgId, groupId ) {

        // var scaleData = [
        //     {
        //         'name': '工程建设',
        //         'value': 10
        //     },
        //     {
        //         'name': '土地交易',
        //         'value': 20
        //     },
        //     {
        //         'name': '其他交易',
        //         'value': 27
        //     },
        // ];




        var url = "/stats/summary/org/" + orgId + "/group/" + groupId;
        $.ajax({
            type  : "GET",
            async : true,
            url   : url
        }).done( function( r ) {
            console.log(r);

            updateTagsChart( r.equipCountByType );

            $( ".count-startup" ).text( r.eventTypes[StartupEvent] );
            $( ".count-shock" ).text( r.eventTypes[ShockEvent] );
            $( ".count-speeding" ).text( r.eventTypes[SpeedingEvent] );
            $( ".count-proximity" ).text( r.eventTypes[ProximityEvent] );

            $( ".count-pt" ).text( r.equipCountByType[PT] );
            $( ".count-zt" ).text( r.equipCountByType[ZT] );
            $( ".count-vt" ).text( r.equipCountByType[VT] );
            $( ".count-total-equips" ).text( r.equipCountByType[PT] + r.equipCountByType[ZT] + r.equipCountByType[VT] );

    //         var  data = [
    //             {value:  r.eventTypes[StartupEvent], label: 'STARTUP'},
    //             {value:  r.eventTypes[ShockEvent], label: 'SHOCK'},
    //             {value:  r.eventTypes[SpeedingEvent], label: 'SPEEDING'},
    //             {value:  r.eventTypes[ProximityEvent], label: 'PROXIMITY'},
    //         ];
    //
    //         var total = r.eventTypes[ShockEvent] + r.eventTypes[SpeedingEvent] + r.eventTypes[ProximityEvent];
    //         if ( total > 0 ) {
    //             eventTypeChart.setData( data );
    //         } else {
    //             eventTypeChart.setData( [ { value: 0, label: 'N/A' } ] );
    //         }
        }).always( function() {
        });
    }


    function updateTagsChart( rawData ) {
        var scaleData = [];
        $.each( rawData, function( tagType, v ) {
            var name;
            if ( tagType == PT ) {
                name = "Pedestrian";
            } else if ( tagType == ZT ) {
                name = "Zone";
            } else if ( tagType == VT ) {
                name = "Vehicle";
            } else {
                name = "Unknown";
            }
            scaleData.push( { name: name, value: v } );
        });
        console.log( scaleData );

        var data = [];
        for (var i = 0; i < scaleData.length; i++) {
            data.push({
                value: scaleData[i].value,
                name: scaleData[i].name,
                itemStyle: {
                    normal: {
                        borderWidth: 5,
                        shadowBlur: 30,
                        borderColor: new echarts.graphic.LinearGradient(0, 0, 1, 1, [{
                            offset: 0,
                            color: '#7777eb'
                        }, {
                            offset: 1,
                            color: '#70ffac'
                        }]),
                        shadowColor: 'rgba(142, 152, 241, 0.6)'
                    }
                }
            }, {
                value: 4,
                name: '',
                itemStyle: placeHolderStyle
            });
        }
        var seriesObj = [{
            name: '',
            type: 'pie',
            clockWise: false,
            radius: [95, 100],
            hoverAnimation: false,
            itemStyle: {
                normal: {
                    label: {
                        show: true,
                        position: 'outside',
                        color: '#ddd',
                        formatter: function(params) {
                            // var percent = 0;
                            var total = 0;
                            for (var i = 0; i < scaleData.length; i++) {
                                total += scaleData[i].value;
                            }
                            var percent = ((params.value / total) * 100).toFixed(0);
                            if (params.name !== '') {
                                return params.name + '\n{white|' + '' + params.value + '}';
                            } else {
                                return '';
                            }
                        },
                        rich: rich
                    },
                    labelLine: {
                        show: false
                    }
                }
            },
            data: data
        }];
        var option = {
            // backgroundColor: '#04243E',
            tooltip: {
                show: false
            },
            legend: {
                show: false
            },
            toolbox: {
                show: false
            },
            series: seriesObj
        }

        if (option && typeof option === "object") {
            myChart.setOption(option, true);
        }
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
    //     var url = '/getRealTimeLogs?limit=5';
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

});

